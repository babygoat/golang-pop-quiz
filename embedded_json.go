package quiz

import (
	"encoding/json"
	"time"
)

// Source: https://twitter.com/katie_hockman/status/1242527795777069062
// The embedded time.Time implements MarshalJSON method
// Thus, the anonymous struct t was bound with the method from embedded
// unless it declares one itself.
func EmbeddedJson() []byte {
	t := struct {
		time.Time
		N int
	}{
		Time: time.Unix(1610039949, 0).UTC(),
		N:    5,
	}

	m, _ := json.Marshal(t)
	return m
	// "2020-03-25T12:08:53.806735+08:00"
}
