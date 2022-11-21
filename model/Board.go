package model

type Board struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Board    string `form:"board" json:"board" binding:"required"`
}

type RealBoard struct {
	ID       int
	Username string
	Board    string
}
