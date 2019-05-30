package now

import (
	"github.com/beevik/ntp"
	"log"
)

// Now prints current time synced with NTP
func Now() string {
	time, err := ntp.Time("ntp1.stratum1.ru")
	if err != nil {
		log.Fatal(err)
	}

	return time.String()
}
