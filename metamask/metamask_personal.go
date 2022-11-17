package metamask

import (
	web3login "github.com/CypherBabel/web3-login"
	"github.com/CypherBabel/web3-login/util/strutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/storyicon/sigverify"
)

const (
	challengeStringLengthPersonal = 32
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

func (s *ServicePersonal) Challenge() string {
	s.ChallengeStr = strutil.Rand(challengeStringLengthPersonal)
	return s.ChallengeStr
}

func (s *ServicePersonal) Verify(verifyStr string) (bool, error) {
	return sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(s.Address),
		[]byte(s.ChallengeStr),
		verifyStr,
	)
}
