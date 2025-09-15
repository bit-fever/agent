//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package core

//=============================================================================
//===
//=== TradingSystemMap
//===
//=============================================================================

type TradingSystemMap struct {
	TradingSystems map[string]*TradingSystem
}

//=============================================================================

func NewTradingSystemMap() *TradingSystemMap {
	ss := &TradingSystemMap{}
	ss.TradingSystems = map[string]*TradingSystem{}
	return ss
}

//=============================================================================
//===
//=== TradingSystem
//===
//=============================================================================

type TradingSystem struct {
	Name       string
	DataSymbol string

	TradeLists []*TradeList
}

//=============================================================================

func NewTradingSystem() *TradingSystem {
	ts := &TradingSystem{}
	ts.TradeLists = []*TradeList{}
	return ts
}

//=============================================================================
//===
//=== TradeList
//===
//=============================================================================

type TradeList struct {
	Trades       []*Trade
	DailyProfits []*DailyProfit
}

//=============================================================================

func NewTradeList() *TradeList {
	tl := TradeList{}
	tl.Trades       = []*Trade{}
	tl.DailyProfits = []*DailyProfit{}
	return &tl
}

//=============================================================================
//===
//=== Trade
//===
//=============================================================================

type Trade struct {
	EntryDate   int
	EntryTime   int64
	EntryPrice  float64
	EntryLabel  string
	ExitDate    int
	ExitTime    int64
	ExitPrice   float64
	ExitLabel   string
	GrossProfit float64
	Contracts   int64
	Position    int64
}

//=============================================================================

func NewTrade() *Trade {
	return &Trade{}
}

//=============================================================================

type DailyProfit struct {
	Date        int
	Time        int64
	GrossProfit float64
	Trades      int64
}

//=============================================================================

func NewDailyProfit() *DailyProfit {
	return &DailyProfit{}
}

//=============================================================================
