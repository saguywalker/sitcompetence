package handlers

import (
	"database/sql"
	"net/http"

	"github.com/sit-competence/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetStaffs(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetStaffs(db))
	}
}

func PostStaff(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var staff models.Staff
		c.Bind(&staff)
		id, err := models.PostStaff(db, staff)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusCreated, H{
				"inserted": id,
			})
		}
	}
}

func DeleteStaff(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := c.Param("id")
		_, err := models.DeleteStaff(db, id)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		}
	}

}
