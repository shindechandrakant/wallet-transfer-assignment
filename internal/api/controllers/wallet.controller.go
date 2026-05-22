package controllers

import (
	"fmt"
	"shindechandrakant/internal/api/dtos"
	"shindechandrakant/internal/wallet"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type WalletController struct {
	service *wallet.Service
}

func NewWalletController(service *wallet.Service) *WalletController {
	return &WalletController{
		service: service,
	}
}

func (h *WalletController) Transfer(ctx fiber.Ctx) error {

	var body dtos.TransferRequest
	// json parse
	if err := ctx.Bind().Body(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	// validate struct
	if err := validate.Struct(body); err != nil {

		errors := err.(validator.ValidationErrors)
		var response []string

		for _, e := range errors {
			response = append(response, fmt.Sprintf("%s failed on %s", e.Field(), e.Tag()))
		}

		return ctx.Status(422).JSON(fiber.Map{
			"errors": response,
		})
	}

	_ = ctx.Status(200).JSON(body)
	return nil
}
