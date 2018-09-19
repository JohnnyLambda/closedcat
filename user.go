package closedcat

import "time"

type User struct {
	Join          time.Time
	Email         string
	EmailVerified bool
}
