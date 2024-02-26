package data_types

type Operation int

const (
	Greater Operation = iota
	GreaterOrEqual
	Smaller
	SmallerOrEqual
	Equal
	NotEqual
	In
	NotIn
)

func (op Operation) Name() string {
	switch op {
	default:
		return ""
	case Greater:
		return "gt"
	case GreaterOrEqual:
		return "gte"
	case Smaller:
		return "lt"
	case SmallerOrEqual:
		return "lte"
	case Equal:
		return "eq"
	case NotEqual:
		return "neq"
	case In:
		return "in"
	case NotIn:
		return "notin"
	}
}
