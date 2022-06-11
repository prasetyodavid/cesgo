package users

import (
	models "cesgo/models/users"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"cesgo/models/response"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) (err error) {

	user := new(models.Users)
	c.Bind(user)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
		file, err := c.FormFile("ktp")
		if err != nil {
			fmt.Println("Ktp kosong")
		} else {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			dst, err := os.Create(file.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			user.Ktp = file.Filename
			fmt.Println("Ada file, akan disimpan")
		}
	}
	response := new(response.General)
	if user.CreateUser() != nil { // method create user
		response.ErrorCode = 10
		response.Message = "Gagal create data user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create data user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)
}

func SearchUser(c echo.Context) (err error) {
	response := new(response.General)
	users, err := models.GetAll(c.QueryParam("keywords")) // method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Failed"
	} else {
		response.Message = "Success"
		response.Data = users
	}
	return c.JSON(http.StatusOK, response)
}
