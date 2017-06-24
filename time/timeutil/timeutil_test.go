package timeutil

import (
	"testing"
	"time"
)

var dt6ForDt14Tests = []struct {
	v    int64
	want int32
}{
	{int64(20060102150405), int32(20060102)}}

func TestDt6ForDt14(t *testing.T) {
	for _, tt := range dt6ForDt14Tests {
		got := Dt6ForDt14(tt.v)
		if got != tt.want {
			t.Errorf("Dt6ForDt14(%v): want %v, got %v", tt.v, tt.want, got)
		}
	}
}

var dt8ForStringTests = []struct {
	v    string
	want int32
}{
	{"2006-01-02T15:04:05Z", int32(20060102)}}

func TestDt8ForString(t *testing.T) {
	for _, tt := range dt8ForStringTests {
		got, err := Dt8ForString(time.RFC3339, tt.v)
		if err != nil {
			t.Errorf("Dt8ForString(%v): want %v, error %v", tt.v, tt.want, err)
		}
		if got != tt.want {
			t.Errorf("Dt8ForString(%v): want %v, got %v", tt.v, tt.want, got)
		}
	}
}

var monthToQuarterTests = []struct {
	v    int
	want int
}{
	{1, 1}, {2, 1}, {3, 1},
	{4, 2}, {5, 2}, {6, 2},
	{7, 3}, {8, 3}, {9, 3},
	{10, 4}, {11, 4}, {12, 4}}

func TestMonthToQuarter(t *testing.T) {
	for _, tt := range monthToQuarterTests {
		got := MonthToQuarter(tt.v)
		if got != tt.want {
			t.Errorf("MonthToQuarter(%v): want %v, got %v", tt.v, tt.want, got)
		}
	}
}

var quarterToMonthTests = []struct {
	v    int
	want int
}{
	{1, 1}, {2, 4}, {3, 7}, {4, 10}}

func TestQuarterToMonth(t *testing.T) {
	for _, tt := range quarterToMonthTests {
		got := QuarterToMonth(tt.v)
		if got != tt.want {
			t.Errorf("QuarterToMonth(%v): want %v, got %v", tt.v, tt.want, got)
		}
	}
}

/*
var fromToTests = []struct {
	v    string
	want string
}{
	{"Wed, 25 May 2016 11:07:20 +0000", "2016-05-25T11:07:20Z"}}

func FromToTest(t *testing.T) {
	for _, tt := range fromToTests {
		got, err := FromTo(tt.v, time.RFC1123Z, time.RFC3339)
		if err != nil {
			t.Errorf("FromTo(%v): want %v, error ", tt.v, tt.want, err)
		}
		if got != tt.want {
			t.Errorf("FromTo(%v): want %v, got %v", tt.v, tt.want, got)
		}
	}
}
*/
