package stein

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Stein struct
type Stein struct {
	*http.Client
	BaseURL string
	Option  *Option
}

// Option struct
type Option struct {
	Username string
	Password string
	Timeout  time.Duration
}

type ReadOption struct {
	Search map[string]interface{}
}

// default http timeout
var defHTTPTimeout = 15 * time.Second

// NewClient this function will always be called when the library is in use
func NewClient(baseURL string, option *Option) *Stein {
	httpClient := new(http.Client)

	// checking for default timeout
	if option != nil {
		httpClient.Timeout = option.Timeout
	} else {
		httpClient.Timeout = defHTTPTimeout
	}
	return &Stein{
		Client:  httpClient,
		BaseURL: baseURL,
		Option:  option,
	}
}

// Read Stein API
func (s *Stein) Read(sheetName string, option ReadOption, v interface{}) (status int, err error) {
	var search string
	searchByte, err := json.Marshal(option.Search)
	if err != nil {
		return
	}
	if len(option.Search) != 0 {
		search = fmt.Sprintf("search=%s", string(searchByte))
	}
	status, err = s.Call(http.MethodGet, fmt.Sprintf("%s?%s", sheetName, search), nil, v)
	if err != nil {
		return
	}

	return
}

// Insert to Stein
func (s *Stein) Insert(sheetName string, body interface{}) (status int, res InsertReponse, err error) {
	reqByte, err := json.Marshal(body)
	if err != nil {
		return
	}

	reqBody := bytes.NewReader(reqByte)

	status, err = s.Call(http.MethodPost, sheetName, reqBody, &res)
	if err != nil {
		return
	}

	return
}

// Update rows by condition
func (s *Stein) Update(sheetName string, set interface{}, condition map[string]interface{}) (status int, res UpdateReponse, err error) {
	body := UpdateRequest{
		Condition: condition,
		Set:       set,
	}

	reqByte, err := json.Marshal(body)
	if err != nil {
		return
	}

	reqBody := bytes.NewReader(reqByte)

	status, err = s.Call(http.MethodPut, sheetName, reqBody, &res)
	if err != nil {
		return
	}

	return
}

// Delete rows by condition
func (s *Stein) Delete(sheetName string, condition map[string]interface{}) (status int, res DeleteResponse, err error) {
	body := DeleteRequest{
		Condition: condition,
	}

	reqByte, err := json.Marshal(body)
	if err != nil {
		return
	}

	reqBody := bytes.NewReader(reqByte)

	status, err = s.Call(http.MethodDelete, sheetName, reqBody, &res)
	if err != nil {
		return
	}

	return
}

// ===================== HTTP CLIENT ================================================

// NewRequest send new request
func (s *Stein) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// ExecuteRequest execute request
func (s *Stein) executeRequest(req *http.Request, v interface{}) (status int, err error) {
	res, err := s.Client.Do(req)
	status = res.StatusCode
	if err != nil {
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if v != nil {
		if err = json.Unmarshal(resBody, v); err != nil {
			return
		}
	}
	return
}

// Call the Stein at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occurred, the return of this function is the error
// itself, otherwise nil.
func (s *Stein) Call(method, path string, body io.Reader, v interface{}) (status int, err error) {
	req, err := s.newRequest(method, s.BaseURL+"/"+path, body)
	fmt.Println(path)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return s.executeRequest(req, v)
}

// ===================== END HTTP CLIENT ================================================
