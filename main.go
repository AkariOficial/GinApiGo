package main

import (
	"github.com/AkariOficial/GinApiGo/ola"
	"github.com/gin-gonic/gin"
)

func main() {

    ola.Ola()

    r := gin.Default()
    r.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "Frutas": "Banana",
            "Jogos": "Genshin Impact",
        })
    })
    r.Run(":6765")
}
