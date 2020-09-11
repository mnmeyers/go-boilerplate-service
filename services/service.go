package services

import (
	"context"
	"fmt"
	"go-boilerplate/models"
	"go-boilerplate/repositories"
	"sync"
	"time"
)

// Service defines the interface for the service that deals with payments
type Service interface {
	Get(ctx context.Context, customerId string) ([]models.DTO, error)
	Create(ctx context.Context, postBody models.PostBody) (*models.DTO, error)
}

// ServiceImpl defines an instance of the payment service
type ServiceImpl struct {
	repository repositories.Repository
}

var _ Service = (*ServiceImpl)(nil)
var service Service
var onceService sync.Once

// GetService returns a thread-safe singleton of the payment service.
func GetService() Service {
	onceService.Do(func() {
		service = &ServiceImpl{
			repository: repositories.GetDatabase()}
	})

	return service
}

func (service *ServiceImpl) Get(ctx context.Context, customerId string) ([]models.DTO, error) {
	ch := make(chan []models.DAO, 1)
	defer close(ch)

	var payments []models.DAO
	var err error
	go func() {
		payments = service.getPaymentsFromStripe(customerId)
		ch <- payments
	}()

	select {
	case paymentsFromStripe := <-ch:
		fmt.Println("Successfully connected to Stripe API")
		if len(paymentsFromStripe) > 0 {
			go service.repository.Update(ctx, customerId, paymentsFromStripe)
		}
	case <-time.After(2 * time.Second):
		fmt.Println("Stripe API exceeded time limit of 1 second")
		payments, err = service.repository.GetById(ctx, customerId)
	}
	if err != nil {
		return nil, err
	}

	paymentDTOs, err := service.convertPaymentDAOsToDTOs(payments)

	if err != nil {
		return nil, err
	}
	return paymentDTOs, nil
}

func (service *ServiceImpl) Create(ctx context.Context, postBody models.PostBody) (*models.DTO, error) {

}
