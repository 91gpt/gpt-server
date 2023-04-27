/**
 * @Author Nil
 * @Description model/historic_msg.go
 * @Date 2023/4/21 10:29
 **/

package model

import (
	"errors"
	"gorm.io/gorm"
)

func (c *Controller) HistoricMsgModel(m *HistoricMsg) *HistoricMsg {
	m.controller = c.controller
	return m
}

type HistoricMsg struct {
	controller *gorm.DB
	gorm.Model
	UserId   uint
	Title    string
	Messages []Message `gorm:"foreignKey:HistoricMsgId"`
}

func init() {
	c := new(HistoricMsg)
	c.Registry()
}

func (c *HistoricMsg) TableName() string {
	return "historic_msg"
}

func (c *HistoricMsg) Registry() {
	//base.TableRegister = append(base.TableRegister, &HistoricMsg{})
}

type HistoricMsgQuery struct {
	PageSize  *int
	PageIndex *int
	PreLoad   bool
	ID        *uint
	IDs       *[]uint

	UserId *uint
	Desc   bool
}

func (c *HistoricMsg) Condition(q *HistoricMsgQuery) *gorm.DB {
	if q.ID != nil {
		c.controller = c.controller.Where("id = ?", *q.ID)
	}
	if q.UserId != nil {
		c.controller = c.controller.Where("user_id = ?", *q.UserId)
	}
	if q.Desc {
		c.controller = c.controller.Order("created_at desc")
	}
	if q.PreLoad {
		c.controller = c.controller.Preload("Messages")
	}
	return c.controller
}

func (c *HistoricMsg) Save() (err error) {
	return c.controller.Save(c).Error
}

func (c *HistoricMsg) Get(q *HistoricMsgQuery) (res HistoricMsg, err error) {
	if errors.Is(c.Condition(q).Last(&res).Error, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	return
}

func (c *HistoricMsg) List(q *HistoricMsgQuery) (res []*HistoricMsg, total int64, err error) {
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
