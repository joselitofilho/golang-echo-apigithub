package resources

import (
	"github.com/joselitofilho/golang-echo-apigithub/internal/core"
	"github.com/joselitofilho/golang-echo-apigithub/internal/helpers"
	"gorm.io/gorm"
)

type AbstractResourceGormDB struct {
	dbProvider   *gorm.DB
	resourceName string
}

func (r *AbstractResourceGormDB) Create(value interface{}) error {
	tx := r.dbProvider.Create(value)
	return tx.Error
}

func (r *AbstractResourceGormDB) Get(entityID uint64, dest interface{}) error {
	tx := r.dbProvider.First(dest, uint(entityID))
	return tx.Error
}

func (r *AbstractResourceGormDB) List(dest interface{}, options *core.ListRequestOptions) (int64, error) {
	var count int64
	tx := r.dbProvider.Limit(options.Limit).Offset(options.Offset).Order("id ASC").Find(dest)
	r.dbProvider.Model(dest).Count(&count)
	return count, tx.Error
}

func (r *AbstractResourceGormDB) Update(newEntity interface{}, dest interface{}) error {
	newEntityMap, err := helpers.StructToMap(newEntity)
	if err != nil {
		return err
	}
	delete(newEntityMap, "Model")
	tx := r.dbProvider.Model(dest).Updates(newEntityMap)
	return tx.Error
}

func (r *AbstractResourceGormDB) Delete(entityID uint64, value interface{}) error {
	tx := r.dbProvider.Unscoped().Delete(value, uint(entityID))
	return tx.Error
}

func (r *AbstractResourceGormDB) SoftDelete(entityID uint64, value interface{}) error {
	tx := r.dbProvider.Delete(value, uint(entityID))
	return tx.Error
}
