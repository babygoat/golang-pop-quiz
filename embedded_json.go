package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t := struct {
		time.Time
		N int
	}{
		Time: time.Now(),
		N:    5,
	}

	m, _ := json.Marshal(t)
	// Source: https://twitter.com/katie_hockman/status/1242527795777069062
	// The embedded time.Time implements MarshalJSON method
	// Thus, the anonymous struct t was bound with the method from embedded
	// unless it delcares one itself.
	// "2020-03-25T12:08:53.806735+08:00"
	fmt.Printf("%s", m)
}
