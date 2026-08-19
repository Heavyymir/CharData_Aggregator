package api

// Move this file to internal fightdataapi package once testing of get request completed.
import (
	"fmt"
	"io"
	"net/http"
)

// Get request to scrape for HTML data from wikis. Needs inputs to be updated to be selctable later.
func (c *Client) Fetch(url string) ([]byte, error) {

	// Check to see if the URL is present inside the Cache
	data, ok := c.Cache.Get(url)
	if ok && len(data) > 0 {
		return data, nil
	}

	//Basic request logic using input URL
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Read the HTTP Body. If empty, return an error.
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty response body from %s", url)
	}

	// Check res.StatusCode. If it is above 299, return an error with the status and body of the response.
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf(
		"response failed: status= %s body= %q", 
		res.Status, 
		string(data[:min(len(data), 500)]),
		)
	}

	// Add data from the request to the Cache.
	c.Cache.Add(url, data)

	return data, nil
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
