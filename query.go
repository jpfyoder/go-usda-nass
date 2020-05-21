// Copyright (c) 2020 Joshua Yoder
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.

package nass

// Represents a Query to the NASS.
type Query struct {
    Client      *Client
    Params      map[string]string
}

// Create a new Query object from the client
func NewQuery(client *Client) *Query {
    query := Query{
        Client:         client, // client object
        Params:         make(map[string]string), // list of query parameters
    }
    return &query
}

// Add a filter to the corresponding query object
func (q Query) Filter(param string, value string) Query {
    q.Params[param] = value
    return q
}

// Add a filter with a range operator to the query object
// Valid operator strings include:
// LE: <=
// LT: <
// GT: >
// LIKE: like
// NOT_LIKE: not like
// NE: not equal
func (q Query) FilterRange(param string, value string, operator string) Query {
    q.Params[param + "__" + operator] = value
    return q
}

// Return count of records to be returned by the query if executed
func (q Query) Count() int {
    return q.Client.count_query(q)
}

// Execute the query and return the serialized results
func (q Query) Execute() map[string]interface{} {
    return q.Client.call_query(q)
}
