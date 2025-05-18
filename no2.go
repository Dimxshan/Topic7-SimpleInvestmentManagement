package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Asset struct {
	Name         string
	Price        float64
	Quantity     int
	TotalPrice   float64
	NAV          float64 
	Volatility   float64 
	DailyChange  float64 
}

type Portfolio struct {
	Balance     float64
	Assets      map[string]Asset
	CurrentDay  int    
	CurrentDate time.Time 
}

var sp500 = []Asset{
	{"Tesla", 750.0, 0, 0, 0, 0.045, 0},
	{"Apple", 145.0, 0, 0, 0, 0.025, 0},
	{"Microsoft", 300.0, 0, 0, 0, 0.022, 0},
	{"Amazon", 3400.0, 0, 0, 0, 0.032, 0},
	{"Google", 2800.0, 0, 0, 0, 0.028, 0},
	{"Facebook", 330.0, 0, 0, 0, 0.035, 0},
	{"Berkshire Hathaway", 420000.0, 0, 0, 0, 0.018, 0},
	{"Johnson & Johnson", 175.0, 0, 0, 0, 0.015, 0},
	{"Visa", 230.0, 0, 0, 0, 0.020, 0},
	{"Nvidia", 670.0, 0, 0, 0, 0.055, 0},
}

var commodities = []Asset{
	{"Gold", 1800.0, 0, 0, 0, 0.012, 0},
	{"Silver", 25.0, 0, 0, 0, 0.025, 0},
}

var cryptocurrencies = []Asset{
	{"Bitcoin", 103972.67, 0, 0, 0, 0.075, 0},
	{"Ethereum", 2520.55, 0, 0, 0, 0.080, 0},
	{"Tether", 1.00, 0, 0, 0, 0.002, 0},
	{"XRP", 2.39, 0, 0, 0, 0.090, 0},
	{"BNB", 647.41, 0, 0, 0, 0.085, 0},
	{"Solana", 171.34, 0, 0, 0, 0.095, 0},
	{"USD Coin", 0.9997, 0, 0, 0, 0.001, 0},
	{"Dogecoin", 0.2243, 0, 0, 0, 0.110, 0},
	{"Cardano", 0.7619, 0, 0, 0, 0.085, 0},
	{"Tron", 0.2729, 0, 0, 0, 0.088, 0},
}

// Helper function to generate random price changes based on asset volatility
func randomPriceChange(volatility float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	
	return z * volatility
}

// SimulateDailyPriceChange updates all asset prices based on their volatility
func SimulateDailyPriceChange() {
	marketSentiment := randomPriceChange(0.01)

	for i := range sp500 {
		priceChange := randomPriceChange(sp500[i].Volatility) + marketSentiment
		
		if priceChange > 0.1 {
			priceChange = 0.1
		} else if priceChange < -0.1 {
			priceChange = -0.1
		}
		
		oldPrice := sp500[i].Price
		sp500[i].Price *= (1 + priceChange)
		
		sp500[i].Price = math.Round(sp500[i].Price*100) / 100
		
		sp500[i].DailyChange = (sp500[i].Price - oldPrice) / oldPrice
	}

	for i := range commodities {
		commoditySentiment := randomPriceChange(0.005)
		priceChange := randomPriceChange(commodities[i].Volatility) + commoditySentiment
		
		if priceChange > 0.08 {
			priceChange = 0.08
		} else if priceChange < -0.08 {
			priceChange = -0.08
		}
		
		oldPrice := commodities[i].Price
		commodities[i].Price *= (1 + priceChange)
		commodities[i].Price = math.Round(commodities[i].Price*100) / 100
		commodities[i].DailyChange = (commodities[i].Price - oldPrice) / oldPrice
	}

	cryptoSentiment := randomPriceChange(0.02)
	for i := range cryptocurrencies {
		priceChange := randomPriceChange(cryptocurrencies[i].Volatility) + cryptoSentiment
		
		if priceChange > 0.2 {
			priceChange = 0.2
		} else if priceChange < -0.2 {
			priceChange = -0.2
		}
		
		oldPrice := cryptocurrencies[i].Price
		cryptocurrencies[i].Price *= (1 + priceChange)
		
		if cryptocurrencies[i].Price < 1 {
			cryptocurrencies[i].Price = math.Round(cryptocurrencies[i].Price*10000) / 10000
		} else {
			cryptocurrencies[i].Price = math.Round(cryptocurrencies[i].Price*100) / 100
		}
		
		cryptocurrencies[i].DailyChange = (cryptocurrencies[i].Price - oldPrice) / oldPrice
	}
}

// UpdatePortfolioNAV updates all asset NAVs based on current prices
func (p *Portfolio) UpdatePortfolioNAV() {
	for name, asset := range p.Assets {
		// Find current price of the asset
		currentPrice := 0.0
		
		// Check in stocks
		for _, stock := range sp500 {
			if stock.Name == name {
				currentPrice = stock.Price
				break
			}
		}
		
		// Check in commodities
		for _, commodity := range commodities {
			if commodity.Name == name {
				currentPrice = commodity.Price
				break
			}
		}
		
		// Check in cryptocurrencies
		for _, crypto := range cryptocurrencies {
			if crypto.Name == name {
				currentPrice = crypto.Price
				break
			}
		}
		
		// Update NAV with current price
		if currentPrice > 0 {
			asset.NAV = float64(asset.Quantity) * currentPrice
			asset.Price = currentPrice // Update the price in the portfolio
			p.Assets[name] = asset
		}
	}
}

