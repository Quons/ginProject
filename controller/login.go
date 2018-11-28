package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/sirupsen/logrus"
	_ "ginProject/ginProject/log"
)

func Getting(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "hello %s\n", name)
}
func AnotherGetting(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, "%s\n", message)
}

func StrGetting(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	//这个是c.Request.URL.Query().Get("lastname")的快捷方式
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s\n", firstname, lastname)
}

func PostFunc(c *gin.Context) {
	message := c.PostForm("message")
	age := c.PostForm("age")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
		"age":     age,
	})
}

func ApostFunc(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func UploadFunc(c *gin.Context) {
	//单文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadGateway, "error")
		return
	}
	log.Info(file.Filename)
	c.String(http.StatusOK, "filename:%s", file.Filename)
}
func MultiUploadFunc(c *gin.Context) {
	//单文件
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusOK, "error:%v", err)
		return
	}
	files := form.File["upload[]"]
	for _, file := range files {
		log.Info(file.Filename)
	}
	c.String(http.StatusOK, "%d files upload!", len(files))
}
