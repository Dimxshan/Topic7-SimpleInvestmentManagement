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
	InitialBalance float64
	ATHValue float64 
}

// TimeInterval represents different time intervals for simulation :3
type TimeInterval struct {
	Name     string
	Display  string
	Duration time.Duration
	Factor   float64 // Volatility factor based on time interval 
}

// Available time intervals = kalian bisa pilih saja loh , mau 1D , 1W , 1M , semua adaa
var timeIntervals = []TimeInterval{
	{"1H", "1 Hour", time.Hour, 0.2},
	{"4H", "4 Hours", 4 * time.Hour, 0.4},
	{"1D", "1 Day", 24 * time.Hour, 1.0},
	{"1W", "1 Week", 7 * 24 * time.Hour, 2.5},
	{"1M", "1 Month", 30 * 24 * time.Hour, 5.0},
}

var sp500 = []Asset{
	{"Tesla", 750.0, 0, 0, 0, 0.05, 0},
	{"Apple", 145.0, 0, 0, 0, 0.035, 0},
	{"Microsoft", 300.0, 0, 0, 0, 0.032, 0},
	{"Amazon", 3400.0, 0, 0, 0, 0.042, 0},
	{"Google", 2800.0, 0, 0, 0, 0.038, 0},
	{"Facebook", 330.0, 0, 0, 0, 0.045, 0},
	{"Berkshire Hathaway", 420000.0, 0, 0, 0, 0.028, 0},
	{"Johnson & Johnson", 175.0, 0, 0, 0, 0.025, 0},
	{"Visa", 230.0, 0, 0, 0, 0.030, 0},
	{"Nvidia", 670.0, 0, 0, 0, 0.055, 0},
}

var commodities = []Asset{
	{"Gold", 1800.0, 0, 0, 0, 0.02, 0},
	{"Silver", 25.0, 0, 0, 0, 0.03, 0},
}

var cryptocurrencies = []Asset{
	{"Bitcoin", 103972.67, 0, 0, 0, 0.15, 0},
	{"Ethereum", 2520.55, 0, 0, 0, 0.18, 0},
	{"Tether", 1.00, 0, 0, 0, 0.01, 0},
	{"XRP", 2.39, 0, 0, 0, 0.22, 0},
	{"BNB", 647.41, 0, 0, 0, 0.20, 0},
	{"Solana", 171.34, 0, 0, 0, 0.25, 0},
	{"USD Coin", 0.9997, 0, 0, 0, 0.008, 0},
	{"Dogecoin", 0.2243, 0, 0, 0, 0.30, 0},
	{"Cardano", 0.7619, 0, 0, 0, 0.22, 0},
	{"Tron", 0.2729, 0, 0, 0, 0.24, 0},
}

// Helper function to generate random price changes based on asset volatility ;3
func randomPriceChange(volatility float64, intervalFactor float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	
	// Apply the interval factor to scale volatility based on time interval
	scaledVolatility := volatility * intervalFactor
	return z * scaledVolatility
}

