package request

type Order struct {
	UserId    int `json:"user_id" extensions:"x-order=1"`
	ProductId int `json:"product_id" extensions:"x-order=2"`
}
