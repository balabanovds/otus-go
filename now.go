package now

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

// Now prints current time synced with NTP
func Now() {
	time, err := ntp.Time("ntp1.stratum1.ru")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.String())
}
