package models

// Coin holds all coin related data
type Coin struct {
	ID        int
	CmcID     int     // exported field
	Symbol    string  // exported field
	Name      string  // exported field
	Slug      string  // exported field
	Rank      int     // exported field
	Price     float64 // exported field
	MarketCap float64 // exported field
	Volume24h float64 // exported field
}

// NewCoin creates a new Coin instance
func NewCoin(cmcID int, symbol string, name string, slug string, rank int, price float64, marketCap float64, volume24h float64) *Coin {
	return &Coin{
		ID:        0,
		CmcID:     cmcID,
		Symbol:    symbol,
		Name:      name,
		Slug:      slug,
		Rank:      rank,
		Price:     price,
		MarketCap: marketCap,
		Volume24h: volume24h,
	}
}
