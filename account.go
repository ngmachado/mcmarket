package mcmarket

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/ngmachado/mcmarket/Entities"
	"github.com/ngmachado/mkm"
)

type Account struct {
	cl *mkm.Client
}

//Returns the Account entity of the authenticated user
func (r *Account) Information() (*Entities.Account, error) {
	body, err := r.cl.Request(mkm.Get, "/account", nil)
	if err != nil {
		return nil, err
	}
	account := &Entities.Account{}
	err = json.Unmarshal(body, account)
	return account, err
}

//Updates the vacation status of the authenticated user; returns the Account entity of the authenticated user
func (r *Account) Vacation(vacation bool) (*Entities.Account, error) {
	resource := fmt.Sprintf("/account/vacation?onVacation=%t", vacation)
	body, err := r.cl.Request(mkm.Put, resource, nil)
	if err != nil {
		return nil, err
	}
	account := &Entities.Account{}
	err = json.Unmarshal(body, account)
	return account, err
}

//Updates the display language of the authenticated user; returns the Account entity of the authenticated user
func (r *Account) Language(lang int) (*Entities.Account, error) {
	resource := fmt.Sprintf("/account/language?idDisplayLanguage=%d", lang)
	body, err := r.cl.Request(mkm.Put, resource, nil)
	if err != nil {
		return nil, err
	}
	account := &Entities.Account{}
	err = json.Unmarshal(body, account)
	return account, err
}

//Redeems one or more coupons
func (r *Account) Redeem(couponCode int) (*Entities.Account, error) {
	postData := []byte(fmt.Sprintf("%s<request><couponCode>%d</couponCode></request>", xml.Header, couponCode))
	body, err := r.cl.Request(mkm.Post, "/account/coupon", postData)
	if err != nil {
		return nil, err
	}
	account := &Entities.Account{}
	err = json.Unmarshal(body, account)
	return account, err
}

func (r *Account) SellerActivation(bank *Entities.BankAccount) ([]byte, error) {
	postData, err := xml.Marshal(bank)
	if err != nil {
		return nil, err
	}
	post := append([]byte(xml.Header), postData...)
	return r.cl.Request(mkm.Post, "/account/selleractivation", post)
}

func (r *Account) SellerCompeleteActivation(firstAmount float32, secondAmount float32) ([]byte, error) {
	postData := []byte(fmt.Sprintf("%s<request><amount1>%.2f</amount1><amount2>%.2f</amount2></request>", xml.Header, firstAmount, secondAmount))
	return r.cl.Request(mkm.Put, "/account/selleractivation", postData)
}
