package closedcat

import "time"

type User struct {
	Join          time.Time
	Email         string
	EmailVerified bool
}

func MakeUser(Join time.Time, Email string, EmailVerified bool) User {
	return User{Join: Join, Email: Email, EmailVerified: EmailVerified}
}
