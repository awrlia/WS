package app

import "1.ProjectGo/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
		{Model: models.Order{}},
		{Model: models.OrderCustomer{}},
		{Model: models.OrderItem{}},
		{Model: models.Pembayaran{}},
		{Model: models.Shipment{}},
	}
}
