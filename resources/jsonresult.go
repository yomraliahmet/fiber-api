package resources

type JSONResult struct {
	Code    int         `json:"code" extensions:"x-order=1"`
	Message string      `json:"message" extensions:"x-order=2"`
	Data    interface{} `json:"data" swaggertype:"array,object" extensions:"x-order=3"`
}

type JSONResultSuccess struct {
	Code    int         `json:"code" extensions:"x-order=1" example:"200"`
	Message string      `json:"message" extensions:"x-order=2" example:"Success"`
	Data    interface{} `json:"data" swaggertype:"array,object" extensions:"x-order=3"`
}

type JSONResultSuccessString struct {
	Code    int    `json:"code" extensions:"x-order=1" example:"200"`
	Message string `json:"message" extensions:"x-order=2" example:"Success"`
	Data    string `json:"data" extensions:"x-order=3" example:"null"`
}

type JSONResultError struct {
	Code    int    `json:"code" extensions:"x-order=1" example:"404"`
	Message string `json:"message" extensions:"x-order=2" example:"Not Found"`
	Data    string `json:"data" extensions:"x-order=3" example:"null"`
}

func JsonResult(code int, message string, data interface{}) JSONResult {

	jsonResult := JSONResult{}
	jsonResult.Code = code
	jsonResult.Message = message
	jsonResult.Data = data

	return jsonResult
}
