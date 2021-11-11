package resources

import (
	"time"

	"github.com/yomraliahmet/fiber-api/models"
)

type Order struct {
	ID        uint      `json:"id" extensions:"x-order=1" example:"1"`
	User      User      `json:"user" extensions:"x-order=2"`
	Product   Product   `json:"product" extensions:"x-order=3"`
	CreatedAt time.Time `json:"order_date" extensions:"x-order=4" example:"2021-11-08 22:38:24"`
}

func OrderResource(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}
