package middleware

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ReqId(ctx *fiber.Ctx) error {
	reqId := randSeq(24)
	ctx.Locals("reqId", reqId)
	return ctx.Next()
}
