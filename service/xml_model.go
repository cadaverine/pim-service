package service

import "encoding/xml"

type currency struct {
	ID   string `xml:"id,attr"`
	Rate string `xml:"rate,attr"`
}

type currencies struct {
	Currencies []currency `xml:"currency"`
}

type category struct {
	ID       int    `xml:"id,attr"`
	ParentID int    `xml:"parentId,attr"`
	Title    string `xml:",chardata"`
}

type categories struct {
	Categories []category `xml:"category"`
}

type param struct {
	Name  string `xml:"name"`
	Value string `xml:",chardata"`
}

type offer struct {
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
	Param       []param `xml:"param"`
}

type offers struct {
	Offers []offer `xml:"offer"`
}

type shop struct {
	// XMLName    xml.Name   `xml:"shop"`
	Name       string     `xml:"name"`
	Company    string     `xml:"company"`
	URL        string     `xml:"url"`
	Platform   string     `xml:"platform"`
	Currencies currencies `xml:"currencies"`
	Categories categories `xml:"categories"`
	Offers     offers     `xml:"offers"`
}

type catalog struct {
	XMLName xml.Name `xml:"yml_catalog"`
	Date    string   `xml:"date,attr"`
	Shops   []shop   `xml:"shop"`
}
