package closedcat

type Location struct {
	Name    string
	Admins  []User
	Notify  []User
	Reports []Report
}
