package models


type User struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Name          string         `json:"name"`
    Subname       string         `json:"subname"`
    Username string `json:"username" gorm:"uniqueIndex"`
    Email string `json:"email" gorm:"uniqueIndex"`
    Password string 
    Status        string         `json:"status"`
    Subscriptions []Subscription `json:"subscriptions" gorm:"many2many:user_subscriptions;"`
    Groups []Group `json:"groups" gorm:"many2many:user_groups;"`
}


