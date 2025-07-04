// File: internal/core/billing.go
// Purpose: Central billing logic including invoice generation, plan checks, and billing records.

package core

import (
	"errors"
	"time"

	"shorty/internal/db"
	"shorty/internal/lib/invoice"
	"shorty/internal/models"
)

func GetBillingOverview(ctx *models.Context) (*models.BillingOverview, error) {
	switch ctx.Scope {
	case models.ScopeUser:
		return db.GetUserBillingOverview(ctx, ctx.User.ID)
	case models.ScopeTeam:
		return db.GetTeamBillingOverview(ctx, ctx.Team.ID)
	default:
		return nil, errors.New("invalid billing scope")
	}
}

func ListInvoices(ctx *models.Context) ([]*models.Invoice, error) {
	switch ctx.Scope {
	case models.ScopeUser:
		return db.GetInvoicesByUser(ctx, ctx.User.ID)
	case models.ScopeTeam:
		return db.GetInvoicesByTeam(ctx, ctx.Team.ID)
	default:
		return nil, errors.New("invalid scope for invoice listing")
	}
}

func GenerateInvoice(ctx *models.Context, scopeID, name string, amount int64, status string, periodStart, periodEnd time.Time) (*models.Invoice, error) {
	inv := &models.Invoice{
		ScopeID:     scopeID,
		ScopeType:   ctx.Scope,
		Name:        name,
		Amount:      amount,
		Status:      status,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		CreatedAt:   time.Now(),
	}

	if err := db.SaveInvoice(ctx, inv); err != nil {
		return nil, err
	}

	return inv, nil
}

func DownloadInvoice(ctx *models.Context, invoiceID string) ([]byte, error) {
	inv, err := db.GetInvoiceByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}
	return invoice.RenderPDF(inv)
}

func GetAnnualInvoiceSummary(ctx *models.Context) (*models.AnnualSummary, error) {
	switch ctx.Scope {
	case models.ScopeUser:
		return db.GetAnnualSummaryForUser(ctx, ctx.User.ID)
	case models.ScopeTeam:
		return db.GetAnnualSummaryForTeam(ctx, ctx.Team.ID)
	default:
		return nil, errors.New("invalid billing scope")
	}
}

func CheckBillingStatus(ctx *models.Context) (*models.BillingStatus, error) {
	return db.GetBillingStatus(ctx)
}
