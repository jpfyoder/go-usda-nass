package nass

type Query struct {
    Client      *Client
    Params      map[string]string
}

func NewQuery(client *Client) *Query {
    query := Query{
        Client:         client,
        Params:         make(map[string]string),
    }
    return &query
}

func (q Query) Filter(param string, value string, operator string) Query {
    if operator == "" {
        q.Params[param] = value
    } else {
        q.Params[param + "__" + operator] = value
    }
    return q
}

func (q Query) Count() map[string]interface{} {
    return q.Client.count_query(q)
}

func (q Query) Execute() map[string]interface{} {
    return q.Client.call_query(q)
}
