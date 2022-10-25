package metamask

import (
	"encoding/json"

	web3login "github.com/CypherBabel/web3-login"
	"github.com/CypherBabel/web3-login/util/strutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

const (
	challengeStringLength = 32
)

type Service struct {
	Address      string // 钱包地址
	ChallengeStr string
}

func NewService(address string) web3login.Web3login {
	return &Service{
		Address: address,
	}
}

func (s *Service) Challenge() string {
	s.ChallengeStr = strutil.Rand(challengeStringLength)
	return s.ChallengeStr
}

func (s *Service) Verify(verifyStr string) (bool, error) {
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(s.ChallengeStr), &typedData); err != nil {
		return false, err
	}
	return sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress(s.Address),
		typedData,
		verifyStr,
	)
}
