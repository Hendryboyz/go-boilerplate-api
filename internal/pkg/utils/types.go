package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateTime time.Time

const timeLayout = "2006-01-02 15:04:05"

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format(timeLayout)

	return json.Marshal(formatted)
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	err := json.Unmarshal(data, &timeStr)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	t, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		return fmt.Errorf("fail to parse time: %w", err)
	}

	*d = DateTime(t)
	return nil
}
