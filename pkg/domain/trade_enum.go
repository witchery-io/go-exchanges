// Code generated by go-enum
// DO NOT EDIT!

package domain

import (
	"fmt"
)

const (
	// TradeEventTypeExecute is a TradeEventType of type Execute
	TradeEventTypeExecute TradeEventType = iota
	// TradeEventTypeUpdate is a TradeEventType of type Update
	TradeEventTypeUpdate
)

const _TradeEventTypeName = "executeupdate"

var _TradeEventTypeMap = map[TradeEventType]string{
	0: _TradeEventTypeName[0:7],
	1: _TradeEventTypeName[7:13],
}

// String implements the Stringer interface.
func (x TradeEventType) String() string {
	if str, ok := _TradeEventTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("TradeEventType(%d)", x)
}

var _TradeEventTypeValue = map[string]TradeEventType{
	_TradeEventTypeName[0:7]:  0,
	_TradeEventTypeName[7:13]: 1,
}

// ParseTradeEventType attempts to convert a string to a TradeEventType
func ParseTradeEventType(name string) (TradeEventType, error) {
	if x, ok := _TradeEventTypeValue[name]; ok {
		return x, nil
	}
	return TradeEventType(0), fmt.Errorf("%s is not a valid TradeEventType", name)
}
