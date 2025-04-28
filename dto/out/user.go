package out

type UserResponse struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	IsAdmin      bool   `json:"is_admin"`
	CustomerName string `json:"customer_name"`
}
