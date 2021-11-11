package request

type User struct {
	FirstName string `json:"first_name" extensions:"x-order=1"`
	LastName  string `json:"last_name" extensions:"x-order=2"`
	City      string `json:"city" extensions:"x-order=3"`
}
