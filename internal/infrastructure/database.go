package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	err = createUserTable(db)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	err = createPortfolioTable(db)
	if err != nil {
		log.Fatalf("Error creating portfolio table: %v", err)
	}
	err = createInstrumentTable(db)
	if err != nil {
		log.Fatalf("Error creating portfolio table: %v", err)
	}
	err = createPositionsTable(db)
	if err != nil {
		log.Fatalf("Error creating positions table: %v", err)
	}
	err = createOrdersTable(db)
	if err != nil {
		log.Fatalf("Error creating order table: %v", err)
	}
	err = createTradesTable(db)
	if err != nil {
		log.Fatalf("Error creating traders table: %v", err)
	}
	err = createTransactionsTable(db)
	if err != nil {
		log.Fatalf("Error creating transactions table: %v", err)
	}
	err = createAccountsBallances(db)
	if err != nil {
		log.Fatalf("Error creating account_balances table: %v", err)
	}

	err = StartWal(db)
	if err != nil {
		log.Fatalf("Error starting WAL mode: %v", err)
	}

	err = createDummyData(db)
	if err != nil {
		log.Fatalf("Error creating dummy datata: %v", err)
	}
	log.Println("Initialized the database with data")
	return db, nil
}

func createUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL CHECK(age > 0),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT 1
	);`
	_, err := db.Exec(query)
	return err
}

func createPortfolioTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS portfolios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
		name TEXT NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err := db.Exec(query)
	return err
}

func createInstrumentTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS instruments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type TEXT NOT NULL,
		ticker TEXT NOT NULL UNIQUE,
		exchange TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	return err
}

func createPositionsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS positions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		portfolio_id INTEGER NOT NULL,
		instrument_id INTEGER NOT NULL,
		quantity REAL NOT NULL CHECK(quantity >= 0),
		average_price REAL, 
		FOREIGN KEY (portfolio_id) REFERENCES portfolios(id),
		FOREIGN KEY (instrument_id) REFERENCES instruments(id)
	);`

	_, err := db.Exec(query)
	return err
}

func createOrdersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		portfolio_id INTEGER NOT NULL,
		instrument_id INTEGER NOT NULL,
		side TEXT NOT NULL CHECK(side IN ('buy', 'sell')),
		quantity REAL NOT NULL CHECK(quantity > 0),
		price REAL NOT NULL CHECK(price >= 0),
		order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status TEXT NOT NULL CHECK(status IN ('pending', 'completed', 'cancelled')),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    filled_at TIMESTAMP,
		FOREIGN KEY (portfolio_id) REFERENCES portfolios(id),
		FOREIGN KEY (instrument_id) REFERENCES instruments(id)
	)`

	_, err := db.Exec(query)
	return err
}

func createTradesTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS trades (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id INTEGER NOT NULL,
		quantity REAL NOT NULL CHECK(quantity > 0),
		price REAL NOT NULL CHECK(price >= 0),
		executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status TEXT NOT NULL CHECK(status IN ('pending', 'completed', 'cancelled')),
		FOREIGN KEY (order_id) REFERENCES orders(id)
	);`

	_, err := db.Exec(query)
	return err
}

func createTransactionsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
	  portfolio_id INTEGER NOT NULL,
		type TEXT NOT NULL CHECK(type IN ('deposit', 'withdrawal')),
		amount REAL NOT NULL CHECK(amount > 0),
    currency TEXT NOT NULL,
		transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status TEXT NOT NULL CHECK(status IN ('pending', 'completed', 'failed')),
		FOREIGN KEY (portfolio_id) REFERENCES portfolios(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err := db.Exec(query)
	return err
}

func createAccountsBallances(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS account_balances (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		portfolio_id INTEGER NOT NULL,
		currency TEXT NOT NULL,
		balance REAL NOT NULL CHECK(balance >= 0),
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (portfolio_id) REFERENCES portfolios(id)
	);`

	_, err := db.Exec(query)
	return err
}

func createDummyData(db *sql.DB) error {
	query := `
	INSERT INTO users (name, email, age) VALUES
	('John Doe', 'heh1', 30),
	('Jane Smith', 'heh2', 25),
	('Alice Johnson', 'heh3', 28),
	('Bob Brown', 'heh4', 35),
	('Charlie White', 'heh5', 40);
	INSERT INTO portfolios (user_id, name, description) VALUES
	(1, 'Tech Portfolio', 'A portfolio focused on technology stocks'),
	(2, 'Healthcare Portfolio', 'A portfolio focused on healthcare stocks'),
	(3, 'Energy Portfolio', 'A portfolio focused on energy stocks'),
	(4, 'Finance Portfolio', 'A portfolio focused on finance stocks'),
	(5, 'Consumer Goods Portfolio', 'A portfolio focused on consumer goods stocks');
	INSERT INTO instruments (name, type, ticker, exchange) VALUES
	('Apple Inc.', 'stock', 'AAPL', 'NASDAQ'),
	('Microsoft Corporation', 'stock', 'MSFT', 'NASDAQ'),
	('Tesla Inc.', 'stock', 'TSLA', 'NASDAQ'),
	('Amazon.com Inc.', 'stock', 'AMZN', 'NASDAQ'),
	('Alphabet Inc.', 'stock', 'GOOGL', 'NASDAQ');
	INSERT INTO positions (portfolio_id, instrument_id, quantity) VALUES
	(1, 1, 10),
	(1, 2, 5),
	(2, 3, 8 ),
	(2, 4, 12),
	(3, 5, 15);
	INSERT INTO orders (portfolio_id, instrument_id, side, quantity, price, status) VALUES
	(1, 1, 'buy', 10, 150.00, 'completed'),
	(1, 2, 'sell', 5, 250.00, 'completed'),
	(2, 3, 'buy', 8, 700.00, 'completed'),
	(2, 4, 'sell', 12, 1800.00, 'completed'),
	(3, 5, 'buy', 15, 2800.00, 'completed');
	INSERT INTO trades (order_id, quantity, price, status) VALUES
	(1, 10, 150.00, 'completed'),
	(2, 5, 250.00, 'completed'),
	(3, 8, 700.00, 'completed'),
	(4, 12, 1800.00, 'completed'),
	(5, 15, 2800.00, 'completed');
	INSERT INTO transactions (user_id, portfolio_id, type, amount, currency, status) VALUES
	(1, 1, 'deposit', 10000.00, 'USD', 'completed'),
	(2, 2, 'withdrawal', 5000.00, 'USD', 'completed'),
	(3, 3, 'deposit', 15000.00, 'USD', 'completed'),
	(4, 4, 'withdrawal', 2000.00, 'USD', 'completed'),
	(5, 5, 'deposit', 8000.00, 'USD', 'completed');
	INSERT INTO account_balances (portfolio_id, currency, balance) VALUES
	(1, 'USD', 10000.00),
	(2, 'USD', 5000.00),
	(3, 'USD', 15000.00),
	(4, 'USD', 2000.00),
	(5, 'USD', 8000.00);
	`
	_, err := db.Exec(query)
	return err
}

func StartWal(db *sql.DB) error {
	// Enable WAL mode
	_, err := db.Exec("PRAGMA journal_mode=WAL;")
	return err
}
