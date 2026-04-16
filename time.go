package common

import (
	"fmt"
	"time"
)

// ResilientTime Custom time type to handle both RFC3339, no timezone and no fractional seconds.
type ResilientTime struct {
	time.Time
}

func (t *ResilientTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// Remove quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	// Handle null
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}
	// Try RFC3339 with timezone
	if ts, err := time.Parse(time.RFC3339Nano, s); err == nil {
		t.Time = ts
		return nil
	}
	// Try RFC3339 without timezone
	if ts, err := time.Parse("2006-01-02T15:04:05.999999", s); err == nil {
		t.Time = ts
		return nil
	}
	// Try RFC3339 without fractional seconds
	if ts, err := time.Parse("2006-01-02T15:04:05", s); err == nil {
		t.Time = ts
		return nil
	}
	return fmt.Errorf("cannot parse time: %s", s)
}

func ParseResilientTime(s string) (ResilientTime, error) {
	// Handle null
	if s == "null" {
		return ResilientTime{Time: time.Time{}}, nil
	}
	// Try RFC3339 with timezone
	if ts, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return ResilientTime{Time: ts}, nil
	}
	// Try RFC3339 without timezone
	if ts, err := time.Parse("2006-01-02T15:04:05.999999", s); err == nil {
		return ResilientTime{Time: ts}, nil
	}
	// Try RFC3339 without fractional seconds
	if ts, err := time.Parse("2006-01-02T15:04:05", s); err == nil {
		return ResilientTime{Time: ts}, nil
	}
	return ResilientTime{}, fmt.Errorf("cannot parse time: %s", s)
}
