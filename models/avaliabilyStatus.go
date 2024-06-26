package models

type AvaliabilityStatus int

const (
	Unknown AvaliabilityStatus = iota
	Rented
	Available
)

func AvaliabilityStatusFromInt(i int) AvaliabilityStatus {
	switch i {
	case 1:
		return Rented
	case 2:
		return Available
	default:
		return Unknown
	}
}

func (a AvaliabilityStatus) ToInt() int {
	return int(a)
}

func (a AvaliabilityStatus) ToString() string {
	switch a {
	case Rented:
		return "alugado"
	case Available:
		return "disponivel"
	default:
		return "unknown"
	}
}
