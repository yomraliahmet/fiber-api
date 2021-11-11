package controllers

// Login godoc
// @Summary Auth Login
// @Accept  json
// @Produce  json
// @Param default body request.Auth true "Default"
// @Success 200 {object} resources.JSONResultSuccess{data=resources.Auth} "success"
// @Failure default {object} resources.JSONResultError{data=string} "error"
// @Router /auth/login [post]
// @Tags Auth
func SwaggerLogin()
