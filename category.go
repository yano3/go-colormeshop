package colormeshop

import ()

type CategoriesContainer struct {
	Categories []CategoryBig `json:"categories"`
}

type CategoryBig struct {
	IDBig        int64           `json:"id_big,omitempty"`
	IDSmall      int64           `json:"id_small,omitempty"`
	AccountID    string          `json:"account_id,omitempty"`
	Name         string          `json:"name,omitempty"`
	ImageURL     string          `json:"image_url,omitempty"`
	Expl         string          `json:"expl,omitempty"`
	Sort         *int64          `json:"sort,omitempty"`
	DisplayState string          `json:"display_state,omitempty"`
	MakeDate     int             `json:"make_date,omitempty"`
	UpdateDate   int             `json:"update_date,omitempty"`
	Children     []CategorySmall `json:"children,omitempty"`
}

type CategorySmall struct {
	IDBig        int64  `json:"id_big,omitempty"`
	IDSmall      int64  `json:"id_small,omitempty"`
	AccountID    string `json:"account_id,omitempty"`
	Name         string `json:"name,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	Expl         string `json:"expl,omitempty"`
	Sort         *int64 `json:"sort,omitempty"`
	DisplayState string `json:"display_state,omitempty"`
	MakeDate     int    `json:"make_date,omitempty"`
	UpdateDate   int    `json:"update_date,omitempty"`
}

func (c *Client) Categories() (*[]CategoryBig, error) {
	resp, err := c.get("/v1/categories.json")
	if err != nil {
		return nil, err
	}

	var con CategoriesContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Categories, nil
}
