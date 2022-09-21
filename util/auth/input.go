package auth

import (
	"github.com/CypherBabel/web3-login/util"
)

type AuthInput struct {
	Address   string `json:"address" form:"address"`
	Signature string `json:"signature" form:"signature"`
}

func NewAuthInput(addressHex, sigHex string) *AuthInput {
	return &AuthInput{
		Address:   addressHex,
		Signature: sigHex,
	}
}

func (in *AuthInput) Validate() error {
	if err := util.ValidateAddressHex(in.Address); err != nil {
		return err
	}
	if err := util.ValidateSignatureHex(in.Signature); err != nil {
		return err
	}
	return nil
}

func (in *AuthInput) Sig() util.Signature {
	return util.NewSignatureFromHex(in.Signature)
}
