package main

import (
	//	"encoding/json"
	"fmt"
	//	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/EdgarVids/supreme-doodle/internal/stlib"
)

func main() {
	token := os.Getenv("SPACETRADERS")
	fmt.Println("Space Traders Client Practice")
	// fmt.Println("Token:", token)
	c := &http.Client{Timeout: time.Duration(1) * time.Second}
	auth := fmt.Sprintf("Bearer %s", token)
	apiUrl := "https://api.spacetraders.io"
	//	resource := "/v2/my/agent/"
	u, _ := url.ParseRequestURI(apiUrl)
	//	u.Path = resource
	//	urlStr := u.String()

	//	r, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	//	r.Header.Add("Authorization", auth)

	//	resp, _ := c.Do(r)
	//	body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// fmt.Println(resp.Status)
	stlib.ContractListRequest(c, u, auth)
	stlib.AgentIdRequest(c, u, auth)

}
