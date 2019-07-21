package handlers

import (
	"database/sql"
	"net/http"

	"github.com/saguywalker/sit-competence/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetStudents(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetStudents(db))
	}
}

func PostStudent(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var student models.Student
		c.Bind(&student)
		id, err := models.PostStudent(db, student)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusCreated, H{
				"inserted": id,
			})
		}
	}
}

func DeleteStudent(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := c.Param("id")
		_, err := models.DeleteStudent(db, id)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		}
	}

}
