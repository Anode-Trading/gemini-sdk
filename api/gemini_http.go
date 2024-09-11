package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetSecurities() ([]string, error) {
	var securities []string
	url := fmt.Sprintf("%s%s%s", httpBaseUrl, v1, symbols)
	client := NewHttpClient(2)

	symbolResponse, err := client.Get(url)
	if err != nil || symbolResponse == nil || symbolResponse.Body == nil || symbolResponse.StatusCode != http.StatusOK {
		return securities, fmt.Errorf("cannot get symbol details: %s", err)
	}

	defer symbolResponse.Body.Close()
	symbolResponseBody, err := io.ReadAll(symbolResponse.Body)
	if err != nil {
		return securities, fmt.Errorf("err reading symbol details: %s", err)
	}

	err = json.Unmarshal(symbolResponseBody, &securities)
	if err != nil {
		return securities, fmt.Errorf("err creating symbol details: %s", err)
	}
	return securities, nil
}
func GetSecurityInfo(pair string) (SymbolDetails, error) {
	var symbolDetails SymbolDetails
	url := fmt.Sprintf("%s%s%s", httpBaseUrl, v1, fmt.Sprintf(symbolsDetails, pair))
	client := NewHttpClient(2)

	symbolDetailsResponse, err := client.Get(url)
	if err != nil || symbolDetailsResponse == nil || symbolDetailsResponse.Body == nil || symbolDetailsResponse.StatusCode != http.StatusOK {
		return symbolDetails, fmt.Errorf("cannot get symbol details: %s", err)
	}
	defer symbolDetailsResponse.Body.Close()

	symbolDetailResponseBody, err := io.ReadAll(symbolDetailsResponse.Body)
	if err != nil {
		return symbolDetails, fmt.Errorf("err reading symbol details: %s", err)
	}

	err = json.Unmarshal(symbolDetailResponseBody, &symbolDetails)
	if err != nil {
		return symbolDetails, fmt.Errorf("err creating symbol details: %s", err)
	}
	return symbolDetails, nil
}
