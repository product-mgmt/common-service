package types

type Status int

const (
	Pending Status = iota + 1
	Active
	Blocked
	Delete
)

func (s Status) String() string {
	return [...]string{"Pending", "Active", "Blocked", "Delete"}[s-1]
}
