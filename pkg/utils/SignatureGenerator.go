package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"projectsphere/cats-social/pkg/utils/config"
)

func PaymentSignatureGenerator(CardNumber string, amount int) string {
	data := fmt.Sprintf("%v:%d:%v", CardNumber, amount, config.Get().SeaLabsPayAPI.MerchantCode)
	h := hmac.New(sha256.New, []byte(config.Get().SeaLabsPayAPI.APIKey))

	h.Write([]byte(data))

	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func CallbackSealabsPaySignatureGenerator(TxnId int, Amount int, CardNumber string, Status string, Message string) string {
	data := fmt.Sprintf("%d:%d:%v:%v:%v", TxnId, Amount, config.Get().SeaLabsPayAPI.MerchantCode, Status, Message)
	h := hmac.New(sha256.New, []byte(config.Get().SeaLabsPayAPI.APIKey))

	h.Write([]byte(data))

	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
