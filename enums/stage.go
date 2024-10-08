package enums

type Status string

const (
	Pending Status = "Pending"
	Running Status = "Running"
	Success Status = "Success"
	Failed  Status = "Failed"
)
