package atomx

import (
	"errors"
	"time"
)

type Time struct {
	time.Time
}

func TimeFromTime(t time.Time) *Time {
	tt := Time{
		Time: t,
	}

	return &tt
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if tt, err := time.Parse(`"2006-01-02T15:04:05"`, string(data)); err != nil {
		return err
	} else {
		*t = Time{
			Time: tt,
		}
	}

	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len("2006-01-02T15:04:05")+2)
	b = append(b, '"')
	b = t.AppendFormat(b, "2006-01-02T15:04:05")
	b = append(b, '"')
	return b, nil
}
