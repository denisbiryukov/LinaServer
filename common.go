package main

import (
	"fmt"
	"time"
)

var (
	Port = 37
)

type Message uint32

// GetAddress constructs and retrieves a net address.
func GetAddress(host string) string {
	return fmt.Sprintf("%s:%d", host, Port)
}

// GetSeconds returns number of seconds since 00:00 (midnight) 1 January 1900 till 't' time.
func GetSeconds(t time.Time) int64 {
	var seconds int64 = 2208988800 // Seconds since 01 JAN 1900 till 01 JAN 1970 (GMT).
	return t.UTC().Add(0*time.Hour).Unix() + seconds
}
