package handlers

import (
	"asira_borrower/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ClientImageFile(c echo.Context) error {
	defer c.Request().Body.Close()
	imageModel := models.Image{}

	imageId, _ := strconv.Atoi(c.Param("file_id"))
	image, err := imageModel.FindbyID(imageId)
	if err != nil {
		return returnInvalidResponse(http.StatusForbidden, err, "unauthorized")
	}
	log.Println(image.Image_string)
	return c.JSON(http.StatusOK, image)
}
