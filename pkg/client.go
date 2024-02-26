package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/aalekhpatel07/go-vastai-client/pkg/command"
	"github.com/aalekhpatel07/go-vastai-client/pkg/data_types"
	"io"
	"net/http"
	url2 "net/url"
)

type Client struct {
	BaseURL    string
	httpclient *http.Client
	apikey     string
}

func NewClient(baseURL string, apikey string) *Client {
	return &Client{
		BaseURL:    baseURL,
		apikey:     apikey,
		httpclient: &http.Client{},
	}
}

func DefaultClient(apikey string) *Client {
	return NewClient("https://console.vast.ai/api/v0", apikey)
}

// SearchOffers returns all available machines that can be used
func (c *Client) SearchOffers(cmd command.SearchOffersRequest) (command.SearchOffersResponse, error) {
	offersData := command.SearchOffersResponse{}

	path := fmt.Sprintf("%s/bundles/", c.BaseURL)
	url, err := url2.Parse(path)
	if err != nil {
		return offersData, err
	}
	simplified, err := data_types.Query(cmd).Simplify()
	if err != nil {
		return offersData, err
	}
	body, err := json.Marshal(simplified)
	if err != nil {
		return offersData, err
	}

	bodyString := string(body)
	q := url.Query()
	q.Set("q", bodyString)
	q.Set("api_key", c.apikey)
	fullUrl := fmt.Sprintf("%s?%s", url, q.Encode())

	resp, err := http.Get(fullUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if err != nil {
		return offersData, err
	}

	if resp.StatusCode != 200 {
		errorData := data_types.VastAIError{}
		err = json.NewDecoder(resp.Body).Decode(&errorData)
		if err != nil {
			return offersData, err
		}
		return offersData, errorData
	}
	err = json.NewDecoder(resp.Body).Decode(&offersData)
	if err != nil {
		return offersData, err
	}

	return offersData, nil
}
