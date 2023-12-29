package models

type User struct {
	ID     int            `json:"id" form:"id" gorm:"primaryKey"`
	Name   string         `json:"name" form:"name"`
	Locker LockerResponse `json:"locker"`
	Posts  []Post         `json:"posts" `
}

type UserResponse struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" from:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
