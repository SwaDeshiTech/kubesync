package enums

type Status int

const (
	Pending Status = iota
	Running
	Success
	Failed
)
