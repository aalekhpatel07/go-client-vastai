package command

import (
	"bytes"
	"encoding/json"
	"github.com/aalekhpatel07/go-vastai-client/pkg/data_types"
	"testing"
)

func TestNewSearchOffersRequest(t *testing.T) {
	search := NewSearchOffersRequest()
	search.WithOrder(data_types.Order{
		Field:     "score",
		Direction: data_types.Desc,
	}).WithInstanceType(data_types.Bid).WithStorage(10).WithCriteria(data_types.Criterion{
		Field: data_types.Field{
			Name:  "reliability",
			Value: "0.99",
		},
		Operation: data_types.Greater,
	})
	simplified, err := data_types.Query(*search).Simplify()
	if err != nil {
		t.Errorf("Failed to simplify query")
	}
	got, err := json.Marshal(simplified)
	if err != nil {
		t.Errorf("failed to marshal request.")
	}
	wanted := "{\"allocated_storage\":10,\"order\":[[\"score\",\"desc\"]],\"reliability\":{\"gt\":\"0.99\"},\"type\":\"bid\"}"

	if !bytes.Equal(got, []byte(wanted)) {
		t.Errorf("got %s; want %s", got, wanted)
	}
}
