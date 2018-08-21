// main.go

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine
var namedatabase = "test1"
var userdatabase = "root"
var passdatabase = "ilovesomeone"

//-------------- // Struct
type datasetschema struct {
	ID       string
	Username string
	Password string
}

//-------------- // My Engine
func openDB(stringquery string) (rowsreturn *sql.Rows) {
	db, err := sql.Open("mysql", fmt.Sprint(userdatabase, ":", passdatabase, "@/", namedatabase, "?charset=utf8"))
	checkErr(err)
	rows, err := db.Query(stringquery)
	checkErr(err)
	return rows
}
func executeDB(stringquery string) (stringreturn bool) {
	db, err := sql.Open("mysql", fmt.Sprint(userdatabase, ":", passdatabase, "@/", namedatabase, "?charset=utf8"))
	_, err = db.Exec(stringquery)
	if err != nil {
		// fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//--------------
func CachedPageHandler(c *gin.Context) {
	//Do what you need to get the cached html
	yourHtmlString := "<html><body>I am cached HTML!</body></html>"

	//Write your 200 header status (or other status codes, but only WriteHeader once)
	c.Writer.WriteHeader(http.StatusOK)
	//Convert your cached html string to byte array
	c.Writer.Write([]byte(yourHtmlString))
	return
}

func preveiew(c *gin.Context) {
	var dataset = []datasetschema{}
	var rows = openDB("select * from `dbmaccount`")
	for rows.Next() {
		var r datasetschema
		rows.Scan(&r.ID, &r.Username, &r.Password)
		dataset = append(dataset, r)
	}
	c.HTML(http.StatusOK, "show.html",
		gin.H{
			"title":   "funtion prevuew",
			"dataset": dataset,
		},
	)
}
func insert(c *gin.Context) {
	idaccount := c.PostForm("idaccount")
	username := c.PostForm("username")
	password := c.PostForm("password")
	var sql = ""
	if idaccount == "" {
		sql = fmt.Sprint("insert into dbmaccount (`username`,`password`) value ('", username, "','", password, "' )")
	} else {
		sql = fmt.Sprint("update dbmaccount set username='", username, "' , password='", password, "' where id='", idaccount, "' ")
	}
	executeDB(sql)
	preveiew(c)
}
func delete(c *gin.Context) {
	idaccount := c.Param("id")
	var sql = ""
	if idaccount != "" {
		sql = fmt.Sprint("delete from dbmaccount where id='", idaccount, "' ")
		executeDB(sql)
	}
	preveiew(c)
}

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", preveiew)
	router.GET("/preview", preveiew)
	router.GET("/insertpreview", func(c *gin.Context) {
		c.HTML(http.StatusOK, "insert.html", gin.H{})
	})
	router.POST("/insertsystem", insert)
	router.GET("/editpreview/:id", func(c *gin.Context) {
		var idaccount = c.Param("id")
		var rows = openDB(fmt.Sprint("select username,password from dbmaccount where id='", idaccount, "'"))
		var username string
		var password string
		for rows.Next() {
			rows.Scan(&username, &password)
		}
		c.HTML(http.StatusOK, "insert.html", gin.H{
			"idaccount": idaccount,
			"username":  username,
			"password":  password,
		})
	})
	router.GET("delete/:id", delete)

	router.Run()
}
