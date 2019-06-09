package colormeshop

import (
	"strconv"
)

type CustomerContainer struct {
	Customer Customer `json:"customer"`
}

type CustomersContainer struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	ID                  int64  `json:"id,omitempty"`
	AccountID           string `json:"account_id,omitempty"`
	Name                string `json:"name,omitempty"`
	Furigana            string `json:"furigana,omitempty"`
	Hojin               string `json:"hojin,omitempty"`
	Busho               string `json:"busho,omitempty"`
	Sex                 string `json:"sex,omitempty"`
	Postal              string `json:"postal,omitempty"`
	PrefID              int    `json:"pref_id,omitempty"`
	PrefName            string `json:"pref_name,omitempty"`
	Address1            string `json:"address1,omitempty"`
	Address2            string `json:"address2,omitempty"`
	Mail                string `json:"mail,omitempty"`
	Tel                 string `json:"tel,omitempty"`
	Fax                 string `json:"fax,omitempty"`
	TelMobile           string `json:"tel_mobile,omitempty"`
	Memo                string `json:"memo,omitempty"`
	Points              int64  `json:"points,omitempty"`
	Member              bool   `json:"member,omitempty"`
	SalesCount          int64  `json:"sales_count,omitempty"`
	ReceiveMailMagazine bool   `json:"receive_mail_magazine,omitempty"`
}

func (c *Client) Customer(id int64) (*Customer, error) {
	pid := strconv.FormatInt(id, 10)
	resp, err := c.get("/v1/customers/" + pid + ".json")
	if err != nil {
		return nil, err
	}

	var con CustomerContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Customer, nil
}

func (c *Client) Customers() (*[]Customer, error) {
	resp, err := c.get("/v1/customers.json")
	if err != nil {
		return nil, err
	}

	var con CustomersContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Customers, nil
}
