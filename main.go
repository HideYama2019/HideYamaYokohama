package main

import (
    "time"
    "github.com/gin-gonic/gin"
)

func main() {
    now := time.Now().Format("2006-01-02_15:04:05")

    xx := gin.Default()
    xx.Static("static", "./static")
    xx.LoadHTMLGlob("./views/*.html")

    xx.GET("/", func(cc *gin.Context) {
        cc.HTML(200, "index.html", gin.H{
            "now": now,
        })
    })

    xx.GET("/hello", func(cc *gin.Context) {
        cc.HTML(200, "hello.html", gin.H{

        })
    })

    xx.Run(
        ":8080",
    )
}
