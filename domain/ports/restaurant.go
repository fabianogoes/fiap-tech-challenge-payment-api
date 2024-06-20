package ports

type RestaurantClientPort interface {
	Webhook(orderID uint, status string) error
}
