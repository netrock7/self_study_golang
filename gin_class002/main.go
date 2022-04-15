package main

import "github.com/gin-gonic/gin"

type PostParams struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"`
}

func main() {
	r := gin.Default()
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "出错了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})

	r.Run(":8899")
}
