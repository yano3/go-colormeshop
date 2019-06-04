package colormeshop

import ()

type ShopContainer struct {
	Shop Shop `json:"shop"`
}

type Shop struct {
	ID                   string `json:"id,omitempty"`
	State                string `json:"state,omitempty"`
	DomainPlan           string `json:"domain_plan,omitempty"`
	ContractPlan         string `json:"contract_plan,omitempty"`
	ContractStartDate    int    `json:"contract_start_date,omitempty"`
	ContractEndDate      int    `json:"contract_end_date,omitempty"`
	ContractTerm         int    `json:"contract_term,omitempty"`
	ContractPaymentState string `json:"contract_payment_state,omitempty"`
	LastLoginDate        int    `json:"last_login_date,omitempty"`
	SetupDate            int    `json:"setup_date,omitempty"`
	MakeDate             int    `json:"make_date,omitempty"`
	URL                  string `json:"url,omitempty"`
	OpenState            string `json:"open_state,omitempty"`
	LoginID              string `json:"login_id,omitempty"`
	Name1                string `json:"name1,omitempty"`
	Name2                string `json:"name2,omitempty"`
	Name1Kana            string `json:"name1_kana,omitempty"`
	Name2Kana            string `json:"name2_kana,omitempty"`
	Hojin                string `json:"hojin,omitempty"`
	HojinKana            string `json:"hojin_kana,omitempty"`
	UserMail             string `json:"user_mail,omitempty"`
	Tel                  string `json:"tel,omitempty"`
	Fax                  string `json:"fax,omitempty"`
	Postal               string `json:"postal,omitempty"`
	PrefID               int    `json:"pref_id,omitempty"`
	PrefName             string `json:"pref_name,omitempty"`
	Address1             string `json:"address1,omitempty"`
	Address2             string `json:"address2,omitempty"`
	Title                string `json:"title,omitempty"`
	ShopMail1            string `json:"shop_mail_1,omitempty"`
	ShopMail2            string `json:"shop_mail_2,omitempty"`
}

func (c *Client) Shop() (*Shop, error) {
	resp, err := c.get("/v1/shop.json")
	if err != nil {
		return nil, err
	}

	var con ShopContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Shop, nil
}
