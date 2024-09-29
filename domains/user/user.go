package user

import "time"

type User struct {
	Id        string    `json:"id"`
	Role      string    `json:"role"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