// AdvanceToNextDay simulates the passage of a day
func (p *Portfolio) AdvanceToNextDay() {
	// Increment day counter
	p.CurrentDay++
	
	// Advance date by one day
	p.CurrentDate = p.CurrentDate.AddDate(0, 0, 1)
	
	// Simulate price changes
	SimulateDailyPriceChange()
	
	// Update portfolio NAVs
	p.UpdatePortfolioNAV()
	
	fmt.Printf("\n==== DAY %d (%s) ====\n", p.CurrentDay, p.CurrentDate.Format("Monday, January 2, 2006"))
	fmt.Println("Market has closed for the day. All asset prices have been updated.")
}

func (p *Portfolio) BuyAsset(name string, quantity int, price float64) bool {
	totalPrice := float64(quantity) * price
	if totalPrice > p.Balance {
		fmt.Println("Not enough balance to buy", name)
		return false
	}
	p.Balance -= totalPrice
	asset, exists := p.Assets[name]
	if exists {
		asset.Quantity += quantity
		asset.TotalPrice += totalPrice
		asset.NAV = float64(asset.Quantity) * asset.Price // Update NAV with the current price
		p.Assets[name] = asset
	} else {
		p.Assets[name] = Asset{Name: name, Price: price, Quantity: quantity, TotalPrice: totalPrice, NAV: float64(quantity) * price}
	}
	fmt.Printf("Successfully bought %d units of %s\n", quantity, name)
	return true
}

func (p *Portfolio) SellAsset(name string, quantity int, price float64) bool {
	asset, exists := p.Assets[name]
	if !exists || asset.Quantity < quantity {
		fmt.Println("Not enough assets to sell", name)
		return false
	}
	totalPrice := float64(quantity) * price
	p.Balance += totalPrice
	asset.Quantity -= quantity
	asset.TotalPrice -= totalPrice
	asset.NAV = float64(asset.Quantity) * asset.Price // Update NAV with the current price
	p.Assets[name] = asset
	fmt.Printf("Successfully sold %d units of %s\n", quantity, name)
	return true
}

