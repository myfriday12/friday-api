package main

import (
	"fmt"
	fridayroute "friday-api/api/friday-routes"
	friday "friday-api/friday/config"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Friday API")
	friday.AppConfigPath = "../friday/config"

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[FridayAPI] - [%s] | %s | %s | %s | %d | %s | %s | %s |\n",
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	router.Use(func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Add("Access-Control-Max-Age", "300")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		context.Writer.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept,tpa,ses")
		context.Writer.Header().Set("Content-Type", "application/json; charset=latin-1")
		context.Next()
	})

	v1 := router.Group("/v1")
	fridayroute.MountRoute(v1)

	port, hasValue := os.LookupEnv("FRIDAY_PORT")
	if hasValue == false {
		log.Fatal("Unable to load ENV FRIDAY PORT")
	}

	router.Run(":" + port)
}
