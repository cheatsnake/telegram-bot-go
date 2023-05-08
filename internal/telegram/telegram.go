package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	token string
	client http.Client
}

func New(token string) *Client {
	return &Client{
		token: token,
		client: http.Client{},
	}
}

func (c *Client) Updates(offset, limit int) ([]Update, error) {
	query := url.Values{}
	query.Add(offsetQuery, strconv.Itoa(offset))
	query.Add(limitQuery, strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}

	var resp Response[[]Update]

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	query := url.Values{}
	query.Add(chatIDQuery, strconv.Itoa(chatID))
	query.Add(textQuery, text)

	_, err := c.doRequest(sendMessageMethod, query)
	if err != nil {
		return fmt.Errorf("send message failed: %w", err)
	}

	return nil
}

func (c *Client) GetMe() (User, error) {
	data, err := c.doRequest(getMeMethod, url.Values{})
	if (err != nil) {
		return User{}, fmt.Errorf("bot authentication failed: %w", err)
	}

	var resp Response[User]

	if err := json.Unmarshal(data, &resp); err != nil {
		return resp.Result, err
	}

	if !resp.Ok {
		return resp.Result, fmt.Errorf("bot authorization failed")
	}

	return resp.Result, nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	reqUrl := url.URL{
		Scheme: "https",
		Host: apiHost,
		Path: path.Join(botPrefix + c.token, method),
	}

	req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("body parse failed: %w", err)
	}

	return body, nil
}
