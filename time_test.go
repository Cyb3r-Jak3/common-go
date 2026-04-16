package common

import (
	"testing"
	"time"
)

func TestResilientTimeParsesRFC3339WithTimezone(t *testing.T) {
	var rt ResilientTime
	err := rt.UnmarshalJSON([]byte(`"2023-03-15T14:30:00Z"`))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if rt.Format(time.RFC3339) != "2023-03-15T14:30:00Z" {
		t.Errorf("Expected '2023-03-15T14:30:00Z', got %s", rt.Format(time.RFC3339))
	}
}

func TestResilientTimeParsesRFC3339WithoutTimezone(t *testing.T) {
	var rt ResilientTime
	err := rt.UnmarshalJSON([]byte(`"2023-03-15T14:30:00"`))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if rt.Format("2006-01-02T15:04:05") != "2023-03-15T14:30:00" {
		t.Errorf("Expected '2023-03-15T14:30:00', got %s", rt.Format("2006-01-02T15:04:05"))
	}
}

func TestResilientTimeParsesRFC3339WithoutFractionalSeconds(t *testing.T) {
	var rt ResilientTime
	err := rt.UnmarshalJSON([]byte(`"2023-03-15T14:30:00.123456"`))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if rt.Format("2006-01-02T15:04:05.999999") != "2023-03-15T14:30:00.123456" {
		t.Errorf("Expected '2023-03-15T14:30:00.123456', got %s", rt.Format("2006-01-02T15:04:05.999999"))
	}
}

func TestResilientTimeHandlesNullValue(t *testing.T) {
	var rt ResilientTime
	err := rt.UnmarshalJSON([]byte(`null`))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !rt.IsZero() {
		t.Errorf("Expected zero time, got %s", rt.Time)
	}
}

func TestResilientTimeReturnsErrorForInvalidFormat(t *testing.T) {
	var rt ResilientTime
	err := rt.UnmarshalJSON([]byte(`"invalid-time-format"`))
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestResilientTimeHandlesEmptyInput(t *testing.T) {
	var rt ResilientTime
	err := rt.UnmarshalJSON([]byte(`""`))
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestParseResilientTime_ParseRFC3339WithTimezone(t *testing.T) {
	rt, err := ParseResilientTime("2023-03-15T14:30:00Z")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if rt.Format(time.RFC3339) != "2023-03-15T14:30:00Z" {
		t.Errorf("Expected '2023-03-15T14:30:00Z', got %s", rt.Format(time.RFC3339))
	}
}

func TestParseResilientTime_ParseRFC3339WithoutTimezone(t *testing.T) {
	rt, err := ParseResilientTime("2023-03-15T14:30:00")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if rt.Format("2006-01-02T15:04:05") != "2023-03-15T14:30:00" {
		t.Errorf("Expected '2023-03-15T14:30:00', got %s", rt.Format("2006-01-02T15:04:05"))
	}
}

func TestParseResilientTime_ParseRFC3339WithoutFractionalSeconds(t *testing.T) {
	rt, err := ParseResilientTime("2023-03-15T14:30:00.123456")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if rt.Format("2006-01-02T15:04:05.999999") != "2023-03-15T14:30:00.123456" {
		t.Errorf("Expected '2023-03-15T14:30:00.123456', got %s", rt.Format("2006-01-02T15:04:05.999999"))
	}
}

func TestParseResilientTime_HandlesNullValue(t *testing.T) {
	rt, err := ParseResilientTime("null")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !rt.IsZero() {
		t.Errorf("Expected zero time, got %s", rt.Time)
	}
}

func TestParseResilientTime_ReturnsErrorForInvalidFormat(t *testing.T) {
	_, err := ParseResilientTime("invalid-time-format")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}

func TestParseResilientTime_HandlesEmptyInput(t *testing.T) {
	_, err := ParseResilientTime("")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
