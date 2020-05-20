package nass

import (
        "strings"
        "net/http"
        "encoding/json"
        "log"
)

type Client struct {
    BaseURL     string
    Key         string
}

func NewClient(key string) *Client {
    baseURL := "http://quickstats.nass.usda.gov/api"
    client := Client{
        BaseURL:        baseURL,
        Key:            key,
    }
    return &client
}

func (c Client) ParamValues(param string) map[string]interface{} {
    resp, err := http.Get(c.BaseURL + "/get_param_values/" +
                                    "?key=" + c.Key +
                                    "&param=" + param)
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

func (c Client) count_query(query Query) map[string]interface{} {
    var params strings.Builder
    for k, v := range query.Params {
        params.WriteString("&")
        params.WriteString(k)
        params.WriteString("=")
        params.WriteString(v)
    }
    resp, err := http.Get(c.BaseURL + "/get_counts/" +
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

func (c Client) call_query(query Query) map[string]interface{} {
    var params strings.Builder
    for k, v := range query.Params {
        params.WriteString("&")
        params.WriteString(k)
        params.WriteString("=")
        params.WriteString(v)
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
