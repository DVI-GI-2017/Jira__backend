package request

type LoginRequest struct {
	User     string `json:"u"`
	Password string `json:"p"`
}