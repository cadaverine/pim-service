package models

import "encoding/xml"

type Currency struct {
	ID   string `xml:"id,attr"`
	Rate string `xml:"rate,attr"`
}

type Currencies struct {
	Currencies []Currency `xml:"currency"`
}

type Category struct {
	ID       int    `xml:"id,attr"`
	ParentID int    `xml:"parentId,attr" db:"parent_id"`
	Title    string `xml:",chardata"`
}

type Categories struct {
	Categories []Category `xml:"category"`
}

type Param struct {
	Name  string `xml:"name"`
	Value string `xml:",chardata"`
}

type Offer struct {
	ID          string  `xml:"id,attr"`
	Available   bool    `xml:"available,attr"`
	Type        string  `xml:"type,attr"`
	URL         string  `xml:"url"`
	Price       string  `xml:"price"`
	CurrencyID  string  `xml:"currencyId"`
	CategoryID  string  `xml:"categoryId"`
	Name        string  `xml:"name"`
	Picture     string  `xml:"picture"`
	Vendor      string  `xml:"vendor"`
	Description string  `xml:"description"`
	Param       []Param `xml:"param"`
}

type Offers struct {
	Offers []Offer `xml:"offer"`
}

type Shop struct {
	// XMLName    xml.Name   `xml:"shop"`
	Name       string     `xml:"name"`
	Company    string     `xml:"company"`
	URL        string     `xml:"url"`
	Platform   string     `xml:"platform"`
	Currencies Currencies `xml:"currencies"`
	Categories Categories `xml:"categories"`
	Offers     Offers     `xml:"offers"`
}

type Catalog struct {
	XMLName xml.Name `xml:"yml_catalog"`
	Date    string   `xml:"date,attr"`
	Shops   []Shop   `xml:"shop"`
}
