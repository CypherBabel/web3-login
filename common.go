package web3login

// Web3login
type Web3login interface {
	Challenge(string)
	Verify(string) (bool, error)
	GetInviterId() int64
}
