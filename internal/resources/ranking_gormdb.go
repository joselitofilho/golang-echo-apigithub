package resources

import "gorm.io/gorm"

const (
	RankingsResourceName = "rankings"
)

type RankingResourceGormDB struct {
	AbstractResourceGormDB
}

func NewRankingResourceGormDB(dbProvider *gorm.DB) *RankingResourceGormDB {
	return &RankingResourceGormDB{
		AbstractResourceGormDB: AbstractResourceGormDB{
			dbProvider:   dbProvider,
			resourceName: RankingsResourceName,
		},
	}
}
