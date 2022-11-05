package main

type ChartResponse struct {
	Chart TimeSeriesData
}

type TimeSeriesData struct {
	Result []TimeSeriesResult
	Error  interface{}
}

type TimeSeriesResult struct {
	Meta       TimeSeriesMetadata
	Timestamp  []int64
	Indicators TimeSeriesIndicators
}

type TimeSeriesMetadata struct {
	Currency             string
	Symbol               string
	InstrumentTyoe       string
	ExchangeName         string
	FirstTradeDate       int64
	GMTOffset            int64
	Timezone             string
	CurrentTradingPeriod TimeSeriesCurrentTradingPeriod
	TradingPeriods       interface{}
}

type TimeSeriesCurrentTradingPeriod struct {
	Pre     TimeSeriesTradingPeriod
	Regular TimeSeriesTradingPeriod
	Post    TimeSeriesTradingPeriod
}

type TimeSeriesTradingPeriod struct {
	Timezone  string
	Start     int64
	End       int64
	GMTOffset int64
}

type TimeSeriesIndicators struct {
	Quote []TimeSeriesQuote
}

type TimeSeriesQuote struct {
	High   []float64
	Open   []float64
	Low    []float64
	Close  []float64
	Volume []float64
}
