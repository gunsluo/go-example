package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
)

var total int

// 定义一个模拟的订单创建函数
func createOrder() (string, error) {
	total++
	fmt.Println("total:", total)
	// 模拟耗时操作，比如调用数据库、外部服务
	time.Sleep(5 * time.Second)
	// 返回一个模拟的订单ID
	return fmt.Sprintf("order-%d", time.Now().UnixNano()), nil
}

// curl -X POST http://localhost:3000/orders \
//   -H "X-Idempotency-Key: 123e4567-e89b-12d3-a456-426614174000" \
//   -H "Content-Type: application/json" \
//   -d '{}';

func main() {
	app := fiber.New()

	// 配置并使用 Idempotency 中间件
	// 这里使用内存存储来保存已处理的 idempotency key
	// 生产环境建议使用 Redis 等持久化存储
	app.Use(idempotency.New(idempotency.Config{
		Lifetime: 5 * time.Minute, // 设置过期时间，例如 5 分钟内相同的 key 只能被处理一次
	}))

	// 定义创建订单的路由
	app.Post("/orders", func(c fiber.Ctx) error {
		// 获取客户端发送的 idempotency key
		// idempotencyKey := c.Get("X-Idempotency-Key")
		// if idempotencyKey == "" {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"error": "Idempotency-Key header is required",
		// 	})
		// }

		// 尝试执行创建订单的操作
		orderID, err := createOrder()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create order",
			})
		}

		// 如果是第一次处理该 key，则返回新创建的订单信息
		// 如果是重复请求，中间件会自动返回之前的结果
		return c.JSON(fiber.Map{
			"id":      orderID,
			"message": "Order created successfully",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
