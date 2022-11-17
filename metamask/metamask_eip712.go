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

type ServiceEip712 struct {
	Address      string // 钱包地址
	InviterId    int64  // 邀请人id
	ChallengeStr string
}

func NewServiceEip712(address string, inviterId int64) web3login.Web3login {
	return &ServiceEip712{
		Address:   address,
		InviterId: inviterId,
	}
}
func (s *ServiceEip712) GetInviterId() int64 {
	return s.InviterId
}

func (s *ServiceEip712) Challenge() string {
	s.ChallengeStr = strutil.Rand(challengeStringLength)
	return s.ChallengeStr
}

func (s *ServiceEip712) Verify(verifyStr string) (bool, error) {
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
