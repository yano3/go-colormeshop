package colormeshop

import (
	"strconv"
)

type ProductContainer struct {
	Product Product `json:"product"`
}

type ProductsContainer struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID                int64     `json:"id,omitempty"`
	AccountID         string    `json:"account_id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Category          Category  `json:category,omitempty"`
	GroupIds          []int64   `json:"group_ids,omitempty"`
	DisplayState      string    `json:"display_state,omitempty"`
	SalesPrice        *int64    `json:"sales_price,omitempty"`
	Price             *int64    `json:"price,omitempty"`
	MembersPrice      *int64    `json:"members_price,omitempty"`
	Cost              *int64    `json:"cost,omitempty"`
	CoolCharge        *int64    `json:"cool_charge,omitempty"`
	DeliveryCharge    *int64    `json:"delivery_charge,omitempty"`
	StockManaged      bool      `json:"stock_managed,omitempty"`
	Stocks            *int64    `json:"stocks,omitempty"`
	MaxNum            *int64    `json:"max_num,omitempty"`
	MinNum            *int64    `json:"min_num,omitempty"`
	SaleStartDate     *int64    `json:"sale_start_date,omitempty"`
	SaleEndDate       *int64    `json:"sale_end_date,omitempty"`
	Unit              *int64    `json:"unit,omitempty"`
	Weight            *int64    `json:"weight,omitempty"`
	SoldOutDisplay    bool      `json:"soldout_display,omitempty"`
	FewNum            *int64    `json:",omitempty"`
	Sort              *int64    `json:"sort,omitempty"`
	Expl              string    `json:"expl,omitempty"`
	SimpleExpl        string    `json:"simple_expl,omitempty"`
	SmartPhoneExpl    string    `json:"smartphone_expl,omitempty"`
	MakeDate          int       `json:"make_date,omitempty"`
	UpdateDate        int       `json:"update_date,omitempty"`
	ImageURL          string    `json:"image_url,omitempty"`
	ThumbnailImageURL string    `json:"thumbnail_image_url,omitempty"`
	Images            []Image   `json:images,omitempty"`
	Options           []Option  `json:options,omitempty"`
	Variants          []Variant `json:variants,omitempty"`
	Pickups           []Pickup  `json:pickups,omitempty"`
}

type Category struct {
	IDBig   int64 `json:"id_big,omitempty"`
	IDSmall int64 `json:"id_small,omitempty"`
}

type Image struct {
	Src      string `json:"src,omitempty"`
	Position int64  `json:"posision,omitempty"`
}

type Option struct {
	ID         int64  `json:"id,omitempty"`
	ProductID  int64  `json:"product_id,omitempty"`
	AccountID  string `json:"account_id,omitempty"`
	Name       string `json:"name,omitempty"`
	MakeDate   int    `json:"make_date,omitempty"`
	UpdateDate int    `json:"update_date,omitempty"`
	Values     []string
}

type Variant struct {
	ProductID                     int64   `json:"product_id,omitempty"`
	AccountID                     string  `json:"account_id,omitempty"`
	Option1Value                  string  `json:"option1_value,omitempty"`
	Option2Value                  string  `json:"option2_value,omitempty"`
	Title                         string  `json:"title,omitempty"`
	Stocks                        *int64  `json:"stocks,omitempty"`
	ModelNumber                   *string `json:"model_number,omitempty"`
	FewNum                        *int64  `json:"few_num,omitempty"`
	OptionPrice                   *int64  `json:"option_price,omitempty"`
	OptionPriceIncludingTax       *int64  `json:"option_price_including_tax,omitempty"`
	OptionPriceTax                *int64  `json:"option_price_tax,omitempty"`
	OptionMemberPrice             *int64  `json:"option_members_price,omitempty"`
	OptionMemberPriceIncludingTax *int64  `json:"option_members_price_including_tax,omitempty"`
	OptionMemberPriceTax          *int64  `json:"option_members_price_tax,omitempty"`
	MakeDate                      int     `json:"make_date,omitempty"`
	UpdateDate                    int     `json:"update_date,omitempty"`
}

type Pickup struct {
	PickupType int64 `json:"pickup_type,omitempty"`
	OrderNum   int64 `json:"order_num,omitempty"`
	MakeDate   int   `json:"make_date,omitempty"`
	UpdateDate int   `json:"update_date,omitempty"`
}

func (c *Client) Product(id int64) (*Product, error) {
	pid := strconv.FormatInt(id, 10)
	resp, err := c.get("/v1/products/" + pid + ".json")
	if err != nil {
		return nil, err
	}

	var con ProductContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Product, nil
}

func (c *Client) Products() (*[]Product, error) {
	resp, err := c.get("/v1/products.json")
	if err != nil {
		return nil, err
	}

	var con ProductsContainer
	if err := decodeJSON(resp, &con); err != nil {
		return nil, err
	}

	return &con.Products, nil
}
