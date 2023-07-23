package main

import (
	"b48s1/connection"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
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

var projectData = []Project{
	// {
	// 	ProjectName:"Project 1",
	// 	StartDate:  "2020-01-15",
	// 	EndDate:    "2020-02-15",
	// 	Duration:   countDuration("2020-01-15", "2020-02-15"),
	// 	Description: "<b>This is the description of project 1</b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
	// 	NodeJs:     true,
	// 	ReactJs:    false,
	// 	Golang:     true,
	// 	Javascript: true,
	// },
}


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

	// Daftar Routes GET
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/add-project", addProject)
	e.GET("/edit-project/:id", editProject)
	e.GET("/project/:id", projectDetail)
	
	//Daftar Routes POST
	e.POST("/", submitProject)
	e.POST("/edit-project/:id", submitEditedProject)
	e.POST("/delete-project/:id", deleteProject)

	// Server
	e.Logger.Fatal(e.Start("localhost:8000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Worldl!")
}

func home(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name_project, start_date, end_date, description, technologies, image FROM tb_projects")

	projectData = []Project{}
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
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

	projects := map[string]interface{}{
		"Projects": projectData,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}


	return tmpl.Execute(c.Response(), projects)
}

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range projectData {
		if id == i {
			ProjectDetail = Project{
				ProjectName:    data.ProjectName,
				StartDate:  	data.StartDate,
				EndDate:    	data.EndDate,
				Duration:   	data.Duration,
				Description: 	data.Description,
				NodeJs:     	data.NodeJs,
				ReactJs:    	data.ReactJs,
				Golang:     	data.Golang,
				Javascript: 	data.Javascript,
				Image:			data.Image,
			}
		}
	}
	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"startDateString" 	: ProjectDetail.StartDate.Format("2006-01-02"),
		"endDateString"		: ProjectDetail.EndDate.Format("2006-01-02"),
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range projectData {
		if id == i {
			ProjectDetail = Project{
				ProjectName:    data.ProjectName,
				StartDate:  	data.StartDate,
				EndDate:    	data.EndDate,
				Duration:   	data.Duration,
				Description: 	data.Description,
				NodeJs:     	data.NodeJs,
				ReactJs:    	data.ReactJs,
				Golang:     	data.Golang,
				Javascript: 	data.Javascript,
			}
		}
	}
	data := map[string]interface{}{
		"Project"			:   ProjectDetail,
		
	}

	var tmpl, err = template.ParseFiles("views/edit-project.html")
	if err != nil {
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

		start, _ := time.Parse("2006-01-02", startDate)
		end, _ := time.Parse("2006-01-02", endDate)

		var newProject = Project{
			ProjectName: projectName,
			StartDate:  start,
			EndDate:    end,
			Duration:   countDuration(start, end),
			Description: description,
			NodeJs:     (nodeJs == "nodeJs"),
			ReactJs:    (reactJs == "reactJs"),
			Golang:     (golang == "golang"),
			Javascript: (javascript == "javascript"),
		}

		projectData = append(projectData, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
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

		start, _ := time.Parse("2006-01-02", startDate)
		end, _ := time.Parse("2006-01-02", endDate)

		var editedProject = Project{
			ProjectName: projectName,
			StartDate:  start,
			EndDate:    end,
			Duration:   countDuration(start, end),
			Description: description,
			NodeJs:     (nodeJs == "nodeJs"),
			ReactJs:    (reactJs == "reactJs"),
			Golang:     (golang == "golang"),
			Javascript: (javascript == "javascript"),
		}

		projectData[id] = editedProject
		return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	projectData = append(projectData[:id], projectData[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
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