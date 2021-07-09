package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/joselitofilho/golang-echo-apigithub/internal/api"
	"github.com/joselitofilho/golang-echo-apigithub/internal/controllers"
	"github.com/joselitofilho/golang-echo-apigithub/internal/models"
	"github.com/joselitofilho/golang-echo-apigithub/internal/resources"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	databaseDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
	gormDB, err := gorm.Open(postgres.Open(databaseDetails), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := gormDB.AutoMigrate(&models.Ranking{}, &models.Language{}); err != nil {
		panic(err)
	}

	rankingResource := resources.NewRankingResourceGormDB(gormDB)
	rankingController := controllers.NewRankingController(rankingResource)
	e.GET("/rankings", rankingController.List)
	e.GET("/rankings/:id", rankingController.Get)

	ghLoopCh := make(chan struct{})
	ghClient := github.NewClient(nil)
	org := "Instituto-Atlantico"
	// org := "github"
	go api.GHInfos(ghLoopCh, ghClient, org, rankingResource)

	e.Logger.Fatal(e.Start(":1323"))
}
