package main

import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
		Id int
		ProjectName string
		StartDate string
		EndDate string
		Duration string
		Description string
		NodeJs     bool
		ReactJs    bool
		Golang     bool
		Javascript bool
		Image string
}

var projectData = []Project{
	{
		ProjectName:"Project 1",
		StartDate:  "2020-01-15",
		EndDate:    "2020-02-15",
		Duration:   countDuration("2020-01-15", "2020-02-15"),
		Description: "<b>This is the description of project 1</b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
		NodeJs:     true,
		ReactJs:    false,
		Golang:     true,
		Javascript: true,
		Image: 		"img1.png",
	},
	{
		ProjectName:"Project 2",
		StartDate:  "2023-06-05",
		EndDate:    "2023-06-06",
		Duration:   countDuration("2023-06-05", "2023-06-06"),
		Description: "<b>This is the description of project 2</b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
		NodeJs:     false,
		ReactJs:    false,
		Golang:     true,
		Javascript: false,
		Image: 		"img2.png",
	},
	{
		ProjectName:"Project 3",
		StartDate:  "2022-06-05",
		EndDate:    "2023-06-06",
		Duration:   countDuration("2022-06-05", "2023-06-06"),
		Description: "<b>This is the description of project 3</b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
		NodeJs:     true,
		ReactJs:    true,
		Golang:     true,
		Javascript: true,
		Image: 		"img3.png",
	},
	{
		ProjectName:"Project 4",
		StartDate:  "2022-06-05",
		EndDate:    "2023-06-06",
		Duration:   countDuration("2022-06-05", "2023-06-06"),
		Description: "<b>This is the description of project 4</b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
		NodeJs:     true,
		ReactJs:    true,
		Golang:     true,
		Javascript: true,
		Image: 		"img4.png",
	},
	{
		ProjectName:"Project 5",
		StartDate:  "2022-01-05",
		EndDate:    "2023-06-06",
		Duration:   countDuration("2022-06-05", "2023-06-06"),
		Description: "<b>This is the description of project 5 </b> <br> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis ",
		NodeJs:     true,
		ReactJs:    true,
		Golang:     true,
		Javascript: true,
		Image: 		"img5.png",
	},
}


func main() {

	e := echo.New()

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
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}

	projects := map[string]interface{}{
		"Projects": projectData,
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
				Image: 			data.Image,
			}
		}
	}
	data := map[string]interface{}{
		"Project":   ProjectDetail,
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
				Image: 			data.Image,
			}
		}
	}
	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"Id": id,
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
		image := c.FormValue("input-image")

		var newProject = Project{
			ProjectName: projectName,
			StartDate:  startDate,
			EndDate:    endDate,
			Duration:   countDuration(startDate, endDate),
			Description: description,
			NodeJs:     (nodeJs == "nodeJs"),
			ReactJs:    (reactJs == "reactJs"),
			Golang:     (golang == "golang"),
			Javascript: (javascript == "javascript"),
			Image: 		image,
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
		image := c.FormValue("input-image")

		var editedProject = Project{
			ProjectName: projectName,
			StartDate:  startDate,
			EndDate:    endDate,
			Duration:   countDuration(startDate, endDate),
			Description: description,
			NodeJs:     (nodeJs == "nodeJs"),
			ReactJs:    (reactJs == "reactJs"),
			Golang:     (golang == "golang"),
			Javascript: (javascript == "javascript"),
			Image: 		image,
		}

		projectData[id] = editedProject
		return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	projectData = append(projectData[:id], projectData[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func countDuration(d1 string, d2 string) string {
	date1, _ := time.Parse("2006-01-02", d1)
	date2, _ := time.Parse("2006-01-02", d2)

	diff := date2.Sub(date1)
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