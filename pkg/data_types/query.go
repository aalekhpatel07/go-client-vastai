package data_types

import (
	"encoding/json"
)

type Field struct {
	Name  string
	Value string
}

type Query struct {
	Criteria []Criterion
	Type     InstanceType
	Order    []Order
	Storage  float64
}

type Direction int

const (
	Asc Direction = iota
	Desc
)

func (d Direction) ToString() string {
	var s string
	switch d {
	default:
		s = ""
	case Asc:
		s = "asc"
	case Desc:
		s = "desc"
	}
	return s
}

type Order struct {
	Field     string
	Direction Direction
}

func (o Order) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{o.Field, o.Direction.ToString()})
}

type Value interface {
	Value() string
}

type SimplifiedQuery struct {
	Fields  map[string][]Criterion
	Type    InstanceType
	Order   []Order
	Storage float64
}

type Criterion struct {
	Field     Field
	Operation Operation
}

func (q Query) Simplify() (SimplifiedQuery, error) {
	fields := make(map[string][]Criterion)
	for _, criterion := range q.Criteria {
		_, exists := fields[criterion.Field.Name]
		if !exists {
			fields[criterion.Field.Name] = make([]Criterion, 0)
		}
		// TODO: Implement the actual merging of lt/gt operators here.
		fields[criterion.Field.Name] = append(fields[criterion.Field.Name], Criterion{Operation: criterion.Operation, Field: criterion.Field})
	}
	return SimplifiedQuery{Fields: fields, Type: q.Type, Order: q.Order, Storage: q.Storage}, nil
}

func (q SimplifiedQuery) MarshalJSON() ([]byte, error) {

	allCriteria := make(map[string]any)
	for field, criteria := range q.Fields {
		obj := make(map[string]string)
		for _, criterion := range criteria {
			obj[criterion.Operation.Name()] = criterion.Field.Value
		}
		allCriteria[field] = obj
	}
	if len(q.Order) > 0 {
		allCriteria["order"] = q.Order
	}
	allCriteria["type"] = q.Type
	allCriteria["allocated_storage"] = q.Storage

	return json.Marshal(allCriteria)
}
