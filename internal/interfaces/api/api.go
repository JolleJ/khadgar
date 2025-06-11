package api

import (
	"database/sql"
	balanceApp "jollej/db-scout/internal/application/balance"
	instrumentApplication "jollej/db-scout/internal/application/instrument"
	orderApplication "jollej/db-scout/internal/application/order"
	portfolioApplication "jollej/db-scout/internal/application/portfolio"
	positionApplication "jollej/db-scout/internal/application/position"
	usersApplication "jollej/db-scout/internal/application/user"
	instrumentRepo "jollej/db-scout/internal/infrastructure/repository/instrument"
	orderRepo "jollej/db-scout/internal/infrastructure/repository/order"
	portfolioRepo "jollej/db-scout/internal/infrastructure/repository/portfolio"
	positionRepo "jollej/db-scout/internal/infrastructure/repository/position"
	"jollej/db-scout/internal/infrastructure/repository/transaction"
	userRepo "jollej/db-scout/internal/infrastructure/repository/user"
	"jollej/db-scout/internal/interfaces/api/handlers/instrument"
	"jollej/db-scout/internal/interfaces/api/handlers/order"
	"jollej/db-scout/internal/interfaces/api/handlers/portfolio"
	"jollej/db-scout/internal/interfaces/api/handlers/position"
	"jollej/db-scout/internal/interfaces/api/handlers/user"
	"jollej/db-scout/internal/interfaces/middleware"
	"log"
	"net/http"
)

func NewMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	log.Println("Starting server")

	apiMux := newApiMux(db)
	wrappedMux := middleware.Chainmiddleware(apiMux, middleware.LoggingMiddleware)

	mux.Handle("/api/", http.StripPrefix("/api", wrappedMux))

	return mux
}

// In future this has to be separated into its own mux routings and not have
// all domains in the same
func newApiMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := userRepo.NewUserRepository(db)
	userService := usersApplication.NewUserService(userRepo)

	portfolioRepo := portfolioRepo.NewPortfolioRepo(db)
	portfolioService := portfolioApplication.NewPortfolioService(portfolioRepo)

	instrumentRepo := instrumentRepo.NewInstrumentrepo(db)
	instrumentService := instrumentApplication.NewInstrumentService(instrumentRepo)

	positionRepo := positionRepo.NewPositionRepo(db)
	positionService := positionApplication.NewPositionService(positionRepo)

	orderRepo := orderRepo.NewOrderRepo(db)
	orderService := orderApplication.NewOrderService(orderRepo)

	transactionRepo := transaction.NewTransactionRepo(db)

	balanceService := balanceApp.NewBalanceService(transactionRepo)

	instrumentHandler := instrument.NewInstrumentHandler(instrumentService)
	usersHandler := user.NewUsersHandler(userService)
	portfolioHandler := portfolio.NewPortfolioHandler(portfolioService, balanceService)
	positionHandler := position.NewPositionHandler(positionService)
	orderHandler := order.NewOrderHandler(orderService)
	// Eventually all actions should be split into their own resource areas
	// User actions
	mux.HandleFunc("GET /users", usersHandler.ListUsers)
	mux.HandleFunc("GET /portfolio/{id}/balance", portfolioHandler.GetBalance)
	mux.HandleFunc("POST /portfolio/{id}/orders", orderHandler.CreateOrder)
	mux.HandleFunc("GET /portfolio/{id}", portfolioHandler.GetPortfolio)
	mux.HandleFunc("GET /instruments", instrumentHandler.ListInstruments)
	mux.HandleFunc("GET /positions/{id}", positionHandler.Get)

	return mux
}
