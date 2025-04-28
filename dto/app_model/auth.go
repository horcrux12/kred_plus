package app_model

type UserSession struct {
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	IsAdmin    bool   `json:"is_admin"`
	CustomerId *int64 `json:"customer_id"`
}
