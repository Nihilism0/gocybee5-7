package model

type Store struct {
	Storename string `form:"storename" json:"storename" binding:"required"`
}
