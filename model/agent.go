/**
 * @Author Nil
 * @Description model/agent.go
 * @Date 2023/4/21 15:35
 **/

package model

import (
	"errors"
	"github.com/ha5ky/hu5ky-bot/model/base"
	"gorm.io/gorm"
)

func (c *Controller) AgentModel(m *Agent) *Agent {
	m.controller = c.controller
	return m
}

type Agent struct {
	controller *gorm.DB
	gorm.Model
	Users []User `gorm:"foreignKey:AgentId"`
	PID   uint

	ImageN           int
	ImageSize        string
	SubAccountNumber int
	Daily            int
	Level            int
}

func init() {
	c := new(Agent)
	c.Registry()
}

func (c *Agent) TableName() string {
	return "agent"
}

func (c *Agent) Registry() {
	base.TableRegister = append(base.TableRegister, &Agent{})
}

type AgentQuery struct {
	PageSize  *int
	PageIndex *int
	PreLoad   bool
	ID        *uint
	IDs       *[]uint

	Desc bool
}

func (c *Agent) Condition(q *AgentQuery) *gorm.DB {
	if q.ID != nil {
		c.controller = c.controller.Where("id = ?", *q.ID)
	}
	if q.Desc {
		c.controller = c.controller.Order("created_at desc")
	}
	if q.PreLoad {
		c.controller = c.controller.Preload("Users")
	}
	return c.controller
}

func (c *Agent) Get(q *AgentQuery) (res Agent, err error) {
	if errors.Is(c.Condition(q).Last(&res).Error, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	return
}

func (c *Agent) Save() error {
	return c.controller.Save(c).Error
}

func (c *Agent) List(q *AgentQuery) (res []*Agent, total int64, err error) {
	if err = c.Condition(q).Find(&res).Count(&total).Error; err != nil {
		return
	}
	if q.PageIndex != nil {
		c.controller = c.controller.Offset((*q.PageIndex - 1) * *q.PageSize)
	}
	if q.PageSize != nil {
		c.controller = c.controller.Limit(*q.PageSize)
	}
	if err = c.Condition(q).Find(&res).Error; err != nil {
		return
	}
	return
}
