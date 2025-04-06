package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks []Task
var nextID = 1

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": tasks,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		if title != "" {
			tasks = append(tasks, Task{ID: nextID, Title: title, Done: false})
			nextID++
		}
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/done/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				break
			}
		}
		c.Redirect(http.StatusFound, "/")
	})

	r.Run(":8080")
}
