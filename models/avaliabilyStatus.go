package models

type AvaliabilityStatus int

const (
	Unknown AvaliabilityStatus = iota
	Alugado
	Disponivel
)

func AvaliabilityStatusFromInt(i int) AvaliabilityStatus {
	switch i {
	case 1:
		return Alugado
	case 2:
		return Disponivel
	default:
		return Unknown
	}
}

func (a AvaliabilityStatus) ToInt() int {
	return int(a)
}

func (a AvaliabilityStatus) ToString() string {
	switch a {
	case Alugado:
		return "alugado"
	case Disponivel:
		return "disponivel"
	default:
		return "unknown"
	}
}

func AvaliabilityStatusFromString(s string) AvaliabilityStatus {
	switch s {
	case "alugado":
		return Alugado
	case "disponivel":
		return Disponivel
	default:
		return Unknown
	}
}
