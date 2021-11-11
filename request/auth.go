package request

type Auth struct {
	User string `json:"user" extensions:"x-order=1" example:"john"`
	Pass string `json:"pass" extensions:"x-order=2" example:"doe"`
}
