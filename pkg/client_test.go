package pkg

import (
	"fmt"
	"os"
	"testing"

	"github.com/aalekhpatel07/go-client-vastai/pkg/command"
	"github.com/aalekhpatel07/go-client-vastai/pkg/data_types"
)

func TestClient_SearchOffers(t *testing.T) {
	apikey, found := os.LookupEnv("VASTAI_API_KEY")
	if !found {
		t.Errorf("failed to lookup VASTAI_API_KEY")
	}
	client := DefaultClient(apikey)
	data, err := client.SearchOffers(command.SearchOffersRequest(data_types.Query{
		Criteria: []data_types.Criterion{
			{
				Field: data_types.Field{
					Name:  "reliability",
					Value: "0.99",
				},
				Operation: data_types.Greater,
			},
		},
		Type: data_types.OnDemand,
		Order: []data_types.Order{{
			Field:     "reliability",
			Direction: data_types.Desc,
		}},
		Storage: 5.0,
	}))
	if err != nil {
		t.Errorf("error searching offers:%s", err)
	}

	fmt.Printf("response: %s", data.Offers[0].Webpage)
}
