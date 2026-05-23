package wallet

import (
	"context"
	"shindechandrakant/internal/api/dtos"
)

type Service interface {
	Transfer(ctx context.Context, req dtos.TransferRequest) (*Transfer, error)
}

type walletService struct {
	repo Repository
}

func NewWalletService(repo Repository) Service {
	return &walletService{repo: repo}
}

func (s *walletService) Transfer(ctx context.Context, req dtos.TransferRequest) (*Transfer, error) {
	// first check the balance is available in sender wallet
	// lock the rows
	// make the debit transaction in sender wallet
	// make the credit transaction in received wallet
	// make the ledger entry

	return nil, nil
}
