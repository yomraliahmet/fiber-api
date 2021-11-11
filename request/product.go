package request

type Product struct {
	Name         string `json:"name" extensions:"x-order=1"`
	SerialNumber string `json:"serial_number" extensions:"x-order=2"`
}
