package metamask

import (
	web3login "github.com/CypherBabel/web3-login"
	"github.com/CypherBabel/web3-login/util"
	"github.com/CypherBabel/web3-login/util/strutil"
	"github.com/ethereum/go-ethereum/crypto"
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

func (s *Service) Verify(responseBytes []byte) error {
	if responseBytes[util.SignatureSize-1] >= util.SignatureRIRangeBase {
		responseBytes[util.SignatureSize-1] -= util.SignatureRIRangeBase
	}

	pubkey, err := crypto.SigToPub(
		challenge(s.ChallengeStr).signatureHashBytes(),
		responseBytes,
	)
	if err != nil {
		return err
	}

	address := util.Address(crypto.PubkeyToAddress(*pubkey))
	if address.Hex() != s.Address {
		return util.ErrInvalidSignature
	}
	return nil
}
