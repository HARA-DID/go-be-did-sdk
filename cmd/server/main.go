package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/delivery/http/router"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/logger"

	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	didaliassdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/did-alias"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"

	_ "github.com/meQlause/go-be-did/docs"
)

// @title           DID API
// @version         1.0
// @description     DID Backend API
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	validator.Init()
	config.InitConfig()
	config.InitBlockchain()
	logger.Init(config.GetApp().LogLevel)

	aasdk.InitializeAccountAbstractionSDK(context.Background(), config.GetConfig().HNS.AccountAbstraction, config.Blockchain())
	didrootsdk.InitializeDIDRootSDK(context.Background(), config.GetConfig().HNS.DIDRoot, config.Blockchain())
	didaliassdk.InitializeDIDAliasSDK(context.Background(), config.GetConfig().HNS.DIDAlias, config.Blockchain())

	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	router.Setup(app, config.GetConfig(), config.Blockchain())

	log.Printf("Server starting on port %s", config.GetApp().Port)
	log.Fatal(app.Listen(":" + config.GetApp().Port))
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
