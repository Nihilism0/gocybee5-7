package model

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type Changepassword struct {
	Username    string `form:"username" json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	NewPassword string `form:"newpassword" json:"newpassword" binding:"required"`
}

type Findpassword struct {
	Username string `form:"username" json:"username" binding:"required"`
}
