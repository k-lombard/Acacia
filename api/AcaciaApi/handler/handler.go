package handler

import (
	"bytes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/k-lombard/Acacia/AcaciaApi/database"
	"gopkg.in/olahol/melody.v1"
)

var dbInstance database.Database
var (
	router = gin.Default()
)

type routes struct {
	router *gin.Engine
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 400 {
		fmt.Println("Response body: " + blw.body.String())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RouteHandler(db database.Database, m *melody.Melody) *gin.Engine {
	dbInstance = db
	router.Use(ginBodyLogMiddleware)
	r := routes{
		router: gin.Default(),
	}
	// r.router.Use(CORSMiddleware())
	r.router.Use(cors.Default())
	api := r.router.Group("/api", CORSMiddleware())
	geolocationpositions := api.Group("/geolocationpositions")
	r.geolocationpositions(geolocationpositions)
	sentries := api.Group("/sentries")
	r.sentries(sentries)
	return r.router
}
