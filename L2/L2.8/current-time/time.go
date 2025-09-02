package currenttime

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetCurrentTime() time.Time {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error code: %d - Something went wrong!\n", 1)
		os.Exit(1)
		return time.Time{}
	}

	return time.Now().Add(response.ClockOffset)

}
