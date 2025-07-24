package models

type Subscription struct {
 ID          string    `json:"id"`
 Name       string     `json:"name"`
 Cost 		float32     `json:"cost"`
}