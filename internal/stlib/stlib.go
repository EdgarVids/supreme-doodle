package stlib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type AgentId struct {
	DataAgentId struct {
		AccountId       string `json:"accountId"`
		Symbol          string `json:"symbol"`
		Headquarters    string `json:"headquarters"`
		Credits         int    `json:"credits"`
		StartingFaction string `json:"startingFaction"`
	} `json:"data"`
}

type ContractArray struct {
	Con []Contract `json:"data"`
}
type DeliverType struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}

type Contract struct {
	Id    string `json:"id:"`
	Type  string `json:"type"`
	Terms struct {
		Deadline string `json:"deadline"`
		Payment  struct {
			OnAccepted  int `json:"onAccepted"`
			OnFulfilled int `json:"onFulfiled"`
		} `json:"payment"`
		Deliver []DeliverType `json:"deliver"`
	} `json:"terms"`
	Accepted         bool   `json:"accepted"`
	Fulfilled        bool   `json:"fulfilled"`
	Expiration       string `json:"expiration"`
	DeadlineToAccept string `json:"deadlineToAccept"`
}

func (a AgentId) Display() {
	fmt.Println(a.DataAgentId.Symbol)
	fmt.Println(a.DataAgentId.Headquarters)
	fmt.Println(a.DataAgentId.Credits)
	fmt.Println(a.DataAgentId.StartingFaction)
}

func ContractListRequest(c *http.Client, u *url.URL, auth string) {
	resource := "v2/my/contracts"
	u.Path = resource
	urlStr := u.String()

	r, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Println(err)
		return
	}
	r.Header.Add("Authorization", auth)

	resp, err := c.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	body, _ := io.ReadAll(resp.Body)

	var contractA ContractArray
	err = json.Unmarshal(body, &contractA)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(contractA)
}

func AgentIdRequest(c *http.Client, u *url.URL, auth string) {
	resource := "v2/my/agent"
	u.Path = resource
	urlStr := u.String()

	r, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		log.Println(err)
		return
	}
	r.Header.Add("Authorization", auth)

	resp, err := c.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	var a AgentId
	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Println(err)
		return
	}
	a.Display()
}
