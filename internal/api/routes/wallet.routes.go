package routes

import (
	"shindechandrakant/internal/api/controllers"

	"github.com/gofiber/fiber/v3"
)

func WalletRoutes(app fiber.Router, h *controllers.WalletController) {
	wallet := app.Group("/wallet")
	wallet.Post("/transfer", h.Transfer)
}
