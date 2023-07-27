package main

import (
	"b48s1/connection"
	"context"
	"fmt"
	"log"
	"net/http"
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

// var projectData = []Project{
// 	{
// 		ProjectName:"Project 1",
// 		StartDate:  "2020-01-15",
// 		EndDate:    "2020-02-15",
// 		Duration:   countDuration("2020-01-15", "2020-02-15"),
// 		Description: "<b>This is the description of project 1</b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
// 		NodeJs:     true,
// 		ReactJs:    false,
// 		Golang:     true,
// 		Javascript: true,
// 	},
// }


func main() {

	e := echo.New()
	connection.DatabaseConnect()

	// fmt.println("Hello World!")
	// e = echo package
	// GET/POST = run the method
	// "/" = endpoint/routing (ex. localhost:5000'/' ( ex. dumbways/lms)
	// helloWord = function that will run if the routes are opened

	// Mengatur penanganan file static
	e.Static("/public", "public")

	// To use sessions using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// Daftar Routes GET
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/add-project", addProject)
	e.GET("/edit-project/:id", editProject)
	e.GET("/myproject", myproject)
	e.GET("/project/:id", projectDetail)
	e.GET("/form-register", formRegister)
	e.GET("/form-login", formLogin)
	
	//Daftar Routes POST
	e.POST("/", submitProject)
	e.POST("/edit-project/:id", submitEditedProject)
	e.POST("/delete-project/:id", deleteProject)
	e.POST("/register", register)
	e.POST("/login", login)
	e.POST("/logout", logout)

	// Server
	e.Logger.Fatal(e.Start("localhost:8000"))
}


func home(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM tb_projects")

	var projectData []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Technologies)
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
	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM tb_projects")

	var projectData []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.Technologies)
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
		"DataSession": userData,
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

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_projects WHERE id=$1", id).Scan(
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

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_projects WHERE id=$1", id).Scan(
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

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"DataSession": userData,
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

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_users(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}
		return redirectWithMessage(c, "Register success!", true, "/form-login")

}

func formLogin(c echo.Context) error {

	sess, _ := session.Get("session", c)
	
	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
		} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
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
	sess.Values["Id"] = user.Id
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

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
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
		image := c.FormValue("input-image")


		_, err := connection.Conn.Exec(context.Background(),
			"INSERT INTO tb_projects (name_project, start_date, end_date, description, image, technologies) VALUES ($1, $2, $3, $4, $5, $6)",
			projectName, startDate, endDate, description, image, tech,
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
		image := c.FormValue("input-image")

		start, _ := time.Parse("2006-01-02", startDate)
		end, _ := time.Parse("2006-01-02", endDate)

		// var tech []string
		// if nodeJs == "nodeJs" {
		// 	tech = append(tech, "nodeJs")
		// }
		// if reactJs == "reactJs" {
		// 	tech = append(tech, "reactJs")
		// }
		// if golang == "golang" {
		// 	tech = append(tech, "golang")
		// }
		// if javascript == "javascript" {
		// 	tech = append(tech, "javascript")
		// }
		
		// merge := strings.Join(tech, ",")

		_, err := connection.Conn.Exec(context.Background(),
			"UPDATE tb_projects SET name_project=$1, start_date=$2, end_date=$3, description=$4, image=$5, technologies=$6 WHERE id=$7",
			projectName, start, end, description, image, tech, id,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func countDuration(d1 time.Time, d2 time.Time) string {

	diff := d2.Sub(d1)
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