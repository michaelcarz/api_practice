package main

import (
	"fmt"
	"golangAPI/database"
	"golangAPI/objj"
	"golangAPI/src"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	t := template.Must(template.ParseFiles("./htmldemo/insert.html"))
// 	t.Execute(w, nil)
// }

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	router.LoadHTMLGlob("./htmldemo/*") //Load all html files in the folder

	router.GET("/test", MainPage)
	router.GET("/insert", InsertPage)
	router.POST("/insert", InsertPage)
	router.GET("/update", UpdatePage)
	router.POST("/update", UpdatePage)

	go func() { //when running the main and it will Synchronously connect  to the database
		database.DD()
		_, err1 := database.DBConnect.DB()
		if err1 != nil {
			fmt.Println("get db failed:", err1)
			return
		}
		//產生table
		database.DBConnect.Debug().AutoMigrate(&objj.User{})
		//判斷有沒有table存在
		migrator := database.DBConnect.Migrator()
		has := migrator.HasTable(&objj.User{})
		//has := migrator.HasTable("GG")
		if !has {
			fmt.Println("table not exist")
		}
	}()

	router.Run(":8080")
}

func InsertPage(c *gin.Context) {
	c.HTML(http.StatusOK, "insert.html", nil)

	//c.String(200, "Success")
}
func UpdatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "update.html", nil)

	//c.String(200, "Success")
}

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}
