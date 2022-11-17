package metamask

import (
	web3login "github.com/CypherBabel/web3-login"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/storyicon/sigverify"
)

type ServicePersonal struct {
	Address      string // 钱包地址
	InviterId    int64  // 邀请人id
	ChallengeStr string
}

func NewServicePersonal(address string, inviterId int64) web3login.Web3login {
	return &ServicePersonal{
		Address:   address,
		InviterId: inviterId,
	}
}
func (s *ServicePersonal) GetInviterId() int64 {
	return s.InviterId
}

func (s *ServicePersonal) Challenge(str string) {
	s.ChallengeStr = str
}

func (s *ServicePersonal) Verify(verifyStr string) (bool, error) {
	return sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(s.Address),
		[]byte(s.ChallengeStr),
		verifyStr,
	)
}
