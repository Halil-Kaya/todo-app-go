package utility

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Response struct {
	Result Result `json:"result"`
	Error  Error  `json:"error"`
	Meta   Meta   `json:"meta"`
}

type Error interface {
}

type Result interface {
}

type Meta struct {
	ReqId string    `json:"reqId"`
	Time  time.Time `json:"time"`
}

func OkResponse(ctx *fiber.Ctx, ack interface{}) error {
	var response Response
	response.Result = ack
	response.Error = nil
	response.Meta.ReqId = ctx.Locals("reqId").(string)
	response.Meta.Time = time.Now()
	return ctx.JSON(response)
}

func ErrorResponse(ctx *fiber.Ctx, errorAck interface{}) error {
	var response Response
	response.Result = nil
	response.Error = errorAck
	response.Meta.ReqId = ctx.Locals("reqId").(string)
	response.Meta.Time = time.Now()
	return ctx.JSON(response)
}
