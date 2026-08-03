package api

// Move this file to internal fightdataapi package once testing of get request completed.
import (
        "io"
        "net/http"
        "fmt"
)


// Get request to scrape for HTML data from wikis. Needs inputs to be updated to be selctable later.
func (c *Client) Fetch(url string) ([]byte, error) {

	// Check to see if the URL is present inside the Cache
    data, ok := c.Cache.Get(url)
	if ok {
        return data, nil
 	}
	        
    //Basic request logic using input URL
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
    	return nil, err
   	}

		
   	res, err := c.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

   	defer res.Body.Close()

   	// Check res.StatusCode. If it is above 299, return an error with code.
    if res.StatusCode < 299 || res.StatusCode >= 300 {
   		return nil, fmt.Errorf("Response failed with statuscode: %d\n", res.StatusCode)
    }

	// Read the HTTP Body. If empty, return an error.
    data, err = io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

	// Add data from the request to the Cache.
    c.Cache.Add(url, data)

   	return data, nil
}
