package lightlytics

import "context"
import "errors"
import "fmt"
import "github.com/machinebox/graphql"

type Client struct {
    graphqlClient *graphql.Client
    Token string
    Workspace string
}

func NewClient(host, username, password, workspace_id *string) (*Client, error) {

    c := Client{
        graphqlClient: graphql.NewClient(fmt.Sprintf("%s/graphql", *host)),
    }

    c.Token = ""
    if workspace_id != nil{
        c.Workspace = *workspace_id
    }

    err := c.authenticate(username, password)

    if err != nil {
        return nil, err
    }

    return &c, nil
}

func (c *Client) authenticate(username, password *string) (error) {
    data, err := c.doRequest(`
        mutation ($creds: Credentials) {
            login (credentials:$creds) {
                access_token
            }
        }
    `, map[string]interface{}{
        "creds": map[string]interface{}{
            "email": username,
            "password": password}})

    if err != nil {
        return err
    }

    token, valid := data["login"].(map[string]interface{})["access_token"].(string)

    if !valid {
        return errors.New("oops")
    }

    c.Token = token

    return nil
}  

func (c *Client) doRequest(query string, variables map[string]interface{}) (map[string]interface{}, error) {
    req := graphql.NewRequest(query)

    if c.Token != "" {
        req.Header.Set("Authorization", "Bearer " + c.Token)
    }
    if c.Workspace != "" {
        req.Header.Set("customer", c.Workspace)
    }

    if variables != nil {
        for key, value := range variables {
            req.Var(key, value)
        }
    }

    // define a Context for the request
    ctx := context.Background()

    // run it and capture the response
    var data map[string]interface{}
    if err := c.graphqlClient.Run(ctx, req, &data); err != nil {
        return nil, err
    }

    return data, nil
}