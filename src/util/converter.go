package util

import (
	"omnichannel-backend-service/src/datastruct"
	"omnichannel-backend-service/src/model"
)

func ConvertDatastructProductToModelProduct(product *datastruct.Product) *model.Product {
	return &model.Product{
		ID:                 product.ID,
		Name:               product.Name,
		Price:              product.Price,
		Weight:             product.Weight,
		Stock:              product.Stock,
		Image:              product.Image,
		Description:        product.Description,
		TokopediaProductID: product.TokopediaProductID,
		ShopeeProductID:    product.ShopeeProductID,
	}
}

func ConvertProductToProductMessage(product *datastruct.Product, method datastruct.Method) *datastruct.ProductMessage {
	productMessage := &datastruct.ProductMessage{
		Method:             method,
		ID:                 product.ID,
		Name:               product.Name,
		Price:              product.Price,
		Weight:             product.Weight,
		Stock:              product.Stock,
		Image:              product.Image,
		Description:        product.Description,
		TokopediaProductID: product.TokopediaProductID,
		ShopeeProductID:    product.ShopeeProductID,
	}

	return productMessage
}

func ConvertCreateOrderRequestToOrderDatastruct(order *model.CreateOrderRequest, ID string, userID string) *datastruct.Order {
	status := []string{"RECEIVED", "ACCEPTED", "CANCELLED", "DONE"}
	orderData := &datastruct.Order{
		ID:                 ID,
		UserID:             userID,
		TotalPrice:         order.TotalPrice,
		TokopediaOrderID:   order.TokopediaOrderID,
		ShopeeOrderID:      order.ShopeeOrderID,
		CustomerName:       order.Customer.CustomerName,
		CustomerPhone:      order.Customer.CustomerPhone,
		CustomerAddress:    order.Customer.CustomerAddress,
		CustomerDistrict:   order.Customer.CustomerDistrict,
		CustomerCity:       order.Customer.CustomerCity,
		CustomerProvince:   order.Customer.CustomerProvince,
		CustomerCountry:    order.Customer.CustomerCountry,
		CustomerPostalCode: order.Customer.CustomerPostalCode,
		OrderStatus:        status[order.OrderStatus],
	}

	return orderData
}

func ConvertOrderAndOrderProductToModelOrder(order *datastruct.Order, orderProducts *datastruct.GetOrderProductResponse) *model.GetOrderResponse {
	customer := model.Customer{
		CustomerName:       order.CustomerName,
		CustomerPhone:      order.CustomerPhone,
		CustomerAddress:    order.CustomerAddress,
		CustomerDistrict:   order.CustomerDistrict,
		CustomerCity:       order.CustomerCity,
		CustomerProvince:   order.CustomerProvince,
		CustomerCountry:    order.CustomerCountry,
		CustomerPostalCode: order.CustomerPostalCode,
	}

	modelOrderProducts := ConvertDatastructOrderProductsToModelOderProducts(orderProducts)
	modelOrder := &model.GetOrderResponse{
		ID:               order.ID,
		TotalPrice:       order.TotalPrice,
		TokopediaOrderID: order.TokopediaOrderID,
		ShopeeOrderID:    order.ShopeeOrderID,
		Customer:         customer,
		OrderStatus:      order.OrderStatus,
		Products:         modelOrderProducts,
	}

	return modelOrder
}

func ConvertDatastructOrderProductsToModelOderProducts(orderProducts *datastruct.GetOrderProductResponse) []model.OrderProduct {
	var result []model.OrderProduct
	for _, orderProduct := range orderProducts.OrderProducts {
		modelOrderProduct := model.OrderProduct(*orderProduct)
		result = append(result, modelOrderProduct)
	}

	return result
}
