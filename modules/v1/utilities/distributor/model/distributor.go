package model

import "time"

type LoginMessage struct {
	Success string `json:"success"`
	Reason  string `json:"reason"`
}

type SubmitMessage struct {
	Success string `json:"success"`
	EDisc   string `json:"E_DISC"`
	Reason  string `json:"reason"`
}

type LoginInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SubmitInput struct {
	DiscountType string    `form:"discount_type"`
	DiscountName string    `form:"discount_name"`
	Description  string    `form:"description"`
	DBerlaku     time.Time `form:"D_BERLAKU"`
	DExpire      time.Time `form:"D_EXPIRE"`
	CNosup       int       `form:"C_NOSPUP"`
	CKddivpri    int       `form:"C_KDDIVPRO"`
	CCusno       string    `form:"C_CUSNO"`
}
