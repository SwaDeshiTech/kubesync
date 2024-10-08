package enums

type Priority string

const (
	P0 Priority = "P0"
	P1 Priority = "P1"
	P2 Priority = "P2"
	P3 Priority = "P3"
)

func (p Priority) String() string {
	return string(p)
}

func (p Priority) Values() []string {
	return []string{string(P0), string(P1), string(P2), string(P3)}
}
