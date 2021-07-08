package models

import (
	"gorm.io/gorm"
)

type Ranking struct {
	gorm.Model
	AvatarURL         string
	Name              string
	StargazersCount   int
	FollowersCount    int
	ProjectsCount     int
	ContributtedCount int
	Languages         []Language
}

type Language struct {
	gorm.Model
	RankingID uint
	Name      string
}
