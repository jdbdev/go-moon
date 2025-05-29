package handlers

import (
	"net/http"
	"strconv"

	"github.com/jdbdev/go-moon/data/models"
)

// CoinService defines the interface for coin operations
type CoinService interface {
	CreateCoin(coin *models.Coin) error
	GetCoin(id int) (*models.Coin, error)
	UpdateCoin(coin *models.Coin) error
	DeleteCoin(id int) error
	ListCoins() ([]models.Coin, error)
}

// CoinHandler handles web requests for coin operations
type CoinHandler struct {
	service  CoinService
	renderer Renderer // Using the same Renderer interface as page handlers
}

// NewCoinHandler creates a new CoinHandler with the provided service and renderer
func NewCoinHandler(service CoinService, renderer Renderer) *CoinHandler {
	return &CoinHandler{
		service:  service,
		renderer: renderer,
	}
}

// HandleList displays the list of coins
func (h *CoinHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	coins, err := h.service.ListCoins()
	if err != nil {
		http.Error(w, "Failed to list coins", http.StatusInternalServerError)
		return
	}

	// Pass coins to the template
	data := map[string]interface{}{
		"Coins": coins,
	}

	if err := h.renderer.RenderTemplate(w, "coins.page.tmpl", data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// HandleNewForm displays the form to create a new coin
func (h *CoinHandler) HandleNewForm(w http.ResponseWriter, r *http.Request) {
	if err := h.renderer.RenderTemplate(w, "coin_new.page.tmpl", nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// HandleCreate processes the form submission to create a new coin
func (h *CoinHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form values
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Create coin from form data
	coin := &models.Coin{
		Symbol:    r.FormValue("symbol"),
		Name:      r.FormValue("name"),
		Slug:      r.FormValue("slug"),
		Price:     parseFloat64(r.FormValue("price")),
		MarketCap: parseFloat64(r.FormValue("market_cap")),
		Volume24h: parseFloat64(r.FormValue("volume_24h")),
	}

	if err := h.service.CreateCoin(coin); err != nil {
		http.Error(w, "Failed to create coin", http.StatusInternalServerError)
		return
	}

	// Redirect to the coins list page
	http.Redirect(w, r, "/coins", http.StatusSeeOther)
}

// HandleEditForm displays the form to edit an existing coin
func (h *CoinHandler) HandleEditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid coin ID", http.StatusBadRequest)
		return
	}

	coin, err := h.service.GetCoin(id)
	if err != nil {
		http.Error(w, "Coin not found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Coin": coin,
	}

	if err := h.renderer.RenderTemplate(w, "coin_edit.page.tmpl", data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// HandleUpdate processes the form submission to update a coin
func (h *CoinHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Invalid coin ID", http.StatusBadRequest)
		return
	}

	coin := &models.Coin{
		ID:        id,
		Symbol:    r.FormValue("symbol"),
		Name:      r.FormValue("name"),
		Slug:      r.FormValue("slug"),
		Price:     parseFloat64(r.FormValue("price")),
		MarketCap: parseFloat64(r.FormValue("market_cap")),
		Volume24h: parseFloat64(r.FormValue("volume_24h")),
	}

	if err := h.service.UpdateCoin(coin); err != nil {
		http.Error(w, "Failed to update coin", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/coins", http.StatusSeeOther)
}

// HandleDelete processes the form submission to delete a coin
func (h *CoinHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Invalid coin ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCoin(id); err != nil {
		http.Error(w, "Failed to delete coin", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/coins", http.StatusSeeOther)
}

// Helper function to parse float values from forms
func parseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}
