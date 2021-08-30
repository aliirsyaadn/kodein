package response

import "github.com/aliirsyaadn/kodein/entity"

var OK = entity.Response{
	Message: "success",
	Status: 200,
}

var InternalServerError = entity.Response{
	Message: "internal server error",
	Status: 500,
}

var BadRequest = entity.Response{
	Message: "bad request",
	Status: 400,
}