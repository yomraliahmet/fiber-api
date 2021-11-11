package controllers

// CreateUser godoc
// @Summary Create new User
// @Accept  json
// @Produce  json
// @Param default body request.User true "Default"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.User} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /users [post]
// @Tags Users
// @Security ApiKeyAuth
func SwaggerCreateUser()

// GetUser godoc
// @Summary Show all Users
// @Accept  json
// @Produce  json
// @Success 200 {object} resources.JSONResultSuccess{data=[]resources.User} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /users [get]
// @Tags Users
// @Security ApiKeyAuth
func SwaggerGetUsers()

// GetUser godoc
// @Summary Show a User
// @ID 1
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.User} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /users/{id} [get]
// @Tags Users
// @Security ApiKeyAuth
func SwaggerGetUser()

// UpdateUser godoc
// @Summary Update a User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param default body request.User true "Default"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.User} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /users/{id} [put]
// @Tags Users
// @Security ApiKeyAuth
func SwaggerUpdateUser()

// DeleteUser godoc
// @Summary Delete a User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} resources.JSONResultSuccessString{data=string} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /users/{id} [delete]
// @Tags Users
// @Security ApiKeyAuth
func SwaggerDeleteUser()
