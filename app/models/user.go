package models

type User struct {
 ID          string    `json:"id"`
 Name       string    `json:"name"`
 Subname string    `json:"subname"`
 Subscriptions Subscription     `json:"subscriptions"`
 Status      string    `json:"status"`
}


