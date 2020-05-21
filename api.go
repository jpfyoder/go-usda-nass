// Copyright (c) 2020 Joshua Yoder
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

// Go wrapper for USDA NASS API
// See the USDA NASS QuickStats API docs for more usage details:
// https://quickstats.nass.usda.gov/api
package nass

import (
        "strings"
        "net/http"
        "encoding/json"
        "log"
)

// Represents a client to the USDA
type Client struct {
    BaseURL     string // Base URL for USDA NASS
    Key         string // API key
}

// Create a new client bearing a given key
func NewClient(key string) *Client {
    baseURL := "http://quickstats.nass.usda.gov/api"
    client := Client{
        BaseURL:        baseURL,
        Key:            key,
    }
    return &client
}

// Return all possible values of a given parameter
func (c Client) ParamValues(param string) []string {
    resp, err := http.Get(c.BaseURL + "/get_param_values/" +
                                    "?key=" + c.Key +
                                    "&param=" + param)
    if err != nil {
        log.Fatal(err)
    }
    var data map[string][]string
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }
    return data[param]
}

// Return number of records that will be retrieved by a given query
func (c Client) count_query(query Query) int {
    var params strings.Builder
    for k, v := range query.Params {
        params.WriteString("&" + k + "=" + v)
    }
    resp, err := http.Get(c.BaseURL + "/get_counts/" +
                                   "?key=" + c.Key +
                                   params.String())
    if err != nil {
        log.Fatal(err)
    }
    var data map[string]int
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }
    return data["count"]
}

// Return the result of the query
func (c Client) call_query(query Query) map[string]interface{} {
    var params strings.Builder
    for k, v := range query.Params {
        params.WriteString("&" + k + "=" + v)
    }
    resp, err := http.Get(c.BaseURL + "/api_GET/" +
                                   "?key=" + c.Key +
                                   params.String())
    if err != nil {
        log.Fatal(err)
    }
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&data)
    if err != nil {
        log.Fatal(err)
    }
    return data
}
