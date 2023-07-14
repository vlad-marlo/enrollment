package httpclient

import (
	"encoding/json"
	"fmt"
	"github.com/vlad-marlo/enrollment/internal/model"
	"net/http"
)

func (cli *Client) CreateRecord(user, msgType string) (*model.CreateRecordResponse, error) {
	resp, err := cli.client.R().SetFormData(map[string]string{
		"user":     user,
		"msg_type": msgType,
	}).Post(createRoute)
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	response := new(model.CreateRecordResponse)

	if err = json.Unmarshal(resp.Body(), response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %w", err)
	}
	return response, nil
}

func (cli *Client) GetRecordByID(id int64) (*model.GetRecordResponse, error) {
	resp, err := cli.client.R().Get(fmt.Sprintf(getOneRoute, id))
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status())
	}

	response := new(model.GetRecordResponse)

	if err = json.Unmarshal(resp.Body(), response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %w", err)
	}
	return response, nil
}

func (cli *Client) GetRecordsByUser(user string) (*model.GetUserRecordsResponse, error) {
	resp, err := cli.client.R().Get(fmt.Sprintf(getUsersRoute, user))
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status())
	}

	response := new(model.GetUserRecordsResponse)

	if err = json.Unmarshal(resp.Body(), response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %w", err)
	}
	return response, nil
}

func (cli *Client) GetAll() (*model.GetAllRecordsResponse, error) {
	resp, err := cli.client.R().Get(getAllRoute)
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status())
	}

	response := new(model.GetAllRecordsResponse)

	if err = json.Unmarshal(resp.Body(), response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %w", err)
	}
	return response, nil
}
