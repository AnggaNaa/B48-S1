package main

import (
	"b48s1/connection"
	"b48s1/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
		Id int
		ProjectName string
		StartDate time.Time
		EndDate time.Time
		Duration string
		Description string
		Technologies []string
		NodeJs     bool
		ReactJs    bool
		Golang     bool
		Javascript bool
		Image string
		Author string
}

type User struct {
		Id int 
		Name string
		Email string
		Password string
}

type SessionData struct {
	IsLogin bool
	Name string
}

var userData = SessionData{}

func main() {

	// e = echo package
	e := echo.New()

	// connect to package database
	connection.DatabaseConnect()

	// Set the handling of static files
	e.Static("/public", "public")
	e.Static("/uploads", "uploads")

	// To use sessions using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// List Routes GET
	e.GET("/", 					home)
	e.GET("/contact", 			contact)
	e.GET("/testimonial", 		testimonial)
	e.GET("/add-project", 		addProject)
	e.GET("/edit-project/:id", 	editProject)
	e.GET("/myproject", 		myproject)
	e.GET("/project/:id", 		projectDetail)
	e.GET("/form-register", 	formRegister)
	e.GET("/form-login", 		formLogin)
	
	//List Routes POST
	e.POST("/", 					middleware.UploadFile(submitProject))
	e.POST("/edit-project/:id", 	middleware.UploadFile(submitEditedProject))
	e.POST("/delete-project/:id", 	deleteProject)
	e.POST("/register", 			register)
	e.POST("/login", 				login)
	e.POST("/logout", 				logout)

	// Server
	e.Logger.Fatal(e.Start("localhost:8000"))
}


func home(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_projects.id, name_project, start_date, end_date, description, image, technologies, tb_users.name AS author FROM tb_projects JOIN tb_users ON tb_projects.author_id = tb_users.id")

	var projectData []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Technologies, &each.Author)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Duration = countDuration(each.StartDate, each.EndDate)
		if checkValue(each.Technologies, "nodeJs") {
			each.NodeJs = true
		}
		if checkValue(each.Technologies, "reactJs") {
			each.ReactJs = true
		}
		if checkValue(each.Technologies, "golang") {
			each.Golang = true
		}
		if checkValue(each.Technologies, "javascript") {
			each.Javascript = true
		}

		projectData = append(projectData, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"Projects": projectData,
		"FlashStatus": sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession": userData,
	}

	delete(sess.Values, "status")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}


	return tmpl.Execute(c.Response(), projects)
}

func myproject(c echo.Context) error {
	
	sess, _ := session.Get("session", c)
	var projectData []Project
	
	if sess.Values["isLogin"] != true {
		userData.IsLogin = false 
			return c.Redirect(http.StatusMovedPermanently, "/form-login")
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
		userId := sess.Values["id"]
		data, _ := connection.Conn.Query(context.Background(), "SELECT tb_projects.id, name_project, start_date, end_date, description, image, technologies, tb_users.name AS author FROM tb_projects LEFT JOIN tb_users ON tb_projects.author_id = tb_users.id WHERE author_id = $1 ORDER BY tb_projects.id", userId)
		
		for data.Next() {
			var each = Project{}
	
			err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Technologies, &each.Author)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
			}
	
			each.Duration = countDuration(each.StartDate, each.EndDate)
			if checkValue(each.Technologies, "nodeJs") {
				each.NodeJs = true
			}
			if checkValue(each.Technologies, "reactJs") {
				each.ReactJs = true
			}
			if checkValue(each.Technologies, "golang") {
				each.Golang = true
			}
			if checkValue(each.Technologies, "javascript") {
				each.Javascript = true
			}
	
			projectData = append(projectData, each)
		}
	}


	// sess, _ := session.Get("session", c)

	// if sess.Values["isLogin"] != true {
	// 	userData.IsLogin = false
	// } else {
	// 	userData.IsLogin = sess.Values["isLogin"].(bool)
	// 	userData.Name = sess.Values["name"].(string)
	// }

	projects := map[string]interface{}{
		"Projects": projectData,
		"DataSession": userData,
		// "Author" : userData.Name,
	}

	var tmpl, err = template.ParseFiles("views/myproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}


	return tmpl.Execute(c.Response(), projects)
}

func addProject(c echo.Context) error {


	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
		return c.Redirect(http.StatusMovedPermanently, "/form-login")
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"DataSession": userData,
	}


	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

func contact(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"DataSession": userData,
	}

	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

