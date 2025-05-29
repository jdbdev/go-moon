package models

// Coin Interface holds all coin related methods
type CoinInterface interface {
	Create()
	Update()
	Delete()
	// getAll()
	// getByID()
	// getBySymbol()
	// getBySlug()
	// getByRank()
	// getByPrice()
	// getByMarketCap()
	// getByVolume24h()
}

// Coin Struct holds all coin related data
type Coin struct {
	ID         int
	cmc_ID     int
	symbol     string
	name       string
	slug       string
	rank       int
	price      float64
	market_cap float64
	volume_24h float64
}

func (c *Coin) Create() {
	// business logic
}

func (c *Coin) Update() {
	// business logic
}

func (c *Coin) Delete() {
	// business logic
}

// Constructor method for Coin Struct
func NewCoin(cmc_ID int, symbol string, name string, slug string, rank int, price float64, market_cap float64, volume_24h float64) *Coin {
	return &Coin{
		ID:         0,
		cmc_ID:     cmc_ID,
		symbol:     symbol,
		name:       name,
		slug:       slug,
		rank:       rank,
		price:      price,
		market_cap: market_cap,
		volume_24h: volume_24h,
	}
}
