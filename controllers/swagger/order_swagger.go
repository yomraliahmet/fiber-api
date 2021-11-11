package controllers

// CreateOrder godoc
// @Summary Create new Order
// @Accept  json
// @Produce  json
// @Param default body request.Order true "Default"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.Order} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /orders [post]
// @Tags Orders
// @Security ApiKeyAuth
func SwaggerCreateOrder()

// GetOrders godoc
// @Summary Show all Orders
// @Accept  json
// @Produce  json
// @Success 200 {object} resources.JSONResultSuccess{data=[]resources.Order} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /orders [get]
// @Tags Orders
// @Security ApiKeyAuth
func SwaggerGetOrders()

// GetOrder godoc
// @Summary Show a Order
// @ID 3
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.Order} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /orders/{id} [get]
// @Tags Orders
// @Security ApiKeyAuth
func SwaggerGetOrder()
