package models

type Group struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Name  string  `json:"name"`
    Category string `json:"category"`
    Capacity int `json:"capacity"`
    Users []User  `json:"users" gorm:"many2many:user_groups;"`
}
