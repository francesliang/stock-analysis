import argparse
import json

import yfinance as yf
from iexfinance.stocks import Stock

from data_types import StockSources, StockValueBasics, StockData


def get_stock_sources(stock_symbol: str) -> StockSources:
    yahoo = yf.Ticker(stock_symbol)
    iex = Stock(stock_symbol)
    stock_srcs = StockSources(yahoo, iex)
    return stock_srcs


def get_stock_data(stock_srcs: StockSources) -> StockData:
    yf_info = stock_srcs.yahoo_finance.info
    iex_key_stats = stock_srcs.iex.get_key_stats()
    stock_data = StockData(yf_info, iex_key_stats)
    return stock_data


def get_stock_basics(stock_data: StockData) -> StockValueBasics:
    pb_ratio = stock_data.yahoo_finance.get("priceToBook", None)
    pe_ratio = stock_data.iex.get("peRatio", None)
    peg_ratio = stock_data.yahoo_finance.get("pegRatio", None)
    dividend_yield = stock_data.yahoo_finance.get("dividendYield", None)
    beta = stock_data.iex.get("beta", None)

    basics = StockValueBasics(
        pb_ratio,
        pe_ratio,
        peg_ratio,
        dividend_yield,
        beta
    )

    return basics


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("symbol", help="Stock symbol for trading")
    args = parser.parse_args()
    return args


def main():
    args = parse_args()
    srcs = get_stock_sources(args.symbol)
    data = get_stock_data(srcs)
    basics = get_stock_basics(data)
    print(json.dumps(basics._asdict(), indent=4))


if __name__ == "__main__":
    main()

