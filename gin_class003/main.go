package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()

// 	r.GET("/someJSON", func(c *gin.Context) {
// 		data := map[string]interface{}{
// 			"lang": "GO语言",
// 			"tag":  "<br>",
// 		}

// 		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
// 		c.AsciiJSON(http.StatusOK, data)
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/JSONP", func(c *gin.Context) {
// 		data := map[string]interface{}{
// 			"foo": "bar",
// 		}

// 		// /JSONP?callback=x
// 		// 将输出：x({\"foo\":\"bar\"})
// 		c.JSONP(http.StatusOK, data)
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

// type LoginForm struct {
// 	User     string `form:"user" binding:"required"`
// 	Password string `form:"password" binding:"required"`
// }

// func main() {
// 	router := gin.Default()
// 	router.POST("/login", func(c *gin.Context) {
// 		// 你可以使用显式绑定声明绑定 multipart form：
// 		// c.ShouldBindWith(&form, binding.Form)
// 		// 或者简单地使用 ShouldBind 方法自动绑定：
// 		var form LoginForm
// 		// 在这种情况下，将自动选择合适的绑定
// 		if c.ShouldBind(&form) == nil {
// 			if form.User == "user" && form.Password == "password" {
// 				c.JSON(200, gin.H{"status": "you are logged in"})
// 			} else {
// 				c.JSON(401, gin.H{"status": "unauthorized"})
// 			}
// 		}
// 	})
// 	router.Run(":8080")
// }

// func main() {
// 	router := gin.Default()

// 	router.POST("/form_post", func(c *gin.Context) {
// 		message := c.PostForm("message")
// 		nick := c.DefaultPostForm("nick", "anonymous")

// 		c.JSON(200, gin.H{
// 			"status":  "posted",
// 			"message": message,
// 			"nick":    nick,
// 		})
// 	})
// 	router.Run(":8080")
// }

// func main() {
// 	r := gin.Default()

// 	// 提供 unicode 实体
// 	r.GET("/json", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"html": "<b>Hello, world!</b>",
// 		})
// 	})

// 	// 提供字面字符
// 	r.GET("/purejson", func(c *gin.Context) {
// 		c.PureJSON(200, gin.H{
// 			"html": "<b>Hello, world!</b>",
// 		})
// 	})

// 	// 监听并在 0.0.0.0:8080 上启动服务
// 	r.Run(":8080")
// }

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
