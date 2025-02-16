package services

import (
	"context"
	"order-billing/src/models"
	"order-billing/src/repository"
	"time"

	"github.com/google/uuid"
)

type BillingServiceInterface interface {
	GenerateInvoice(ctx context.Context, req models.InvoiceRequest) (*models.InvoiceResponse, error)
}

type BillingService struct {
	repo *repository.BillingRepository
}

func NewBillingService(repo *repository.BillingRepository) *BillingService {
	return &BillingService{repo: repo}
}

func (s *BillingService) GenerateInvoice(ctx context.Context, req models.InvoiceRequest) (*models.InvoiceResponse, error) {
	invoiceID := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)

	invoice := &models.InvoiceResponse{
		InvoiceID: invoiceID,
		OrderID:   req.OrderID,
		UserID:    req.UserID,
		Total:     req.Total,
		Items:     req.Items,
		CreatedAt: createdAt,
	}

	err := s.repo.SaveInvoice(ctx, invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}
