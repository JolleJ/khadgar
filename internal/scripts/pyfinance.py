
import yfinance as yf
import json
import sys

def get_eod_data(tickers, period="1d"):
    results = {}

    for symbol in tickers:
        ticker = yf.Ticker(symbol)
        try:
            hist = ticker.history(period=period, interval="1d")
            if hist.empty:
                results[symbol] = {"error": "No data returned"}
                continue

            hist.reset_index(inplace=True)
            results[symbol] =  hist.to_dict(orient="records")
        except Exception as e:
            results[symbol] = {"error": str(e)}

    return json.dumps(results, indent=2, default=str)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"error": "No tickers provided"}, indent=2))
        sys.exit(1)

    tickers = sys.argv[1:]
    print(get_eod_data(tickers))
