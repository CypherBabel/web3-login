package metamask

import "testing"

func TestSignPersonla(t *testing.T) {
	client := &ServicePersonal{
		Address:      "0xd27b80cf103a33d6f7af46565eff7ab8374e3f72",
		ChallengeStr: "asdasx",
		// ChallengeStr: `Welcome to OpenSea!\nClick to sign in and accept the OpenSea Terms of Service: https://opensea.io/tos\nThis request will not trigger a blockchain transaction or cost any gas fees.\nYour authentication status will reset after 24 hours\n\nWallet address:\n'0xd27b80cf103a33d6f7af46565eff7ab8374e3f72'\n\nNonce:\n'aasdas4654asd65hdasx'.`,
	}
	valid, err := client.Verify("0x194fd71c0451b64f2cdf02cf98c27c95ef44fb12183d96a4e26be6f109bbe6806fe3b9fe67db0a2b35511f6237d7d38898cb9b9cc1128cf097a0c6cba531c8bd1c")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result: ", valid)
}
