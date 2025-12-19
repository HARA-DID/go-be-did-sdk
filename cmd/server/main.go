package main

import (
	"log"
	"math/big"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/delivery/http/router"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/pkg"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func main() {
	cfg, _ := config.Load()
	logger.Init(cfg.App.LogLevel)

	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	net := pkg.NewNetwork([]string{
		"http://20.198.228.24:5625",
		"http://13.214.26.197:5625",
		"http://70.153.16.221:5628",
		"http://70.153.192.125:5625",
		"http://70.153.16.216:5625",
	}, "2.0", 1, utils.DefaultLogConfig())

	bc := pkg.NewBlockchain("seed-phrase-dummy2", net, big.NewInt(int64(1212)))

	router.Setup(app, cfg, bc)

	log.Printf("Server starting on port %s", cfg.App.Port)
	log.Fatal(app.Listen(":" + cfg.App.Port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
