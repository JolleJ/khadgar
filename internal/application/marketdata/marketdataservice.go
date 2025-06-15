package marketdata

import (
	"context"
	"encoding/json"
	"jollej/db-scout/internal/domain/instrument"
	"jollej/db-scout/internal/domain/marketdata"
	"jollej/db-scout/lib/prettylog"
	"os"
	"os/exec"
	"path/filepath"
)

type MarkeDataServcie struct {
	instrumentRepo instrument.InstrumentRepo
	marketDataRepo marketdata.MarketDataRepo
}

func NewMarketDataService(instrumentRepo instrument.InstrumentRepo, marketDataRepo marketdata.MarketDataRepo) *MarkeDataServcie {
	return &MarkeDataServcie{instrumentRepo: instrumentRepo, marketDataRepo: marketDataRepo}
}

func (m *MarkeDataServcie) Load(symbols []string) error {
	type MarketDataResponse map[string][]marketdata.MarketDataPoint
	var response MarketDataResponse
	log := prettylog.NewPrettyLog()

	workingDir, err := os.Getwd() // or use a helper like findProjectRoot()
	if err != nil {
		log.Errorf("Failed to get working directory: %v", err)
		return err
	}
	projectRoot := filepath.Dir(workingDir)

	// Step 2: Build full paths
	scriptPath := filepath.Join(projectRoot, "internal", "scripts", "pyfinance.py")
	pythonPath := filepath.Join(projectRoot, "venv", "bin", "python")

	// Step 3: Build args
	args := append([]string{scriptPath}, symbols...)
	cmd := exec.Command(pythonPath, args...)

	// Step 4: Run and log output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Failed to execute script: %v\nOutput: %s", err, output)
		return err
	}

	err = json.Unmarshal(output, &response)
	if err != nil {
		log.Errorf("Failed to unmarshal JSON response: %v\nOutput: %s", err, output)
	}

	ctx := context.Background()
	for ticker, dataPoints := range response {
		if len(dataPoints) == 0 {
			log.Warningf("No data points found for ticker: %s", ticker)
			continue
		}
		err := m.marketDataRepo.AddMarketData(ctx, ticker, dataPoints)

		if err != nil {
			log.Errorf("Failed to add market data for ticker %s: %v", ticker, err)
			return err
		}

	}
	log.Infof("Successfully loaded market data for symbols: %v", symbols)

	return nil
}
