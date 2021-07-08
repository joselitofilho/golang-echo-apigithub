package controllers

import (
	"net/http"

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
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, "Hello, World! "+id)
}

func (c *RankingController) List(ctx echo.Context) error {
	rankings := []models.Ranking{}
	if _, err := c.resource.List(&rankings, &core.ListRequestOptions{Offset: 0, Limit: 100}); err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, rankings)
}
