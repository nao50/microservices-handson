package main

import "context"

// TODOModel is
type TODOModel struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	UID  string `json:"uid" validate:"required"`
	Done bool   `json:"done" validate:"required"`
}

// TODORepository is
type TODORepository interface {
	CreateTODO(ctx context.Context, todo *TODOModel) (*TODOModel, error)
	ListTODOs(ctx context.Context) ([]*TODOModel, error)
	GetTODO(ctx context.Context, id string) (*TODOModel, error)
	UpdateTODO(ctx context.Context, id string, todo *TODOModel) (*TODOModel, int, error)
	DeleteTODO(ctx context.Context, id string) (int, error)
}
