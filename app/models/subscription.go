package models

type Subscription struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Name  string  `json:"name"`
    Category string `json:"category"`
    Amount  float32 `json:"amount"`
    Capacity int `json:"capacity"`
    IsPersonal bool `json:"ispersonal"`
    Users []User  `json:"users" gorm:"many2many:user_subscriptions;"`
}

