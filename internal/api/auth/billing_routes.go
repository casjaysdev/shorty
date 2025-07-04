// File: internal/api/auth/billing_routes.go
// Purpose: Provides authenticated API endpoints for billing overview, invoices, and downloads.

package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"shorty/internal/core"
	"shorty/internal/lib/httpx"
	"shorty/internal/middleware"
)

func RegisterBillingRoutes(r chi.Router) {
	r.Route("/billing", func(r chi.Router) {
		r.Use(middleware.RequireAuth)

		r.Get("/", httpx.Handler(getBillingOverview))
		r.Get("/invoices", httpx.Handler(listInvoices))
		r.Get("/invoices/{id}", httpx.Handler(getInvoicePDF))
		r.Get("/summary", httpx.Handler(getAnnualSummary))
	})
}

func getBillingOverview(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	overview, err := core.GetBillingOverview(ctx)
	if err != nil {
		return err
	}
	return httpx.JSON(w, overview)
}

func listInvoices(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	invoices, err := core.ListInvoices(ctx)
	if err != nil {
		return err
	}
	return httpx.JSON(w, invoices)
}

func getInvoicePDF(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	invoiceID := chi.URLParam(r, "id")

	pdf, err := core.DownloadInvoice(ctx, invoiceID)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"invoice.pdf\"")
	_, err = w.Write(pdf)
	return err
}

func getAnnualSummary(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	summary, err := core.GetAnnualInvoiceSummary(ctx)
	if err != nil {
		return err
	}
	return httpx.JSON(w, summary)
}
