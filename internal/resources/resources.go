package resources

import "github.com/joselitofilho/golang-echo-apigithub/internal/core"

type BaseResource interface {
	Create(interface{}) error
	Get(uint64, interface{}) error
	List(interface{}, *core.ListRequestOptions) (int64, error)
	Update(newEntity interface{}, dest interface{}) error
	Delete(uint64, interface{}) error
	SoftDelete(uint64, interface{}) error
}
