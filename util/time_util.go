package util

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
	IsNotZero bool
}

func (d *CustomTime) UnmarshalJSON(b []byte) error {
	dateStr := string(b) // something like `"2017-08-20"`

	dateStr = strings.ReplaceAll(dateStr, "\"", "")

	if dateStr == "null" || dateStr == "" || dateStr == "\"\"" {
		d.IsNotZero = false
		return nil
	}

	if strings.Contains(dateStr, "T") {

		t, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return fmt.Errorf("cant parse date: %#v", err)
		}
		d.Time = t
	} else {
		t, err := time.Parse(`2006-01-02 15:04:05`, dateStr)
		if err != nil {
			return fmt.Errorf("cant parse date: %#v", err)
		}
		d.Time = t
	}
	d.IsNotZero = !d.Time.IsZero()
	return nil
}

func (d CustomTime) MarshalJSON() ([]byte, error) {
	json, err := d.Time.MarshalJSON()
	return json, err
}

func NewCustomTime(t time.Time) CustomTime {
	return CustomTime{
		Time:      t,
		IsNotZero: !t.IsZero(),
	}
}
