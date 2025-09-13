package in

type ProductDTOIn struct {
	ID          int64   `json:"id"`
	ShopID      int64   `json:"shop_id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
