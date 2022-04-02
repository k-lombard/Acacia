package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/k-lombard/Acacia/AcaciaApi/database"
	"github.com/k-lombard/Acacia/AcaciaApi/models"
)

func (r routes) images(rg *gin.RouterGroup) {
	rg.GET("/", getAllImages)
	rg.POST("/", addImage)
	rg.GET("/:id", getImageById)
	rg.GET("/sentry/:sentryid", getImagesBySentryId)
	rg.DELETE("/:id", deleteImage)
}

func addImage(c *gin.Context) {
	image := &models.Image{}
	r := c.Request
	if err := render.Bind(r, image); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	imageOut, err := dbInstance.AddImage(image)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, imageOut)
}

func getAllImages(c *gin.Context) {
	images, err := dbInstance.GetAllImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, images)
}

func getImageById(c *gin.Context) {
	imageID := uuid.MustParse(c.Param("id"))
	image, err := dbInstance.GetImageById(imageID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, image)
}

func deleteImage(c *gin.Context) {
	imageId := uuid.MustParse(c.Param("id"))
	err := dbInstance.DeleteImage(imageId)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, imageId)
}

func getImagesBySentryId(c *gin.Context) {
	sentryID := uuid.MustParse(c.Param("sentryid"))
	images, err := dbInstance.GetImagesBySentryId(sentryID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, images)
}
