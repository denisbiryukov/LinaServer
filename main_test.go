package main

import (
	"fmt"
	"testing"
	"time"
)

// TestTimeCompare compares retrieved times from the localhost and another remote time server.
func TestTimeCompare(t *testing.T) {
	// Start the local server.
	go func() {
		server := Server{
			isRunning: true,
			listener:  nil,
		}
		server.Run()
	}()

	// Start the local client.
	client := Client{}
	host1 := "time.nist.gov"
	message1 := client.Request(host1)
	fmt.Println(host1, " Received: ", message1)

	host2 := "localhost"
	message2 := client.Request(host2)
	fmt.Println(host2, "Received: ", message2)

	// Compare results.
	diff := message2 - message1
	diff *= diff
	if diff > 1 {
		t.Errorf("Test failed: results not match: %d %d.", message1, message2)
	}
}

// TestGetSeconds verifies the GetSeconds functionality.
func TestGetSeconds(t *testing.T) {
	dateValues := []struct {
		date    string
		seconds int64
	}{
		{"01-Jan-1970", 2208988800},
		{"01-Jan-1976", 2398291200},
		{"01-Jan-1980", 2524521600},
		{"01-May-1983", 2629584000},
		{"17-Nov-1858", -1297728000},
	}
	for _, dateValue := range dateValues {
		date, err := time.Parse("02-Jan-2006", dateValue.date)
		if err != nil {
			t.Fatalf("time.Parse() %v : %v", dateValue.date, err)
		}
		seconds := GetSeconds(date)
		if seconds != dateValue.seconds {
			t.Errorf("Test failed: results not match for date %v: %d %d.", dateValue.date, seconds, dateValue.seconds)
		}
	}
}
