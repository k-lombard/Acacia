package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/k-lombard/Acacia/AcaciaApi/database"
	"github.com/k-lombard/Acacia/AcaciaApi/models"
)

func (r routes) geolocationpositions(rg *gin.RouterGroup) {
	rg.GET("/", getAllGeolocationPositions)
	rg.POST("/", addGeolocationPosition)
	rg.GET("/:id", getGeolocationPositionBySentryId)
	rg.PUT("/:id", updateGeolocationPosition)
	rg.DELETE("/:id", deleteGeolocationPosition)
}

func addGeolocationPosition(c *gin.Context) {
	loc := &models.GeolocationPosition{}
	r := c.Request
	if err := render.Bind(r, loc); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	geolocationPositionOut, err := dbInstance.AddGeolocationPosition(loc)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, geolocationPositionOut)
}

func getAllGeolocationPositions(c *gin.Context) {
	locs, err := dbInstance.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, locs)
}

func getGeolocationPositionBySentryId(c *gin.Context) {
	sentryID := uuid.MustParse(c.Param("id"))
	loc, err := dbInstance.GetGeolocationPositionBySentryId(sentryID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, loc)
}

func deleteGeolocationPosition(c *gin.Context) {
	sentryId := uuid.MustParse(c.Param("id"))
	err := dbInstance.DeleteGeolocationPosition(sentryId)
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
func updateGeolocationPosition(c *gin.Context) {
	r := c.Request
	sentryId := uuid.MustParse(c.Param("id"))
	locData := models.GeolocationPosition{}
	if err := render.Bind(r, &locData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	loc, err := dbInstance.UpdateGeolocationPosition(sentryId, locData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, loc)
}