func testimonial(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"DataSession": userData,
	}

	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT tb_projects.id, name_project, start_date, end_date, description, image, technologies, tb_users.name AS author FROM tb_projects JOIN tb_users ON tb_projects.author_id = tb_users.id WHERE tb_projects.id =$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Description, &ProjectDetail.Image, &ProjectDetail.Technologies, &ProjectDetail.Author,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	ProjectDetail.Duration = countDuration(ProjectDetail.StartDate, ProjectDetail.EndDate)
	
	if checkValue(ProjectDetail.Technologies, "nodeJs") {
		ProjectDetail.NodeJs = true
	}
	if checkValue(ProjectDetail.Technologies, "reactJs") {
		ProjectDetail.ReactJs = true
	}
	if checkValue(ProjectDetail.Technologies, "golang") {
		ProjectDetail.Golang = true
	}
	if checkValue(ProjectDetail.Technologies, "javascript") {
		ProjectDetail.Javascript = true
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	
	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"startDateString" 	: ProjectDetail.StartDate.Format("2006-01-02"),
		"endDateString"		: ProjectDetail.EndDate.Format("2006-01-02"),
		"DataSession": userData,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project-detail.html")
	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func submitProject(c echo.Context) error {
	
	projectName := c.FormValue("input-name")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("input-description")
	nodeJs := c.FormValue("nodeJs")
	reactJs := c.FormValue("reactJs")
	golang := c.FormValue("golang")
	javascript := c.FormValue("javascript")
	tech := []string{nodeJs,reactJs,golang,javascript}
	image := c.Get("dataFile").(string)
	
	sess, _ := session.Get("session", c)
	author := sess.Values["id"].(int)
	
	
	_, err := connection.Conn.Exec(context.Background(),
			"INSERT INTO tb_projects (name_project, start_date, end_date, description, image, technologies, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			projectName, startDate, endDate, description, image, tech, author,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	
		

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func submitEditedProject(c echo.Context) error {
	
	// Menangkap Id dari Query Params
		id, _:= strconv.Atoi(c.Param("id"))

		projectName := c.FormValue("input-name")
		startDate := c.FormValue("startDate")
		endDate := c.FormValue("endDate")
		description := c.FormValue("input-description")
		nodeJs := c.FormValue("nodeJs")
		reactJs := c.FormValue("reactJs")
		golang := c.FormValue("golang")
		javascript := c.FormValue("javascript")
		tech := []string{nodeJs,reactJs,golang,javascript}
		image := c.Get("dataFile").(string)

		start, _ := time.Parse("2006-01-02", startDate)
		end, _ := time.Parse("2006-01-02", endDate)

		oldImageName, err := getProjectImageName(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching project image name"})
		}

		// Jika ada gambar baru yang diunggah, hapus gambar lama dan ganti dengan yang baru
		if image != "" {
			// Hapus file gambar lama dari folder "uploads"
			err := deleteImageFile("uploads/" + oldImageName)
			if err != nil {
				// Jika terjadi kesalahan saat menghapus file gambar, berikan tanggapan tetapi tidak hentikan proses
				fmt.Println("Kesalahan saat menghapus file gambar:", err)
			}
		} else {
			// Jika tidak ada gambar baru yang diunggah, gunakan gambar lama
			image = oldImageName
		}

		_, err1 := connection.Conn.Exec(context.Background(),
			"UPDATE tb_projects SET name_project=$1, start_date=$2, end_date=$3, description=$4, image=$5, technologies=$6 WHERE id=$7",
			projectName, start, end, description, image, tech, id,
		)

		if err1 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	projectImageName, err := getProjectImageName(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching project image name"})
	}

	err = deleteImageFile("uploads/" + projectImageName)
	if err != nil {
		fmt.Println("Kesalahan saat menghapus file gambar:", err)
	}

	_, err1 := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}


	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func editProject(c echo.Context) error {

	if userData.IsLogin != true {
		return c.Redirect(http.StatusMovedPermanently, "/form-login")
	}


	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name_project, start_date, end_date, description, image, technologies FROM tb_projects WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Description, &ProjectDetail.Image, &ProjectDetail.Technologies,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	ProjectDetail.Duration = countDuration(ProjectDetail.StartDate, ProjectDetail.EndDate)
	
	if checkValue(ProjectDetail.Technologies, "nodeJs") {
		ProjectDetail.NodeJs = true
	}
	if checkValue(ProjectDetail.Technologies, "reactJs") {
		ProjectDetail.ReactJs = true
	}
	if checkValue(ProjectDetail.Technologies, "golang") {
		ProjectDetail.Golang = true
	}
	if checkValue(ProjectDetail.Technologies, "javascript") {
		ProjectDetail.Javascript = true
	}
	
	start := ProjectDetail.StartDate.Format("2006-01-02")
	end := ProjectDetail.EndDate.Format("2006-01-02")
	
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	
	data := map[string]interface{}{
		"Project"			:   ProjectDetail,
		"StartDate": start,
		"EndDate":   end,
		"DataSession": userData,
	}
	
	var tmpl, errTemplate = template.ParseFiles("views/edit-project.html")
	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}

func formRegister(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if userData.IsLogin != false {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}


	data := map[string]interface{}{
		"DataSession": userData,
		"FlashStatus": sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func register(c echo.Context)error {
	// to makse sure request body is form data format, not JSON, XML, etc.
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("input-name")
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_users WHERE email =$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err == nil {
		return redirectWithMessage(c, "Email already in use !", false,  "/form-register")
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_users(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}
		return redirectWithMessage(c, "Register success!", true, "/form-login")

}

func formLogin(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if userData.IsLogin != false {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	
	
	flash :=map[string]interface{}{
		"FlashStatus": sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession": userData,
	}
	
	delete(sess.Values, "status")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request(). ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_users WHERE email =$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false,  "/form-login")
	}

	// membandingkan password yg di db dengan yg di html
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM MASA BERLAKU LOGIN
	sess.Values["message"] = "Login success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(),  c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge= -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}


func countDuration(dateStart time.Time, dateEnd time.Time) string {

	diff := dateEnd.Sub(dateStart)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	
	if months > 12 {
		return strconv.Itoa(months/12) + " tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " minggu"
	}
	return strconv.Itoa(days) + " hari"
}

func checkValue(slice []string, object string) bool {
	for _, data := range slice {
		if data == object {
			return true
		}
	}
	return false
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}


func getProjectImageName(id int) (string, error){
	var imageName string
	err := connection.Conn.QueryRow(context.Background(), "SELECT image FROM tb_projects WHERE id=$1", id).Scan(&imageName,)
	if err != nil {
		return "", err
	}

	fmt.Println("Image name retrieved from database:", imageName)
	return imageName, nil
}

func deleteImageFile(filePath string) error{
	err := os.Remove(filePath)
	return err
}