// SimulatePriceChange updates all asset prices based on their volatility and the selected time interval :>
func SimulatePriceChange(intervalFactor float64) {
	// Create a seeded random source for more varied outcomes
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	
	// Market sentiment varies by asset class and time interval
	stockMarketSentiment := randomPriceChange(0.01, intervalFactor)
	commoditySentiment := randomPriceChange(0.005, intervalFactor)
	cryptoSentiment := randomPriceChange(0.02, intervalFactor)
	
	// Ensure we get at least some meaningful movement per interval
	// This prevents the situation where prices remain static
	minMoveMultiplier := 0.3 * intervalFactor

	// Update stock prices = Update Harga S&P 500
	for i := range sp500 {
		// Base price change on volatility with some randomness
		priceChange := randomPriceChange(sp500[i].Volatility, intervalFactor) + stockMarketSentiment
		
		// Ensure there's always some minimum movement
		if math.Abs(priceChange) < sp500[i].Volatility * minMoveMultiplier {
			// If change is too small, enforce a minimum change in the original direction
			if priceChange > 0 {
				priceChange = sp500[i].Volatility * minMoveMultiplier * (0.5 + r.Float64())
			} else {
				priceChange = -sp500[i].Volatility * minMoveMultiplier * (0.5 + r.Float64())
			}
		}
		
		// Cap the price changes based on asset class (±5% for stocks) = 5 Persen karna US Stock , Aseets Moderate Risk
		maxMove := 0.05 * intervalFactor
		if priceChange > maxMove {
			priceChange = maxMove
		} else if priceChange < -maxMove {
			priceChange = -maxMove
		}
		
		oldPrice := sp500[i].Price
		sp500[i].Price *= (1 + priceChange)
		sp500[i].Price = math.Round(sp500[i].Price*100) / 100
		if sp500[i].Price < 0.01 {
			sp500[i].Price = 0.01
		}
		sp500[i].DailyChange = (sp500[i].Price - oldPrice) / oldPrice
	}

	// Update commodity prices = Update Harga Gold & Silver 
	for i := range commodities {
		priceChange := randomPriceChange(commodities[i].Volatility, intervalFactor) + commoditySentiment
		
		// Ensure there's always some minimum movement
		if math.Abs(priceChange) < commodities[i].Volatility * minMoveMultiplier {
			if priceChange > 0 {
				priceChange = commodities[i].Volatility * minMoveMultiplier * (0.5 + r.Float64())
			} else {
				priceChange = -commodities[i].Volatility * minMoveMultiplier * (0.5 + r.Float64())
			}
		}
		
		// Cap the price changes based on asset class (±3% for commodities) = Karna Harga Commodity Memiliki Kecendrugan Kenaiakan Secara Constant
		maxMove := 0.03 * intervalFactor
		if priceChange > maxMove {
			priceChange = maxMove
		} else if priceChange < -maxMove {
			priceChange = -maxMove
		}
		
		oldPrice := commodities[i].Price
		commodities[i].Price *= (1 + priceChange)
		commodities[i].Price = math.Round(commodities[i].Price*100) / 100
		if commodities[i].Price < 0.01 {
			commodities[i].Price = 0.01
		}
		commodities[i].DailyChange = (commodities[i].Price - oldPrice) / oldPrice
	}

	// Update cryptocurrency prices - these should be more volatile = RISKY ASSETS YESS SIRRRR
	for i := range cryptocurrencies {
		// Cryptocurrencies have more dramatic movement
		priceChange := randomPriceChange(cryptocurrencies[i].Volatility, intervalFactor) + cryptoSentiment
		
		// Force significant movements for crypto
		if math.Abs(priceChange) < cryptocurrencies[i].Volatility * minMoveMultiplier {
			if priceChange > 0 {
				priceChange = cryptocurrencies[i].Volatility * minMoveMultiplier * (0.5 + r.Float64())
			} else {
				priceChange = -cryptocurrencies[i].Volatility * minMoveMultiplier * (0.5 + r.Float64())
			}
		}
		
		// Add some noise to ensure prices change every interval
		noiseComponent := (r.Float64()*0.02 - 0.01) * intervalFactor
		priceChange += noiseComponent
		
		// Cap the price changes based on asset class (10-30% for crypto) = 30 PERSEN PER DAY LMAO Let's GOOO
		maxMove := cryptocurrencies[i].Volatility * 2 * intervalFactor
		if priceChange > maxMove {
			priceChange = maxMove
		} else if priceChange < -maxMove {
			priceChange = -maxMove
		}
		
		oldPrice := cryptocurrencies[i].Price
		cryptocurrencies[i].Price *= (1 + priceChange)
		
		// Format price properly based on value
		if cryptocurrencies[i].Price < 1 {
			cryptocurrencies[i].Price = math.Round(cryptocurrencies[i].Price*10000) / 10000
		} else {
			cryptocurrencies[i].Price = math.Round(cryptocurrencies[i].Price*100) / 100
		}
		
		if cryptocurrencies[i].Price < 0.01 {
			cryptocurrencies[i].Price = 0.01
		}
		cryptocurrencies[i].DailyChange = (cryptocurrencies[i].Price - oldPrice) / oldPrice
	}
}

// UpdatePortfolioNAV updates all asset NAVs based on current prices = NAV / NET ASSETS VALUE BREE biar ga lupaa 
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

