package Entities

import "encoding/xml"

type Account struct {
	User struct {
		IDUser                 int          `json:"idUser"`
		Username               string       `json:"username"`
		Country                string       `json:"country"`
		IsCommercial           int          `json:"isCommercial"`
		MaySell                bool         `json:"maySell"`
		SellerActivation       int          `json:"sellerActivation"`
		RiskGroup              int          `json:"riskGroup"`
		LossPercentage         string       `json:"lossPercentage"`
		Reputation             int          `json:"reputation"`
		ShipsFast              int          `json:"shipsFast"`
		SellCount              int          `json:"sellCount"`
		SoldItems              int          `json:"soldItems"`
		AvgShippingTime        int          `json:"avgShippingTime"`
		OnVacation             bool         `json:"onVacation"`
		IDDisplayLanguage      int          `json:"idDisplayLanguage"`
		Name                   Name         `json:"name"`
		Home                   HomeAddress  `json:"homeAddress"`
		Email                  string       `json:"email"`
		PhoneNumber            string       `json:"phoneNumber"`
		Vat                    string       `json:"vat"`
		LegalInformation       string       `json:"legalInformation"`
		RegisterDate           string       `json:"registerDate"`
		IsActivated            bool         `json:"isActivated"`
		MoneyDetails           MoneyDetails `json:"moneyDetails"`
		BankAccount            BankAccount  `json:"bankAccount"`
		ArticlesInShoppingCart int          `json:"articlesInShoppingCart"`
		UnreadMessages         int          `json:"unreadMessages"`
	} `json:"account"`
	Links Links `json:"links"`
}

type ThreadPool struct {
	Thread []ThreadView `json:"thread"`
	Links  Links        `json:"links"`
}
type ThreadView struct {
	Partner        Partner `json:"partner"`
	Message        Message `json:"message"`
	UnreadMessages int     `json:"unreadMessages"`
	Links          Links   `json:"links"`
}
type Thread struct {
	Partner        Partner   `json:"partner"`
	Message        []Message `json:"message"`
	UnreadMessages int       `json:"unreadMessages"`
	Links          Links     `json:"links"`
}

type Message struct {
	IDMessage int    `json:"idMessage"`
	IsSending bool   `json:"isSending"`
	Date      string `json:"date"`
	Text      string `json:"text"`
	Unread    bool   `json:"unread"`
}

type Partner struct {
	IDUser           int         `json:"idUser"`
	Username         string      `json:"username"`
	RegistrationDate string      `json:"registrationDate"`
	IsCommercial     int         `json:"isCommercial"`
	IsSeller         bool        `json:"isSeller"`
	Name             Name        `json:"name"`
	Home             HomeAddress `json:"address"`
	PhoneNumber      string      `json:"phone"`
	Email            string      `json:"email"`
	Vat              string      `json:"vat"`
	LegalInformation string      `json:"legalInformation"`
	RiskGroup        int         `json:"riskGroup"`
	LossPercentage   string      `json:"lossPercentage"`
	UnsentShipments  int         `json:"unsentShipments"`
	Reputation       int         `json:"reputation"`
	ShipsFast        int         `json:"shipsFast"`
	SellCount        int         `json:"sellCount"`
	SoldItems        int         `json:"soldItems"`
	AvgShippingTime  int         `json:"avgShippingTime"`
	OnVacation       bool        `json:"onVacation"`
}

type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type HomeAddress struct {
	Name    string      `json:"name"`
	Extra   interface{} `json:"extra"`
	Street  string      `json:"street"`
	Zip     string      `json:"zip"`
	City    string      `json:"city"`
	Country string      `json:"country"`
}

type Links []struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}

type MoneyDetails struct {
	TotalBalance           int `json:"totalBalance"`
	MoneyBalance           int `json:"moneyBalance"`
	BonusBalance           int `json:"bonusBalance"`
	UnpaidAmount           int `json:"unpaidAmount"`
	ProviderRechargeAmount int `json:"providerRechargeAmount"`
}

type BankAccount struct {
	XMLName      xml.Name `xml:"request"`
	AccountOwner string   `json:"accountOwner" xml:"bankAccountOwner"`
	Iban         string   `json:"iban" xml:"iban"`
	Bic          string   `json:"bic" xml:"bic"`
	BankName     string   `json:"bankName" xml:"bankName"`
	PhoneNumber  string   `json:"phoneNumber" xml:"phoneNumber"`
}
