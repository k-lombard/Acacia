package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/k-lombard/Acacia/AcaciaApi/database"
	"github.com/k-lombard/Acacia/AcaciaApi/models"
)

func (r routes) sentries(rg *gin.RouterGroup) {
	rg.GET("/", getAllSentries)
	rg.POST("/", addSentry)
	rg.GET("/:id", getSentryById)
	rg.PUT("/:id", updateSentry)
	rg.DELETE("/:id", deleteSentry)
}

func addSentry(c *gin.Context) {
	sentry := &models.Sentry{}
	r := c.Request
	if err := render.Bind(r, sentry); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	sentryOut, err := dbInstance.AddSentry(sentry)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, sentryOut)
}

func getAllSentries(c *gin.Context) {
	sentries, err := dbInstance.GetAllSentries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, sentries)
}

func getSentryById(c *gin.Context) {
	sentryID := uuid.MustParse(c.Param("id"))
	sentry, err := dbInstance.GetSentryById(sentryID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, sentry)
}

func deleteSentry(c *gin.Context) {
	sentryId := uuid.MustParse(c.Param("id"))
	err := dbInstance.DeleteSentry(sentryId)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, sentryId)
}
func updateSentry(c *gin.Context) {
	r := c.Request
	sentryId := uuid.MustParse(c.Param("id"))
	sentryData := models.Sentry{}
	if err := render.Bind(r, &sentryData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	sentry, err := dbInstance.UpdateSentry(sentryId, sentryData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, sentry)
}
