package api

// Move this file to internal fightdataapi package once testing of get request completed.
import (
        "io"
        "net/http"
        "fmt"
)


// Get request to scrape for HTML data from wikis. Needs inputs to be updated to be selctable later.
func (c *Client) FightDataGet(baseURL string) ([]byte, error) {

        // URL update hardcoded for testing. Need to update to change based on user input later.
    url := DustloopURL + "BBCF/" + "Kokonoe"

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
        if res.StatusCode > 299 {
                return nil, fmt.Errorf("Response failed with statuscode: %d\n", res.StatusCode)
        }

        data, err := io.ReadAll(res.Body)
        if err != nil {
                return nil, err
        }

        return data, nil
}
