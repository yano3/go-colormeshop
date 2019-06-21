package colormeshop

import ()

type GroupsContainer struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	ID           int64  `json:"id,omitempty"`
	AccountID    string `json:"account_id,omitempty"`
	Name         string `json:"name,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	Expl         string `json:"expl,omitempty"`
	Sort         *int64 `json:"sort,omitempty"`
	DisplayState string `json:"display_state,omitempty"`
}

func (c *Client) Groups() (*[]Group, error) {
	resp, err := c.get("/v1/groups.json")
	if err != nil {
		return nil, err
	}

	var con GroupsContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Groups, nil
}
