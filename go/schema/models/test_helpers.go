package models

import (
	"fmt"
	"strconv"
)

// Place to hold constants for testing.

const (
	testDatabase = "file::memory:?cache=shared"

	testSymbol = "BTCUSD"
	testHot    = false

	testTimestamp   = 123456
	testPrice       = 12.34
	testSize        = 200
	testMarketMaker = false
	testIsBid       = true
	testLevel       = 0
)

func testErr(actual interface{}, expected interface{}, metric string) string {
	var string1 string
	var string2 string

	switch v := actual.(type) {
	case int:
		string1 = strconv.Itoa(v)
	case int8:
		string1 = strconv.Itoa(int(v))
	case int16:
		string1 = strconv.Itoa(int(v))
	case int32:
		string1 = strconv.Itoa(int(v))
	case int64:
		string1 = strconv.Itoa(int(v))
	case uint8:
		string1 = strconv.Itoa(int(v))
	case uint16:
		string1 = strconv.Itoa(int(v))
	case uint32:
		string1 = strconv.Itoa(int(v))
	case uint64:
		string1 = strconv.Itoa(int(v))
	case float32:
		string1 = fmt.Sprintf("%f", v)
	case float64:
		string1 = fmt.Sprintf("%f", v)
	default:
		string1 = fmt.Sprintf("%v", v)
	}

	switch v := expected.(type) {
	case int:
		string2 = strconv.Itoa(v)
	case int8:
		string2 = strconv.Itoa(int(v))
	case int16:
		string2 = strconv.Itoa(int(v))
	case int32:
		string2 = strconv.Itoa(int(v))
	case int64:
		string2 = strconv.Itoa(int(v))
	case uint8:
		string2 = strconv.Itoa(int(v))
	case uint16:
		string2 = strconv.Itoa(int(v))
	case uint32:
		string2 = strconv.Itoa(int(v))
	case uint64:
		string2 = strconv.Itoa(int(v))
	case float32:
		string2 = fmt.Sprintf("%f", v)
	case float64:
		string2 = fmt.Sprintf("%f", v)
	default:
		string2 = fmt.Sprintf("%v", v)
	}

	return fmt.Sprintf(
		"database %s of %s did not match expected %s of %s",
		metric,
		string1,
		metric,
		string2,
	)
}
