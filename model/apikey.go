/**
 * @Author Nil
 * @Description model/apikey.go
 * @Date 2023/4/21 10:38
 **/

package model

import (
	"github.com/ha5ky/hu5ky-bot/model/base"
	"gorm.io/gorm"
)

func (c *Controller) ApiKeyModel(m *ApiKey) *ApiKey {
	m.controller = c.controller
	return m
}

type ApiKey struct {
	controller *gorm.DB
	gorm.Model
	Key      string
	SystemId uint
	//UserId   uint
}

func init() {
	c := new(ApiKey)
	c.Registry()
}

func (c *ApiKey) TableName() string {
	return "apikey"
}

func (c *ApiKey) Registry() {
	base.TableRegister = append(base.TableRegister, &ApiKey{})
}

type ApiKeyQuery struct {
	PageSize  *int
	PageIndex *int
	PreLoad   bool
	ID        *uint
	IDs       *[]uint

	Desc bool
}

func (c *ApiKey) Condition(q *ApiKeyQuery) *gorm.DB {
	if q.ID != nil {
		c.controller = c.controller.Where("id = ?", *q.ID)
	}
	if q.Desc {
		c.controller = c.controller.Order("created_at desc")
	}
	return c.controller
}
