package dao

import (
	"fmt"
	"strings"
	"zeus/pkg/api/dto"
	"zeus/pkg/api/model"

	"github.com/jinzhu/gorm"
)

type Menu struct {
}

// List
func (m Menu) List(treeDto dto.GeneralTreeDto) ([]model.Menu, int64) {
	var menus []model.Menu
	var total int64
	db := GetDb()
	// todo: data permission control
	for sk, sv := range dto.TransformSearch(treeDto.Q, dto.MenuListSearchMapping) {
		db = db.Where(fmt.Sprintf("%s = ?", sk), sv)
	}
	db.Preload("Domain").Order("name asc").Find(&menus)
	db.Model(&model.Menu{}).Count(&total)
	return menus, total
}

// GetMenusByIds
func (m Menu) GetMenusByIds(ids string) []model.Menu {
	var menus []model.Menu
	db := GetDb()
	db.Where("id in (?) and menu_type=1", strings.Split(ids, ",")).Find(&menus)
	return menus
}

// GetMenusPermByIds - get permissions in menu table
func (m Menu) GetMenusPermByIds(ids string) []model.Menu {
	var menus []model.Menu
	db := GetDb()
	db.Where("id in (?) and menu_type=2", strings.Split(ids, ",")).Find(&menus)
	return menus
}

// GetByAlias - get row by alias
func (m Menu) GetByAlias(alias string) model.Menu {
	var menu model.Menu
	db := GetDb()
	db.Where("alias = ? and menu_type=2", alias).First(&menu)
	return menu
}

//Get - get single menu info
func (m Menu) Get(id int, preload bool) model.Menu {
	var menu model.Menu
	db := GetDb()
	if preload {
		db = db.Preload("Domain")
	}
	db.Where("id = ?", id).First(&menu)
	return menu
}

//GetSubMenus
func (m Menu) GetSubMenus(id int) []model.Menu {
	var menus []model.Menu
	db := GetDb()
	db.Where("parent_id=?", id).First(&menus)
	return menus
}

// Create - new menu
func (m Menu) Create(menu *model.Menu) *gorm.DB {
	db := GetDb()
	return db.Save(menu)
}

// Update - update menu
func (m Menu) Update(menu *model.Menu, ups map[string]interface{}) *gorm.DB {
	db := GetDb()
	return db.Model(menu).Update(ups)
}

// Delete - delete menu
func (m Menu) Delete(menu *model.Menu) *gorm.DB {
	db := GetDb()
	return db.Delete(menu)
}
