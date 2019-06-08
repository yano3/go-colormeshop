package colormeshop

import ()

type SalesStatusContainer struct {
	SalesStatus SalesStatus `json:"sales_stat"`
}

type SalesStatus struct {
	AccountID       string `json:"account_id,omitempty"`
	Date            int    `json:"date,omitempty"`
	AmountToday     int64  `json:"amount_today,omitempty"`
	CountToday      int64  `json:"count_today,omitempty"`
	AmountLast7Days int64  `json:"amount_last_7days,omitempty"`
	CountLast7Days  int64  `json:"count_last_7days,omitempty"`
	AmountThisMonth int64  `json:"amount_this_month,omitempty"`
	CountThisMonth  int64  `json:"count_this_month,omitempty"`
}

func (c *Client) SalesStatus() (*SalesStatus, error) {
	resp, err := c.get("/v1/sales/stat.json")
	if err != nil {
		return nil, err
	}

	var con SalesStatusContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.SalesStatus, nil
}
