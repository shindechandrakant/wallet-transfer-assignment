package wallet

type Service struct {
	repo *Repository
}

func NewWalletService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}
