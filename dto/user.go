package dto

import "time"

type InsertUser struct {
	Username string `json:"username" binding:"required,max=15,min=5"`
	Password string `json:"password" binding:"required,max=15,min=5"`
	Status   int    `json:"status" default:"0"`
	RoleId   int    `json:"role_id" default:"0"`
}

type FindOneByUser struct {
	Username string `json:"username" binding:"required,max=15,min=5"`
	Password string `json:"password" binding:"required,max=15,min=5"`
	Status   int    `json:"status"`
}

type FindOneByUsername struct {
	Username string `json:"username" binding:"required,max=15,min=5"`
	Status   int    `json:"status" binding:"default=0"`
}

type User struct {
	Username       string    `json:"username" binding:"required,max=15,min=5"`
	Status         int       `json:"status" binding:"default=0"`
	RoleId         int       `json:"role_id" binding:"default=0"`
	CreatedTime    time.Time `json:"create_time"`
	LastUpdateTime time.Time `json:"last_update_time"`
}

type UserLogin struct {
	Token string `json:"token"`
}