// AdvanceTime simulates the passage of time based on the chosen interval
func (p *Portfolio) AdvanceTime(interval TimeInterval) {
	// Increment day counter based on interval
	daysElapsed := int(interval.Duration.Hours() / 24)
	if daysElapsed < 1 {
		daysElapsed = 1 // Minimum 1 day increment for tracking purposes
	}
	p.CurrentDay += daysElapsed
	
	// Advance date by the interval duration
	p.CurrentDate = p.CurrentDate.Add(interval.Duration)
	
	// Simulate price changes with the interval factor
	SimulatePriceChange(interval.Factor)
	
	// Update portfolio NAVs
	p.UpdatePortfolioNAV()
	p.TrackDailyHistory()
	
	fmt.Printf("\n==== %s PASSED (%s) ====\n", interval.Display, p.CurrentDate.Format("Monday, January 2, 2006 15:04 MST"))
	fmt.Printf("Market has updated after %s interval. All asset prices have been adjusted.\n", interval.Display)
	
	// Print summary of major price movements
	fmt.Println("\nMajor Price Movements:")
	fmt.Println("------------------------")
	
	// Track significant price movements to show to the user
	printedMovements := 0
	
	// Check for significant stock movements (>2%)
	for _, asset := range sp500 {
		if math.Abs(asset.DailyChange) > 0.02 {
			direction := "increased"
			if asset.DailyChange < 0 {
				direction = "decreased"
			}
			fmt.Printf("%-20s: %s by %.2f%% to $%.2f\n", 
				asset.Name, direction, math.Abs(asset.DailyChange*100), asset.Price)
			printedMovements++
			if printedMovements >= 3 {
				break
			}
		}
	}
	
	// Check for significant crypto movements (>5%)
	for _, asset := range cryptocurrencies {
		if math.Abs(asset.DailyChange) > 0.05 {
			direction := "increased"
			if asset.DailyChange < 0 {
				direction = "decreased"
			}
			
			// Format price based on value
			var priceStr string
			if asset.Price < 1 {
				priceStr = fmt.Sprintf("$%.4f", asset.Price)
			} else {
				priceStr = fmt.Sprintf("$%.2f", asset.Price)
			}
			
			fmt.Printf("%-20s: %s by %.2f%% to %s\n", 
				asset.Name, direction, math.Abs(asset.DailyChange*100), priceStr)
			printedMovements++
			if printedMovements >= 6 {
				break
			}
		}
	}
	
	// If no significant movements, show this message
	if printedMovements == 0 {
		fmt.Println("No significant price movements in this interval.")
	}
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
	
	// Calculate the correct proportion of the original total purchase price
	// This ensures we're correctly accounting for partial sales
	sellRatio := float64(quantity) / float64(asset.Quantity)
	proportionalCost := asset.TotalPrice * sellRatio
	
	// Calculate the sale proceeds based on current price
	saleProceeds := float64(quantity) * price
	
	// Add the sale proceeds to the balance
	p.Balance += saleProceeds
	
	// Update the asset quantity
	asset.Quantity -= quantity
	
	// Update the remaining total purchase price correctly
	asset.TotalPrice -= proportionalCost
	
	// Update NAV based on remaining quantity
	asset.NAV = float64(asset.Quantity) * asset.Price
	
	// Update the asset in the portfolio
	if asset.Quantity > 0 {
		p.Assets[name] = asset
	} else {
		// If no quantity remains, remove the asset from portfolio
		delete(p.Assets, name)
	}
	
	fmt.Printf("Successfully sold %d units of %s for $%.2f\n", quantity, name, saleProceeds)
	return true
}

