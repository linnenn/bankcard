package backcard

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BankCardInfoStruct struct {
	Validated       bool   `json:"validated"`
	ValidatedMsg    string `json:"validated_msg"`
	BankENShortName string `json:"bank_en_short_name"`
	BankENFullName  string `json:"bank_en_full_name"`
	BankCNShortName string `json:"bank_cn_short_name"`
	BankCNFullName  string `json:"bank_cn_full_name"`
	BankImg         string `json:"bank_img"`
	CardType        string `json:"card_type"`
	CardTypeName    string `json:"card_type_name"`
}

func BankCardInfo(cardNO string) (error, *BankCardInfoStruct) {
	form := url.Values{}
	form.Add("_input_charset", "utf-8")
	form.Add("cardNo", cardNO)
	form.Add("cardBinCheck", "true")
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(getBankInfo, form.Encode()), nil)
	if err != nil {
		return err, nil
	}
	resp, err := commonHttpClient5s.Do(req)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	type AliPayBankInfo struct {
		Messages []struct {
			ErrorCodes string `json:"errorCodes"`
			Name       string `json:"name"`
		} `json:"messages"`
		CardType  string `json:"cardType"`
		Bank      string `json:"bank"`
		Key       string `json:"key"`
		Validated bool   `json:"validated"`
		Stat      string `json:"stat"`
	}
	result := new(AliPayBankInfo)
	if err := json.Unmarshal(body, result); err != nil {
		return err, nil
	}
	if !result.Validated {
		err := errors.New(result.Messages[0].ErrorCodes)
		return err, &BankCardInfoStruct{
			Validated:    false,
			ValidatedMsg: "cardNO error",
		}
	}
	bankName := ""
	if bank, ok := bankShortNameMap[result.Bank]; ok {
		bankName = bank
	}
	return nil, &BankCardInfoStruct{
		Validated:       true,
		BankENShortName: result.Bank,
		BankENFullName:  result.Bank,
		BankCNShortName: bankName,
		BankCNFullName:  bankName,
		BankImg:         fmt.Sprintf(bankImageUrl, result.Bank),
		CardType:        result.CardType,
		CardTypeName:    bankCardTypeMap[result.CardType],
	}
}
