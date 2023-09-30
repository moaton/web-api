package usecase

type revenueService interface {
	GetRevenues() entity.Revenue
}

type UseCase struct {
	revenueSvc revenueService
}
