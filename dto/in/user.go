package in

type UserRequest struct {
	AbstractRequest
	ID         int64  `json:"id"`
	Username   string `json:"username" validate:"required,username"`
	Password   string `json:"password" validate:"required"`
	IsAdmin    bool   `json:"is_admin"`
	CustomerId int64  `json:"customer_id"`
}
