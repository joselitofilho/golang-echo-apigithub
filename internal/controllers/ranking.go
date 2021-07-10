package controllers

import (
	"net/http"
	"strconv"

	"github.com/joselitofilho/golang-echo-apigithub/internal/core"
	"github.com/joselitofilho/golang-echo-apigithub/internal/models"
	"github.com/joselitofilho/golang-echo-apigithub/internal/resources"
	"github.com/labstack/echo/v4"
)

type RankingController struct {
	resource resources.RankingResource
}

func NewRankingController(resource resources.RankingResource) *RankingController {
	return &RankingController{resource}
}

func (c *RankingController) Get(ctx echo.Context) error {
	strId := ctx.Param("id")

	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	ranking := models.Ranking{}
	if err := c.resource.Get(id, &ranking); err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	if ranking.ID == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, ranking)
}

func (c *RankingController) List(ctx echo.Context) error {
	rankings := []models.Ranking{}
	if _, err := c.resource.List(&rankings, &core.ListRequestOptions{Offset: 0, Limit: 100}); err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, rankings)
}
