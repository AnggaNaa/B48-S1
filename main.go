package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

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
	e.GET("/add-project", addProject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)

	//Daftar Routes POST
	e.POST("/", addFormProject)

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
	return tmpl.Execute(c.Response(), nil)
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

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id" : id,
		"Title" : "Dumbways Web App",
		"Content": "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Omnis optio iste sapiente provident fugit, impedit adipisci deserunt obcaecati voluptatum quo, hic culpa, repudiandae consectetur quia itaque eiusaccusamus natus unde. Lorem ipsum dolor sit amet consecteturadipisicing elit. Id sunt porro cupiditate, nesciunt totam aliaslabore fugiat? Accusantium a voluptates quibusdam tempora, evenietquasi sit debitis eaque ut in magnam? Lorem ipsum dolor sit amet,consectetur adipisicing elit. Dolor minima quia recusandae doloribusfacilis fugit optio, quaerat fugiat quod architecto id dignissimoscupiditate perferendis similique saepe totam, nulla nihil ipsam. Loremipsum dolor sit amet consectetur adipisicing elit. Incidunt laboreoptio, eos qui quasi, soluta fugit totam tempore, in amet underepellat perferendis quibusdam ducimus velit deserunt maxime possimusfacilis? Lorem ipsum dolor sit amet consectetur adipisicing elit.Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiat ihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaes Veritatis accusamus eius, consequuntur possimus similique teneturfugiat incidunt minima necessitatibus eum ab enim reiciendis autem iddicta eaque libero. Ea ex ducimus cupiditate veniam voluptatibus,labore nam eius vero debitis, minima doloremque, nostrum sit etconsequuntur in velit totam nemo adipisci possimus consectetur placeatodit. Sed voluptas illo, praesentium, sapiente magni pariatur evenietmaiores nesciunt quibusdam inventore eligendi ex, aliquid fugiatnihil. Voluptates ipsa magnam rerum atque in magni, harum repudiandaeipsam quae quo similique! Neque quibusdam facilis debitis nonrepellendus asperiores laudantium obcaecati nihil necessitatibusplaceat, eum animi, veniam harum. Lorem ipsum dolor sit ametconsectetur adipisicing elit. Officia praesentium eum perspiciatisquas, nemo magni explicabo aspernatur natus dolore, hic totam,adipisci id. Obcaecati sequi, officiis explicabo in ducimus et aliquidinventore molestiae nam suscipit quisquam accusamus? Ratione autemdicta dolores animi illo veniam pariatur ipsa eveniet nulla id velitminima totam mollitia beatae porro iusto, numquam saepe iure velvoluptates quisquam et? Tempore, illum unde, ea explicabo oditlaudantium totam, vero accusamus natus aliquid eveniet neque odio.Incidunt ducimus quia quibusdam ab perferendis culpa doloresaccusantium nesciunt voluptatibus est necessitatibus quod omnis, seddeserunt asperiores cumque quo odio. Possimus quod debitis hicvoluptates earum minima, saepe quasi. Amet error nemo sapiente quidem,nesciunt, ea cupiditate eos ab temporibus reiciendis doloribus quaeratmaiores ex. Quia nihil hic ratione facilis quas fuga et ducimusexpedita ad omnis quis corporis dignissimos accusamus minus doloribusaut, quidem ipsa voluptatibus officia voluptas temporibus? Ab temporevitae illo iusto, exercitationem debitis neque, aspernatur sequi fugitassumenda amet voluptatibus veritatis laborum nam, non dolorem eiusperspiciatis facere possimus maxime pariatur! Eius, quae? Doloremquesunt dolorum magni aperiam, iste illum optio incidunt officiis",
	}

	var tmpl, err = template.ParseFiles("views/blog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addFormProject(c echo.Context) error {
	
		nameProject := c.FormValue("input-name")
		startDate := c.FormValue("startDate")
		endDate := c.FormValue("endDate")
		description := c.FormValue("input-description")
		nodeJs := c.FormValue("nodeJs")
		reactJs := c.FormValue("reactJs")
		golang := c.FormValue("golang")
		javascript := c.FormValue("javascript")
		image := c.FormValue("input-image")


	println("Name Project : " + nameProject)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Description : " + description)
	println("Technologies : " + nodeJs)
	println("Technologies : " + reactJs)
	println("Technologies : " + golang)
	println("Technologies : " + javascript)
	println("Image : " + image)

	return c.Redirect(http.StatusMovedPermanently, "/")
}