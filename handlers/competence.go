package handlers

import (
	"database/sql"
	"net/http"

	"github.com/saguywalker/sit-competence/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetCompetencies(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetCompetencies(db))
	}
}

func PostCompetence(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var competence models.Competence
		c.Bind(&competence)
		id, err := models.PostCompetence(db, competence)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusCreated, H{
				"inserted": id,
			})
		}
	}
}

func DeleteCompetence(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := c.Param("id")
		_, err := models.DeleteCompetence(db, id)
		if err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		}
	}

}
