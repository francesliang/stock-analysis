from collections import namedtuple


StockValueBasics = namedtuple("StockValueBasics", [
    "price_to_book_ratio",
    "price_to_earnings_ratio",
    "peg_ratio",
    "dividend_yield",
    "beta"
])

StockData = namedtuple("StockData", [
    "yahoo_finance",
    "iex"
])


StockSources = namedtuple("StockSources", [
    "yahoo_finance",
    "iex"
])
