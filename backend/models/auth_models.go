package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Workspace struct {
	gorm.Model
	Name     string `json:"name"`
	JoinCode string `json:"join_code"`
	UserID   uint
}

type Channel struct {
	gorm.Model
	Name        string `json:"name"`
	WorkspaceID uint   `json:"workspace_id"`
}

type Message struct {
	gorm.Model
	Body        string `json:"body"`
	ChannelID   uint   `json:"channel_id"`
	WorkspaceID uint   `json:"workspace_id"`
	UserID      uint   `json:"user_id"`
}