func (p *Portfolio) ShowPortfolio() {
	fmt.Printf("\nCurrent Balance: $%.2f\n", p.Balance)
	fmt.Printf("Current Date: %s\n", p.CurrentDate.Format("Monday, January 2, 2006 15:04 MST"))
	
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
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%-20s | %-10s | %-15s | %-10s\n", "Company", "Price ($)", "Daily Change (%)", "Volatility")
	fmt.Println("------------------------------------------------------------------")
	for _, asset := range sp500 {
		changeStr := fmt.Sprintf("%.2f%%", asset.DailyChange*100)
		if asset.DailyChange > 0 {
			changeStr = "+" + changeStr
		}
		volatilityStr := fmt.Sprintf("±%.1f%%", asset.Volatility*100)
		fmt.Printf("%-20s | $%-9.2f | %-15s | %-10s\n", asset.Name, asset.Price, changeStr, volatilityStr)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("US Stocks typically have a volatility range of around ±5%")
}

func showCommodities() {
	fmt.Println("\nAvailable Commodities:")
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%-20s | %-10s | %-15s | %-10s\n", "Commodity", "Price ($)", "Daily Change (%)", "Volatility")
	fmt.Println("------------------------------------------------------------------")
	for _, asset := range commodities {
		changeStr := fmt.Sprintf("%.2f%%", asset.DailyChange*100)
		if asset.DailyChange > 0 {
			changeStr = "+" + changeStr
		}
		volatilityStr := fmt.Sprintf("±%.1f%%", asset.Volatility*100)
		fmt.Printf("%-20s | $%-9.2f | %-15s | %-10s\n", asset.Name, asset.Price, changeStr, volatilityStr)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Commodities typically have a volatility range of around ±3%")
}

func showCryptocurrencies() {
	fmt.Println("\nAvailable Cryptocurrencies (Top 10 Market Cap):")
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%-20s | %-12s | %-15s | %-10s\n", "Cryptocurrency", "Price ($)", "Daily Change (%)", "Volatility")
	fmt.Println("------------------------------------------------------------------")
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
		
		volatilityStr := fmt.Sprintf("±%.1f%%", asset.Volatility*100)
		fmt.Printf("%-20s | %-12s | %-15s | %-10s\n", asset.Name, priceStr, changeStr, volatilityStr)
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Cryptocurrencies have a higher volatility range of around ±10-30%")
}

func showTimeIntervals() {
	fmt.Println("\nAvailable Time Intervals:")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%-5s | %-10s | %-30s\n", "Code", "Duration", "Description")
	fmt.Println("-------------------------------------------------------")
	for _, interval := range timeIntervals {
		fmt.Printf("%-5s | %-10s | %-30s\n", 
			interval.Name, 
			interval.Display, 
			fmt.Sprintf("Advance time by %s", interval.Display))
	}
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Note: Longer time intervals will result in greater price movements")
}


type HistoryEntry struct {
    Day   int
    Value float64
    ROI   float64
}

var history []HistoryEntry

func (p *Portfolio) TrackDailyHistory() {
	p.UpdatePortfolioNAV()
    p.UpdatePortfolioNAV()
    totalValue := p.Balance
    for _, asset := range p.Assets {
        totalValue += asset.NAV
    }
    
	athBefore := p.ATHValue
	if totalValue > p.ATHValue {
		p.ATHValue = totalValue
	}
	roi := ((totalValue - athBefore) / athBefore) * 100

    history = append(history, HistoryEntry{
        Day:   p.CurrentDay,
        Value: totalValue,
        ROI:   roi,
    })
}

func ShowInvestmentHistory() {
    fmt.Println("\nInvestment History:")
    fmt.Println("------------------------------------------------------")
    fmt.Printf("%-5s | %-15s | %-15s\n", "Day", "Total Value ($)", "ROI (%)")
    fmt.Println("------------------------------------------------------")
    for _, h := range history {
        fmt.Printf("%-5d | $%-14.2f | %-14.2f%%\n", h.Day, h.Value, h.ROI)
    }
    fmt.Println("------------------------------------------------------")
}

func main() {
	var name string
	var initialBalance float64

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Introduction and asking for initial balance = Opening Mas Bro 
	fmt.Println("Welcome to the Interactive Investment Simulator!")
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your initial balance: $")
	fmt.Scanln(&initialBalance)

	portfolio := Portfolio{
		InitialBalance: initialBalance,
		Balance:     initialBalance,
		Assets:      make(map[string]Asset),
		CurrentDay:  1,
		CurrentDate: time.Now(),
	}

	fmt.Printf("\nHello, %s! Your simulation begins on %s\n", 
		name, portfolio.CurrentDate.Format("Monday, January 2, 2006 15:04 MST"))
	fmt.Println("Your investment journey begins today.")

	// Main loop = Loopingan 
	for {
		fmt.Println("\n======================================")
		fmt.Printf("CURRENT DATE: %s\n", portfolio.CurrentDate.Format("Monday, January 2, 2006 15:04 MST"))
		fmt.Println("======================================")
		fmt.Println("Choose an option:")
		fmt.Println("1. View Portfolio")
		fmt.Println("2. View Available Stocks (S&P 500)")
		fmt.Println("3. View Commodities (Gold & Silver)")
		fmt.Println("4. View Cryptocurrencies (Top 10 Market Cap)")
		fmt.Println("5. Buy Asset")
		fmt.Println("6. Sell Asset")
		fmt.Println("7. Advance Time")
		fmt.Println("8. View Time Interval Options")
		fmt.Println("9. View Investment Summary")
		fmt.Println("10. Exit")
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
			if !bought {
				for _, asset := range commodities {
					if strings.EqualFold(asset.Name, assetName) {
						bought = portfolio.BuyAsset(asset.Name, quantity, asset.Price)
						break
					}
				}
			}
			if !bought {
				for _, asset := range cryptocurrencies {
					if strings.EqualFold(asset.Name, assetName) {
						bought = portfolio.BuyAsset(asset.Name, quantity, asset.Price)
						break
					}
				}
			}

			if !bought {
				fmt.Println("Asset not found or insufficient funds!")
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
			if !sold {
				for _, asset := range commodities {
					if strings.EqualFold(asset.Name, assetName) {
						sold = portfolio.SellAsset(asset.Name, quantity, asset.Price)
						break
					}
				}
			}
			if !sold {
				for _, asset := range cryptocurrencies {
					if strings.EqualFold(asset.Name, assetName) {
						sold = portfolio.SellAsset(asset.Name, quantity, asset.Price)
						break
					}
				}
			}

			if !sold {
				fmt.Println("Asset not found or insufficient quantity to sell!")
			}
			
		case 7:
			// Show available interval options
			showTimeIntervals()
			
			// Ask user to select an interval
			var intervalChoice string
			fmt.Print("\nSelect time interval (1H, 4H, 1D, 1W, 1M): ")
			fmt.Scanln(&intervalChoice)
			
			// Find the selected interval
			var selectedInterval TimeInterval
			validInterval := false
			
			for _, interval := range timeIntervals {
				if strings.EqualFold(interval.Name, intervalChoice) {
					selectedInterval = interval
					validInterval = true
					break
				}
			}
			
			if validInterval {
				// Advance time by the selected interval
				portfolio.AdvanceTime(selectedInterval)
			} else {
				fmt.Println("Invalid time interval selected. Please try again.")
			}
		
		case 8:
			// Show available time intervals and their effects
			showTimeIntervals()
			
		case 9:
            ShowInvestmentHistory()
        case 10:
			fmt.Printf("\nThank you for using the Investment Simulator, %s!\n", name)
			fmt.Printf("Your simulation ran for %d days.\n", portfolio.CurrentDay)
			fmt.Println("Final portfolio summary:")
			portfolio.ShowPortfolio()
			fmt.Println("\nHappy investing in the real world!")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
// ShowInvestmentSummary displays the day-by-day investment progress including total value and ROI = INI ADALAH BAGIAN YANG BARU / FITUR TAMBAHAN NAISEE
func (p *Portfolio) ShowInvestmentSummary(initialBalance float64) {
    fmt.Println("\nInvestment Summary:")
    fmt.Println("------------------------------------------------------")
    fmt.Printf("%-5s | %-15s | %-15s\n", "Day", "Total Value ($)", "ROI (%)")
    fmt.Println("------------------------------------------------------")

    totalInvestmentDays := p.CurrentDay
    if totalInvestmentDays < 1 {
        totalInvestmentDays = 1
    }

    totalValue := p.Balance
    for _, asset := range p.Assets {
        totalValue += asset.NAV
    }

    
	athBefore := p.ATHValue
	if totalValue > p.ATHValue {
		p.ATHValue = totalValue
	}
	roi := ((totalValue - athBefore) / athBefore) * 100


    fmt.Printf("%-5d | $%-14.2f | %-14.2f%%\n", p.CurrentDay, totalValue, roi)
    fmt.Println("------------------------------------------------------")
}
