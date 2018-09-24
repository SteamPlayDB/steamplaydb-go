package Controllers

import (
	"../Models"
	"github.com/labstack/echo"
	"net/http"
)

func Games(c echo.Context) error {
	o := new(Models.GameFilterOptionsModel)
	if err := c.Bind(o); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	results, err := Models.FetchGamesByOptions(*o)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, results)
}
