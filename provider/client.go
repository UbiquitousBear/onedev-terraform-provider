package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	user       string
	authToken  string
	httpClient *http.Client
}

func NewClient(hostname string, user string, token string) *Client {
	return &Client{
		hostname:   hostname,
		user:       user,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAllProjects() (*map[string]Project, error) {
	body, err := c.httpRequest("projects?offset=0&count=100", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	projects := map[string]Project{}
	err = json.NewDecoder(body).Decode(&projects)
	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetProject(id int) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("projects/%d", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	item := &Project{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) NewProject(project Project) (*Project, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("projects", "POST", buf)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	item := &Project{}
	json.NewDecoder(body).Decode(item)

	return &project, nil
}

func (c *Client) UpdateProject(project *Project) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("projects/%d", project.Id), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProject(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("projects/%d", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/api/%s", c.hostname, path)
}
