package command

import (
	"github.com/aalekhpatel07/go-client-vastai/pkg/data_types"
)

type SearchOffersRequest data_types.Query

type SearchOffersResponse struct {
	Offers []data_types.Instance `json:"offers"`
}

func NewSearchOffersRequest() *SearchOffersRequest {
	search := SearchOffersRequest{
		Criteria: make([]data_types.Criterion, 0),
		Type:     data_types.OnDemand,
		Order:    make([]data_types.Order, 0),
		Storage:  5,
	}
	return &search
}

func (search *SearchOffersRequest) WithInstanceType(instanceType data_types.InstanceType) *SearchOffersRequest {
	search.Type = instanceType
	return search
}

func (search *SearchOffersRequest) WithStorage(storage float64) *SearchOffersRequest {
	search.Storage = storage
	return search
}

func (search *SearchOffersRequest) WithOrder(order data_types.Order) *SearchOffersRequest {
	search.Order = append(search.Order, order)
	return search
}

func (search *SearchOffersRequest) WithCriteria(criteria data_types.Criterion) *SearchOffersRequest {
	search.Criteria = append(search.Criteria, criteria)
	return search
}
