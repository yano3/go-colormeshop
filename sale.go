package colormeshop

import (
	"strconv"
)

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

type SaleContainer struct {
	Sale Sale `json:"sale"`
}

type SalesContainer struct {
	Sales []Sale `json:"sales"`
}

type Sale struct {
	ID                    int64          `json:"id,omitempty"`
	AccountID             string         `json:"account_id,omitempty"`
	MakeDate              int            `json:"make_date,omitempty"`
	UpdateDate            int            `json:"update_date,omitempty"`
	Memo                  string         `json:"memo,omitempty"`
	PaymantID             int64          `json:"payment_id,omitempty"`
	Paid                  bool           `json:"paid,omitempty"`
	Delivered             bool           `json:"delivered,omitempty"`
	Canceled              bool           `json:"canceled,omitempty"`
	AccepptedMailState    string         `json:"accepted_mail_state,omitempty"`
	PaidMailState         string         `json:"paid_mail_state,omitempty"`
	DeliveredMailState    string         `json:"delivered_mail_state,omitempty"`
	PointState            string         `json:"point_state,omitempty"`
	GMOPointState         string         `json:"gmo_point_state,omitempty"`
	ProductTotalPrice     int64          `json:"product_total_price,omitempty"`
	DeliveryTotalCharge   int64          `json:"delivery_total_charge,omitempty"`
	Tax                   int64          `json:"tax,omitempty"`
	Fee                   int64          `json:"fee,omitempty"`
	NoshiTotalCharge      int64          `json:"noshi_total_charge,omitempty"`
	CardTotalCharge       int64          `json:"card_total_charge,omitempty"`
	WrappingTotalCharge   int64          `json:"wrapping_total_charge,omitempty"`
	PointDiscount         int64          `json:"point_discount,omitempty"`
	GMOPointDiscount      int64          `json:"gmo_point_discount,omitempty"`
	OtherDiscount         int64          `json:"other_discount,omitempty"`
	OtherDiscount_name    string         `json:"other_discount_name,omitempty"`
	TotalPrice            int64          `json:"total_price,omitempty"`
	GrantedPoints         int64          `json:"granted_points,omitempty"`
	UsePoints             int64          `json:"use_points,omitempty"`
	GrantedGMOPoints      int64          `json:"granted_gmo_points,omitempty"`
	UseGMOPoints          int64          `json:"use_gmo_points,omitempty"`
	AcceptedMailSentDate  *int64         `json:"accepted_mail_sent_date,omitempty"`
	PaidMailSentDate      *int64         `json:"paid_mail_sent_date,omitempty"`
	DeliveredMailSentDate *int64         `json:"delivered_mail_sent_date,omitempty"`
	Customer              Customer       `json:"customer,omitempty"`
	Details               []SaleDetail   `json:"details,omitempty"`
	Deliveries            []SaleDelivery `json:"sale_deliveries,omitempty"`
}

type SaleDetail struct {
	ID                       int64  `json:"id,omitempty"`
	SaleID                   int64  `json:"sale_id,omitempty"`
	AccountID                string `json:"account_id,omitempty"`
	ProductID                int64  `json:"product_id,omitempty"`
	SaleDeliveryID           int64  `json:"sale_delivery_id,omitempty"`
	Option1Value             string `json:"option1_value,omitempty"`
	Option2Value             string `json:"option2_value,omitempty"`
	Option1Index             int64  `json:"option1_index,omitempty"`
	Option2Index             int64  `json:"option2_index,omitempty"`
	ProductModelNumber       string `json:"product_model_number,omitempty"`
	ProductName              string `json:"product_name,omitempty"`
	ProductCost              *int64 `json:"product_cost,omitempty"`
	ProductImageURL          string `json:"product_image_url,omitempty"`
	ProductThumbnailImageURL string `json:"product_thumbnail_image_url,omitempty"`
	Price                    *int64 `json:"price,omitempty"`
	PriceWithTax             *int64 `json:"price_with_tax,omitempty"`
	ProductNum               int64  `json:"product_num,omitempty"`
	Unit                     string `json:"unit,omitempty"`
	SubtotalPrice            *int64 `json:"subtotal_price,omitempty"`
}

type SaleDelivery struct {
	ID              int64   `json:"id,omitempty"`
	SaleID          int64   `json:"sale_id,omitempty"`
	AccountID       string  `json:"account_id,omitempty"`
	DeliveryID      int64   `json:"delivery_id,omitempty"`
	DetailsIDs      []int64 `json:"detail_ids,omitempty"`
	Memo            string  `json:"memo,omitempty"`
	Delivered       bool    `json:"delivered,omitempty"`
	Name            string  `json:"name,omitempty"`
	Furigana        string  `json:"furigana,omitempty"`
	Postal          string  `json:"postal,omitempty"`
	PrefID          int     `json:"pref_id,omitempty"`
	PrefName        string  `json:"pref_name,omitempty"`
	Address1        string  `json:"address1,omitempty"`
	Address2        string  `json:"address2,omitempty"`
	Tel             string  `json:"tel,omitempty"`
	PreferredDate   string  `json:"preferred_date,omitempty"`
	PreferredPeriod string  `json:"preferred_period,omitempty"`
	SlipNumber      string  `json:"slip_number,omitempty"`
	TrackingURL     string  `json:"tracking_url,omitempty"`
	NoshiText       string  `json:"noshi_text,omitempty"`
	NoshiCharge     *int64  `json:"noshi_charge,omitempty"`
	CardName        string  `json:"card_name,omitempty"`
	CardText        string  `json:"card_text,omitempty"`
	CardCharge      *int64  `json:"card_charge,omitempty"`
	WrappingText    string  `json:"wrapping_text,omitempty"`
	WrappingCharge  *int64  `json:"wrapping_charge,omitempty"`
	DeliveryCharge  *int64  `json:"delivery_charge,omitempty"`
	TotalCharge     *int64  `json:"total_charge,omitempty"`
}

func (c *Client) Sale(id int64) (*Sale, error) {
	pid := strconv.FormatInt(id, 10)
	resp, err := c.get("/v1/sales/" + pid + ".json")
	if err != nil {
		return nil, err
	}

	var con SaleContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}
	return &con.Sale, nil
}

func (c *Client) Sales() (*[]Sale, error) {
	resp, err := c.get("/v1/sales.json")
	if err != nil {
		return nil, err
	}

	var con SalesContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}
	return &con.Sales, nil
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