func (p *Portfolio) ShowPortfolio() {
	fmt.Printf("\nCurrent Balance: $%.2f\n", p.Balance)
	fmt.Printf("Current Day: %d (%s)\n", p.CurrentDay, p.CurrentDate.Format("Monday, January 2, 2006"))
	
	// Calculate total portfolio value
	totalValue := p.Balance
	for _, asset := range p.Assets {
		totalValue += asset.NAV
	}
	
	fmt.Printf("Total Portfolio Value: $%.2f\n", totalValue)
	fmt.Println("\nYour Assets:")
	
	if len(p.Assets) == 0 {
		fmt.Println("You don't own any assets yet.")
	} else {
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		fmt.Printf("%-20s | %-10s | %-20s | %-20s | %-15s | %-15s\n", 
			"Asset", "Quantity", "Total Purchase ($)", "Current Value ($)", "Profit/Loss ($)", "Return (%)")
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
		
		for _, asset := range p.Assets {
			if asset.Quantity > 0 { // Only show assets with non-zero quantity
				profitLoss := asset.NAV - asset.TotalPrice
				returnPct := 0.0
				if asset.TotalPrice > 0 {
					returnPct = (profitLoss / asset.TotalPrice) * 100
				}
				
				fmt.Printf("%-20s | %-10d | $%-19.2f | $%-19.2f | $%-14.2f | %-15.2f%%\n", 
					asset.Name, asset.Quantity, asset.TotalPrice, asset.NAV, profitLoss, returnPct)
			}
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	}
}

func showSP500() {
	fmt.Println("\nAvailable Stocks (Top S&P 500 Companies):")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("%-20s | %-10s | %-15s\n", "Company", "Price ($)", "Daily Change (%)")
	fmt.Println("--------------------------------------------------------")
	for _, asset := range sp500 {
		changeStr := fmt.Sprintf("%.2f%%", asset.DailyChange*100)
		if asset.DailyChange > 0 {
			changeStr = "+" + changeStr
		}
		fmt.Printf("%-20s | $%-9.2f | %-15s\n", asset.Name, asset.Price, changeStr)
	}
	fmt.Println("--------------------------------------------------------")
}

func showCommodities() {
	fmt.Println("\nAvailable Commodities:")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("%-20s | %-10s | %-15s\n", "Commodity", "Price ($)", "Daily Change (%)")
	fmt.Println("--------------------------------------------------------")
	for _, asset := range commodities {
		changeStr := fmt.Sprintf("%.2f%%", asset.DailyChange*100)
		if asset.DailyChange > 0 {
			changeStr = "+" + changeStr
		}
		fmt.Printf("%-20s | $%-9.2f | %-15s\n", asset.Name, asset.Price, changeStr)
	}
	fmt.Println("--------------------------------------------------------")
}

func showCryptocurrencies() {
	fmt.Println("\nAvailable Cryptocurrencies (Top 10 Market Cap):")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("%-20s | %-10s | %-15s\n", "Cryptocurrency", "Price ($)", "Daily Change (%)")
	fmt.Println("--------------------------------------------------------")
	for _, asset := range cryptocurrencies {
		changeStr := fmt.Sprintf("%.2f%%", asset.DailyChange*100)
		if asset.DailyChange > 0 {
			changeStr = "+" + changeStr
		}
		
		// Format based on price range
		var priceStr string
		if asset.Price < 1 {
			priceStr = fmt.Sprintf("$%.4f", asset.Price)
		} else if asset.Price > 1000 {
			priceStr = fmt.Sprintf("$%.0f", asset.Price)
		} else {
			priceStr = fmt.Sprintf("$%.2f", asset.Price)
		}
		
		fmt.Printf("%-20s | %-10s | %-15s\n", asset.Name, priceStr, changeStr)
	}
	fmt.Println("--------------------------------------------------------")
}

func main() {
	var name string
	var initialBalance float64

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Introduction and asking for initial balance
	fmt.Println("Welcome to the Interactive Investment Program!")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your initial balance: $")
	fmt.Scanln(&initialBalance)

	portfolio := Portfolio{
		Balance:     initialBalance,
		Assets:      make(map[string]Asset),
		CurrentDay:  1,
		CurrentDate: time.Now(),
	}

	fmt.Printf("\nHello, %s! Today is Day 1 (%s)\n", name, portfolio.CurrentDate.Format("Monday, January 2, 2006"))
	fmt.Println("Your investment journey begins today.")

	// Main loop
	for {
		fmt.Println("\n======================================")
		fmt.Printf("DAY %d - %s\n", portfolio.CurrentDay, portfolio.CurrentDate.Format("Monday, January 2, 2006"))
		fmt.Println("======================================")
		fmt.Println("Choose an option:")
		fmt.Println("1. View Portfolio")
		fmt.Println("2. View Available Stocks (S&P 500)")
		fmt.Println("3. View Commodities (Gold & Silver)")
		fmt.Println("4. View Cryptocurrencies (Top 10 Market Cap)")
		fmt.Println("5. Buy Asset")
		fmt.Println("6. Sell Asset")
		fmt.Println("7. Advance to Next Day")
		fmt.Println("8. Exit")
		fmt.Print("\nEnter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			portfolio.ShowPortfolio()
		case 2:
			showSP500()
		case 3:
			showCommodities()
		case 4:
			showCryptocurrencies()
		case 5:
			var assetName string
			var quantity int
			fmt.Print("Enter the asset name you want to buy: ")
			fmt.Scanln(&assetName)
			fmt.Print("Enter the quantity you want to buy: ")
			fmt.Scanln(&quantity)

			// Check if asset is in stocks, commodities, or crypto
			bought := false
			for _, asset := range sp500 {
				if strings.EqualFold(asset.Name, assetName) {
					bought = portfolio.BuyAsset(asset.Name, quantity, asset.Price)
					break
				}
			}
			for _, asset := range commodities {
				if strings.EqualFold(asset.Name, assetName) {
					bought = portfolio.BuyAsset(asset.Name, quantity, asset.Price)
					break
				}
			}
			for _, asset := range cryptocurrencies {
				if strings.EqualFold(asset.Name, assetName) {
					bought = portfolio.BuyAsset(asset.Name, quantity, asset.Price)
					break
				}
			}

			if !bought {
				fmt.Println("Asset not found!")
			}

		case 6:
			var assetName string
			var quantity int
			fmt.Print("Enter the asset name you want to sell: ")
			fmt.Scanln(&assetName)
			fmt.Print("Enter the quantity you want to sell: ")
			fmt.Scanln(&quantity)

			// Check if asset is in stocks, commodities, or crypto
			sold := false
			for _, asset := range sp500 {
				if strings.EqualFold(asset.Name, assetName) {
					sold = portfolio.SellAsset(asset.Name, quantity, asset.Price)
					break
				}
			}
			for _, asset := range commodities {
				if strings.EqualFold(asset.Name, assetName) {
					sold = portfolio.SellAsset(asset.Name, quantity, asset.Price)
					break
				}
			}
			for _, asset := range cryptocurrencies {
				if strings.EqualFold(asset.Name, assetName) {
					sold = portfolio.SellAsset(asset.Name, quantity, asset.Price)
					break
				}
			}

			if !sold {
				fmt.Println("Asset not found!")
			}
			
		case 7:
			// Advance to the next day
			portfolio.AdvanceToNextDay()
			
		case 8:
			fmt.Printf("\nThank you for using the Investment Program, %s!\n", name)
			fmt.Printf("You've simulated %d days of trading.\n", portfolio.CurrentDay)
			fmt.Println("Final portfolio summary:")
			portfolio.ShowPortfolio()
			fmt.Println("\nHappy investing in the real world!")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}