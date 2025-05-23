# Topic7-SimpleInvestmentManagement

IF-48-INT

### ARRANGED BY:

1. EDWARD WU; 103012440003 As Leader 

2. DIMAS RADITYA PUTRA HANDOKO; 103012440016 As Member Group


# Basic Investment Management App - Code Walkthrough

# Introduction

This Go application allows users to track a virtual investment portfolio consisting of different asset classes: equities (S&P 500 stocks), commodities (gold, silver), and cryptocurrencies. Users are allowed to buy and sell assets, track portfolio performance, and simulate movement in the market over time. The application also makes posts on the portfolio's value, daily volatility, and return on investment (ROI).

### Key Functions in the Code

*Asset Struct*

A struct to hold a single investment asset (e.g., a cryptocurrency or stock).
Saves details like the asset name, price, quantity, total price, NAV (Net Asset Value), volatility, and daily price change.

*Portfolio Struct*

Holds the user portfolio information, including the balance (funds available for investing), assets owned, and setting the current day and date.
Stores historical portfolio performance and the maximum value (ATH - All-Time High).

*Simulating Price Changes*

The SimulatePriceChange() method alters asset prices as per their volatility and selected time frame.
It uses random number generation to simulate actual market patterns such that every asset reacts in terms of its volatility.
Stock, commodity, and cryptocurrency prices are altered differently as per their market nature.

*Purchase and Sale of Assets*

BuyAsset() and SellAsset() functions allow the user to purchase or sell assets by entering the asset name and quantity.
The system checks whether the user has sufficient balance for buying and enough quantity for selling.

*Portfolio Display*

The routine ShowPortfolio() displays the current user's balance, portfolio value total (balance + NAV of holdings), and specific details of holdings like profit/loss and return percentage.
Time Simulation
The routine AdvanceTime() moves time ahead by the chosen interval (1 hour, 1 day, 1 week).
It calculates asset prices, portfolio NAV, and tracks principal price movements, illustrating a summary of principal market movement.

*Investment History*

The app tracks each day's portfolio values and ROI, storing them in the history array.
The ShowInvestmentHistory() method displays the portfolio's historical performance.

# Main Program Flow

### User Input

When initialized, the program prompts the user for his name and starting balance.
Portfolio is set up with the input balance and an empty list of assets.

# Interactive Menu

### The main loop presents the user with several options:

- View portfolio
- View available assets (stocks, commodities, cryptocurrencies)
- Buy/Sell assets
- Advance time
- View investment summary

### Time Interval Options

The TimeInterval struct defines several time intervals, which determine the volatility factor and price change scale. These intervals are used to input a new time interval during runtime.
The user may choose a time period (e.g. 1 hour, 1 day) to experiment with how prices evolve over time.

### Investment Summary

The ShowInvestmentSummary() function displays a summary of overall portfolio worth and ROI during the simulation.

# Conclusion

This application offers a comprehensive simulation for managing investments, where users can interactively make decisions about buying and selling assets, track portfolio performance, and simulate market changes based on different time intervals. It's a useful tool for understanding the dynamics of asset management and the impacts of market volatility.
