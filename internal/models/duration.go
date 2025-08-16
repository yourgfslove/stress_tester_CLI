package models

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d Duration) String() string {
	return time.Duration(d).String()
}
