package controllers

// CreateProduct godoc
// @Summary Create new Product
// @Accept  json
// @Produce  json
// @Param default body request.Product true "Default"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.Product} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /products [post]
// @Tags Products
// @Security ApiKeyAuth
func SwaggerCreateProduct()

// GetProducts godoc
// @Summary Show all Products
// @Accept  json
// @Produce  json
// @Success 200 {object} resources.JSONResultSuccess{data=[]resources.Product} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /products [get]
// @Tags Products
// @Security ApiKeyAuth
func SwaggerGetProducts()

// GetProduct godoc
// @Summary Show a Product
// @ID 2
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.Product} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /products/{id} [get]
// @Tags Products
// @Security ApiKeyAuth
func SwaggerGetProduct()

// UpdateProduct godoc
// @Summary Update a Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param default body request.Product true "Default"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.Product} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /products/{id} [put]
// @Tags Products
// @Security ApiKeyAuth
func SwaggerUpdateProduct()

// DeleteProduct godoc
// @Summary Delete a Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} resources.JSONResultSuccessString{data=string} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /products/{id} [delete]
// @Tags Products
// @Security ApiKeyAuth
func SwaggerDeleteProduct()
