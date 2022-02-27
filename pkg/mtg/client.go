package mtg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Ruling struct {
	// TODO should parse dates
	Date string `json:"date"`
	Text string `json:"text"`
}

type CardQuery struct {
	Page     int
	PageSize int
	OrderBy  string
}

func (c *CardQuery) MarshalValues() url.Values {
	values := make(url.Values)

	if c.Page != 0 {
		values["page"] = []string{strconv.Itoa(c.Page)}
	}

	if c.PageSize != 0 {
		values["pageSize"] = []string{strconv.Itoa(c.PageSize)}
	}

	if c.OrderBy != "" {
		values["orderBy"] = []string{c.OrderBy}
	}
	return values
}

func WithOrderBy(orderField string) SearchCardsOption {
	return func(query *CardQuery) error {
		if query.OrderBy != "" {
			return errors.New("query already has orderBy set")
		}
		query.OrderBy = orderField
		return nil
	}
}

func WithPage(pageID int) SearchCardsOption {
	return func(query *CardQuery) error {
		if query.Page != 0 {
			return errors.New("query already has a page ID set")
		}
		query.Page = pageID
		return nil
	}
}

type Client struct{}

type SearchCardsOption func(*CardQuery) error

type SearchCardsResponse struct {
	Cards []Card `json:"cards"`
}

// | is a logical OR, comma is a logical AND
func (c *Client) SearchCards(opts ...SearchCardsOption) ([]Card, error) {
	var ret []Card
	query := &CardQuery{}

	var err error

	for _, v := range opts {
		err = v(query)
		if err != nil {
			return ret, fmt.Errorf("failed to apply option to client: %+w", err)
		}
	}

	req, err := http.NewRequest("GET", "https://api.magicthegathering.io/v1/cards", bytes.NewBuffer([]byte{}))
	if err != nil {
		return ret, fmt.Errorf("failed to create GET request to search cards: %+w", err)
	}

	req.URL.RawQuery = query.MarshalValues().Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, fmt.Errorf("failed to perform HTTP GET request to search cards: %+w", err)
	}
	defer resp.Body.Close()

	var cardsResponse SearchCardsResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, fmt.Errorf("failed to read response body: %+w", err)
	}

	err = json.Unmarshal(body, &cardsResponse)
	if err != nil {
		return ret, fmt.Errorf("failed to unmarshal response: %+w", err)
	}

	return cardsResponse.Cards, nil
}

type ListSetsResponse struct {
	Sets []Set `json:"sets"`
}

func (c *Client) ListSets() ([]Set, error) {
	var ret []Set

	req, err := http.NewRequest("GET", "https://api.magicthegathering.io/v1/sets", bytes.NewBuffer([]byte{}))
	if err != nil {
		return ret, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()

	var setResponse ListSetsResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &setResponse)
	if err != nil {
		return ret, err
	}

	return setResponse.Sets, nil
}

type ListTypesResponse struct {
	Types []string `json:"types"`
}

func (c *Client) ListTypes() ([]string, error) {
	// TODO error handling
	resp, err := http.Get("https://api.magicthegathering.io/v1/types")
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var unmarshalledResp ListTypesResponse

	err = json.Unmarshal(body, &unmarshalledResp)
	if err != nil {
		return []string{}, err
	}

	return unmarshalledResp.Types, nil
}

type ListSubtypesResponse struct {
	Subtypes []string `json:"subtypes"`
}

func (c *Client) ListSubTypes() ([]string, error) {
	// TODO error handling
	resp, err := http.Get("https://api.magicthegathering.io/v1/subtypes")
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var unmarshalledResp ListSubtypesResponse

	err = json.Unmarshal(body, &unmarshalledResp)
	if err != nil {
		return []string{}, err
	}

	return unmarshalledResp.Subtypes, nil
}

type ListSupertypesResponse struct {
	Supertypes []string `json:"supertypes"`
}

func (c *Client) ListSuperTypes() ([]string, error) {
	// TODO error handling
	resp, err := http.Get("https://api.magicthegathering.io/v1/supertypes")
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	var unmarshalledResp ListSupertypesResponse

	err = json.Unmarshal(body, &unmarshalledResp)
	if err != nil {
		return []string{}, err
	}

	return unmarshalledResp.Supertypes, nil
}
