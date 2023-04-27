/**
 * @Author Nil
 * @Description model/init.go
 * @Date 2023/3/28 20:21
 **/

package model

import (
	"github.com/ha5ky/hu5ky-bot/model/base"
)

func Registry() {
	controller := NewController()
	TableRegister = append(TableRegister, base.TableRegister...)
	if err := controller.CreateTables(TableRegister); err != nil {
		panic(err)
	}
}

var (
	TableRegister = []base.AutoRegister{
		&HistoricMsg{},
	}
)
