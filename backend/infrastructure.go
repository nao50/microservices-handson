package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

type DBConnection struct {
	DBConnection *gorm.DB
}

//
func NewTODORepository(conn *gorm.DB) TODORepository {
	return &DBConnection{DBConnection: conn}
}

//
func (dc *DBConnection) CreateTODO(ctx context.Context, todo *TODOModel) (*TODOModel, error) {
	if err := dc.DBConnection.Create(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

//
func (dc *DBConnection) ListTODOs(ctx context.Context) ([]*TODOModel, error) {
	var todosModel []*TODOModel
	dc.DBConnection.Find(&todosModel)
	return todosModel, nil
}

//
func (dc *DBConnection) GetTODO(ctx context.Context, id string) (*TODOModel, error) {
	todo := &TODOModel{ID: id}
	if err := dc.DBConnection.First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

//
func (dc *DBConnection) UpdateTODO(ctx context.Context, id string, todo *TODOModel) (*TODOModel, int, error) {
	if id != todo.ID {
		return nil, 400, nil
	}
	if err := dc.DBConnection.Model(&todo).Update(&todo).Error; err != nil {
		return nil, 500, err
	}
	return todo, 200, nil
}

//
func (dc *DBConnection) DeleteTODO(ctx context.Context, id string) (int, error) {
	// if id != todo.ID {
	// 	return nil, 400, nil
	// }
	// if err := dc.DBConnection.Delete(&todo).Error; err != nil {
	// 	return 500, err
	// }
	// return 200, nil
	todo := &TODOModel{ID: id}

	if err := dc.DBConnection.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return 500, err
	}
	return 204, nil
}
