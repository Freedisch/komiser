package controller

import (
	"context"
	"database/sql"
)

type totalOutput struct {
	Total int `bun:"total" json:"total"`
}

type costOutput struct {
	Total float64 `bun:"sum" json:"total"`
}

type regionOutput struct {
	Region string `bun:"region" json:"region"`
}

type providerOutput struct {
	Provider string `bun:"provider" json:"provider"`
}

type serviceOutput struct {
	Service string `bun:"service" json:"service"`
}

type accountOutput struct {
	Account string `bun:"account" json:"account"`
}

type Repository interface {
	HandleQuery(context.Context, string, interface{}, [][3]string) (sql.Result, error)
}

type Controller struct {
	repo Repository
}

func New(repo Repository) *Controller {
	return &Controller{repo}
}
