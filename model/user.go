/**
 * @Author Nil
 * @Description model/user.go
 * @Date 2023/4/19 20:35
 **/

package model

import (
	"errors"
	"github.com/ha5ky/hu5ky-bot/model/base"
	"github.com/ha5ky/hu5ky-bot/pkg/util"
	"gorm.io/gorm"
)

func (c *Controller) UserModel(m *User) *User {
	m.controller = c.controller
	return m
}

type User struct {
	controller *gorm.DB
	gorm.Model
	Remainder   int
	Daily       int
	Active      bool
	AgentId     uint
	PID         uint
	Users       []User `gorm:"foreignKey:PID"`
	UUID        string
	Account     string
	Pwd         string
	Name        string
	Description string
	HistoricMsg []HistoricMsg `gorm:"foreignKey:UserId"`
	//ApiKeys     []ApiKey      `gorm:"foreignKey:UserId"`
}

func init() {
	c := new(User)
	c.Registry()
}

func (c *User) TableName() string {
	return "user"
}

func (c *User) Registry() {
	base.TableRegister = append(base.TableRegister, &User{})
}

type UserQuery struct {
	PageSize  *int
	PageIndex *int
	PreLoad   bool
	ID        *uint
	UUID      *string
	IDs       *[]uint

	Account *string
	Pwd     *string

	Name    *string
	AgentId *uint
}

func (c *User) Condition(q *UserQuery) *gorm.DB {
	if q.ID != nil {
		c.controller = c.controller.Where("id = ?", *q.ID)
	}
	if q.UUID != nil {
		c.controller = c.controller.Where("uuid = ?", *q.UUID)
	}
	if q.AgentId != nil {
		c.controller = c.controller.Where("agent_id = ?", *q.AgentId)
	}
	if q.Name != nil {
		c.controller = c.controller.Where("agent_id = ?", *q.AgentId)
	}
	if q.Account != nil {
		c.controller = c.controller.Where("account = ?", *q.Account)
	}
	if q.Pwd != nil {
		c.controller = c.controller.Where("pwd = ?", *q.Pwd)
	}
	return c.controller
}

func (c *User) Check(q *UserQuery) (uid uint, ok bool, err error) {
	var res User
	if errors.Is(c.Condition(q).Last(&res).Error, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	uid = res.ID
	ok = true
	return
}

func (c *User) Get(q *UserQuery) (res User, err error) {
	if errors.Is(c.Condition(q).Last(&res).Error, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	return
}

func (c *User) Activate(q *UserQuery) (err error) {
	res := new(User)
	if errors.Is(c.Condition(q).Last(&res).Error, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	res.Active = true
	return c.controller.Save(res).Error
}

func (c *User) Generate(maintainer uint, agent *Agent) (err error) {
	users := make([]*User, 0)
	mainUser := &User{
		Remainder:   agent.Daily,
		Daily:       agent.Daily,
		AgentId:     agent.ID,
		Active:      false,
		PID:         maintainer,
		UUID:        util.GetUUID(),
		Account:     "U-" + util.GetUUID(),
		Pwd:         "P-" + util.GetUUID(),
		Description: "",
	}
	if err = c.controller.Save(mainUser).Error; err != nil {
		return
	}
	for i := 0; i < agent.SubAccountNumber; i++ {
		users = append(users, &User{
			Remainder:   agent.Daily,
			Daily:       agent.Daily,
			AgentId:     agent.ID,
			Active:      false,
			PID:         mainUser.ID,
			UUID:        util.GetUUID(),
			Account:     "U-" + util.GetUUID(),
			Pwd:         "P-" + util.GetUUID(),
			Description: "",
		})
	}
	return c.controller.Save(users).Error
}
