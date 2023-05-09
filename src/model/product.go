package model

import "omnichannel-backend-service/src/util"

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	Weight      float32 `json:"weight" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Image       string  `json:"image" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type CreateProductResponse struct {
	ID string `json:"id" binding:"required"`
}

type GetProductsRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type GetProductsResponse struct {
	Products   []*Product       `json:"products" binding:"required"`
	Pagination *util.Pagination `json:"pagination" binding:"required"`
}

type UpdateProductByNameRequest struct {
	Name        string `json:"name" binding:"required"`
	TokopediaID int    `json:"tokopedia_id"`
	ShopeeID    int    `json:"shopee_id"`
}

type Product struct {
	ID          string  `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	Weight      float32 `json:"weight" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Image       string  `json:"image" binding:"required"`
	Description string  `json:"description" binding:"required"`
	TokopediaID int     `json:"tokopedia_id"`
	ShopeeID    int     `json:"shopee_id"`
}