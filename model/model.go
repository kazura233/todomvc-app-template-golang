package model

import "gorm.io/gorm"

const (
	TODOMVC_STATUS_ACTIVE    = 0
	TODOMVC_STATUS_COMPLETED = 1
)

type Todomvc struct {
	gorm.Model
	Item   string
	Status uint
}

type TodomvcAdd struct {
	Item string `json:"item"`
}

type TodomvcDel struct {
	Id uint `json:"id"`
}

type TodomvcUpdate struct {
	Id     uint   `json:"id"`
	Item   string `json:"item"`
	Status uint   `json:"status"`
}

type TodomvcFind struct {
	Item   string `json:"item"`
	Status int    `json:"status"` // -1 为全部
}
