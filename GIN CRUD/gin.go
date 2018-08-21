package main

// ------------------------------
// type Friend struct {
// 	Fname string
// }

// type Person struct {
// 	UserName string
// 	Emails   []string
// 	Friends  []*Friend
// }

// func main() {
// 	f1 := Friend{Fname: "minux.ma"}
// 	f2 := Friend{Fname: "xushiwei"}
// 	t := template.New("fieldname example")
// 	t, _ = t.Parse(`hello {{.UserName}}!
//             {{range .Emails}}
//                 an email {{.}}
//             {{end}}
//             {{with .Friends}}
//             {{range .}}
//                 my friend name is {{.Fname}}
//             {{end}}
//             {{end}}
//             `)
// 	p := Person{UserName: "Astaxie",
// 		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
// 		Friends: []*Friend{&f1, &f2}}
// 	t.Execute(os.Stdout, p)
// }

// ------------------------------
// func main() {
// 	router := gin.Default()
// 	router.POST("/post", func(c *gin.Context) {
// 		id := c.Query("id")
// 		page := c.DefaultQuery("page", "0")
// 		name := c.PostForm("name")
// 		message := c.PostForm("message")
// 		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
// 	})
// 	router.Run(":8080")
// }

//------------------------------
// func main() {
// 	router := gin.Default()

// 	// Query string parameters are parsed using the existing underlying request object.
// 	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
// 	router.GET("/welcome", func(c *gin.Context) {
// 		firstname := c.DefaultQuery("firstname", "Guest")
// 		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

// 		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// 	})
// 	router.Run(":8080")
// }

//--------------------------------
// func main() {
// 	router := gin.Default()

// 	// This handler will match /user/john but will not match /user/ or /user
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	// However, this one will match /user/john/ and also /user/john/send
// 	// If no other routers match /user/john, it will redirect to /user/john/
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	router.Run(":8080")
// }

//--------------------------------
// func main() {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()

// 	// Creates a gin router with default middleware:
// 	// logger and recovery (crash-free) middleware
// 	router := gin.Default()

// 	router.GET("/someGet", getting)
// 	router.POST("/somePost", posting)
// 	router.PUT("/somePut", putting)
// 	router.DELETE("/someDelete", deleting)
// 	router.PATCH("/somePatch", patching)
// 	router.HEAD("/someHead", head)
// 	router.OPTIONS("/someOptions", options)

// 	// By default it serves on :8080 unless a
// 	// PORT environment variable was defined.
// 	router.Run()
// 	// router.Run(":3000") for a hard coded port
// }
