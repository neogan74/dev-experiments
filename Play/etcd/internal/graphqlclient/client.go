package graphqlclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is a thin helper around the GraphQL HTTP transport.
type Client struct {
	endpoint   string
	httpClient *http.Client
	authToken  string
}

// NewClient constructs a GraphQL client targeting the given endpoint.
// The authToken (if set) is added as a bearer token on every request.
func NewClient(endpoint, authToken string, timeout time.Duration) *Client {
	return &Client{
		endpoint: endpoint,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		authToken: authToken,
	}
}

// Request bundles the GraphQL query and variables as expected by GraphQL over HTTP.
type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// Response captures the GraphQL http response payload.
type Response struct {
	Data   json.RawMessage `json:"data"`
	Errors []struct {
		Message    string                 `json:"message"`
		Path       []interface{}          `json:"path,omitempty"`
		Locations  []map[string]int       `json:"locations,omitempty"`
		Extensions map[string]interface{} `json:"extensions,omitempty"`
	} `json:"errors,omitempty"`
}

// Execute issues the GraphQL request and unmarshals the data field into out.
func (c *Client) Execute(ctx context.Context, query string, variables map[string]interface{}, out interface{}) error {
	payload := Request{
		Query:     query,
		Variables: variables,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("encode graphql payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.authToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.authToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("call graphql endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
		return fmt.Errorf("graphql request failed: status=%d body=%s", resp.StatusCode, string(snippet))
	}

	var gqlResp Response
	if err = json.NewDecoder(resp.Body).Decode(&gqlResp); err != nil {
		return fmt.Errorf("decode graphql response: %w", err)
	}

	if len(gqlResp.Errors) > 0 {
		errMessages := make([]string, 0, len(gqlResp.Errors))
		for _, item := range gqlResp.Errors {
			errMessages = append(errMessages, item.Message)
		}
		return errors.New("graphql errors: " + fmt.Sprint(errMessages))
	}

	if out == nil {
		return nil
	}

	if err = json.Unmarshal(gqlResp.Data, out); err != nil {
		return fmt.Errorf("decode graphql data: %w", err)
	}

	return nil
}
