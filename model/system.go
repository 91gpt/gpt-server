/**
 * @Author Nil
 * @Description model/system.go
 * @Date 2023/4/21 14:04
 **/

package model

import (
	"github.com/ha5ky/hu5ky-bot/model/base"
	"gorm.io/gorm"
)

func (c *Controller) SystemModel(m *System) *System {
	m.controller = c.controller
	return m
}

type System struct {
	controller *gorm.DB
	gorm.Model

	ApiKeys []ApiKey `gorm:"foreignKey:SystemId"`
}

func init() {
	c := new(System)
	c.Registry()
}

func (c *System) TableName() string {
	return "system"
}

func (c *System) Registry() {
	base.TableRegister = append(base.TableRegister, &System{})
}

type SystemQuery struct {
	PageSize  *int
	PageIndex *int
	PreLoad   bool
	ID        *uint
	UUID      *uint
	IDs       *[]uint
}

func (c *System) Condition(q *SystemQuery) *gorm.DB {
	if q.ID != nil {
		c.controller = c.controller.Where("id = ?", *q.ID)
	}
	if q.UUID != nil {
		c.controller = c.controller.Where("uuid = ?", *q.UUID)
	}
	return c.controller
}
