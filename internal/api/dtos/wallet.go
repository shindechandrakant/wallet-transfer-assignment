package dtos

type TransferRequest struct {
	IdempotencyKey string `json:"idempotencyKey" validate:"required,min=3,max=32"`
	FromWalletId   string `json:"fromWalletId" validate:"required"`
	ToWalletId     string `json:"toWalletId" validate:"required"`
	Amount         int64  `json:"amount" validate:"required,gt=0"`
}
