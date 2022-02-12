package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todomvc-app-template-golang/db"
	"todomvc-app-template-golang/model"
)

func Add(c *gin.Context) {
	var p model.TodomvcAdd
	c.ShouldBindJSON(&p)
	fmt.Printf("ShouldBindJSON:%+v\n", p)

	var m = &model.Todomvc{Item: p.Item, Status: 0}
	db.DB.Create(&m)

	c.JSON(http.StatusOK, nil)
}

func Del(c *gin.Context) {
	var p model.TodomvcDel
	c.ShouldBindJSON(&p)
	fmt.Printf("ShouldBindJSON:%+v\n", p)

	db.DB.Where("id", p.Id).Delete(&model.Todomvc{})

	c.JSON(http.StatusOK, nil)

}

func Update(c *gin.Context) {
	var p []model.TodomvcUpdate
	c.ShouldBindJSON(&p)
	fmt.Printf("ShouldBindJSON:%+v\n", p)

	for _, t := range p {
		db.DB.Model(&model.Todomvc{}).Where("id", t.Id).Select("item", "status", "updated_at").Updates(model.Todomvc{
			Item:   t.Item,
			Status: t.Status,
		})
	}

	c.JSON(http.StatusOK, nil)

}

func Find(c *gin.Context) {
	var p model.TodomvcFind
	c.ShouldBindJSON(&p)
	fmt.Printf("ShouldBindJSON:%+v\n", p)

	var m []model.Todomvc
	var tx = db.DB

	if p.Status != -1 {
		tx = tx.Where("status", p.Status)
	}

	if p.Item != "" {
		tx = tx.Where("item LIKE ?", "%"+p.Item+"%")
	}

	tx.Find(&m)

	c.JSON(http.StatusOK, &m)
}
