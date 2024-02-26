package data_types

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestSimplifiedQuery_MarshalJSON(t *testing.T) {
	query := Query{
		Criteria: []Criterion{
			{
				Field: Field{
					Name:  "reliability",
					Value: "2.0",
				},
				Operation: Greater,
			},
		},
		Order: []Order{
			{
				Field:     "score",
				Direction: Desc,
			},
		},
	}
	simplified, err := query.Simplify()
	got, err := json.Marshal(simplified)
	if err != nil {
		t.Errorf("failed to marshal query")
	}
	wanted := "{\"allocated_storage\":0,\"order\":[[\"score\",\"desc\"]],\"reliability\":{\"gt\":\"2.0\"},\"type\":\"on-demand\"}"
	if !bytes.Equal(got, []byte(wanted)) {
		t.Errorf("got %s; want %s", got, wanted)
	}
}

func TestSimplifiedQuery_MarshalJSON_2(t *testing.T) {

	query := Query{
		Criteria: []Criterion{
			{
				Field: Field{
					Name:  "reliability",
					Value: "2.0",
				},
				Operation: Greater,
			},
			{
				Field: Field{
					Name:  "reliability",
					Value: "2.1",
				},
				Operation: SmallerOrEqual,
			},
		},
	}
	simplified, err := query.Simplify()
	got, err := json.Marshal(simplified)
	if err != nil {
		t.Errorf("failed to marshal query")
	}
	wanted := "{\"allocated_storage\":0,\"reliability\":{\"gt\":\"2.0\",\"lte\":\"2.1\"},\"type\":\"on-demand\"}"
	if !bytes.Equal(got, []byte(wanted)) {
		t.Errorf("got %s; want %s", got, wanted)
	}
}

func TestSimplifiedQuery_MarshalJSON_3(t *testing.T) {

	query := Query{
		Criteria: []Criterion{
			{
				Field: Field{
					Name:  "reliability",
					Value: "2.0",
				},
				Operation: Greater,
			},
			{
				Field: Field{
					Name:  "reliability",
					Value: "2.1",
				},
				Operation: SmallerOrEqual,
			},
			{
				Field: Field{
					Name:  "verified",
					Value: "true",
				},
				Operation: Equal,
			},
			{
				Field: Field{
					Name:  "verified",
					Value: "true",
				},
				Operation: Equal,
			},
		},
	}
	simplified, err := query.Simplify()
	got, err := json.Marshal(simplified)
	if err != nil {
		t.Errorf("failed to marshal query")
	}
	wanted := "{\"allocated_storage\":0,\"reliability\":{\"gt\":\"2.0\",\"lte\":\"2.1\"},\"type\":\"on-demand\",\"verified\":{\"eq\":\"true\"}}"
	if !bytes.Equal(got, []byte(wanted)) {
		t.Errorf("got %s; want %s", got, wanted)
	}
}
