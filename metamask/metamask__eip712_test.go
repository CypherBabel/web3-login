package metamask

import "testing"

func TestSignEip712(t *testing.T) {
	client := &ServiceEip712{
		Address: "0xaC39b311DCEb2A4b2f5d8461c1cdaF756F4F7Ae9",
		ChallengeStr: `{
			"types": {
				"EIP712Domain": [
					{
						"name": "name",
						"type": "string"
					},
					{
						"name": "chainId",
						"type": "uint256"
					}
				],
				"RandomAmbireTypeStruct": [
					{
						"name": "identity",
						"type": "address"
					},
					{
						"name": "rewards",
						"type": "uint256"
					}
				]
			},
			"domain": {
				"name": "Ambire Typed test message",
				"chainId": "1"
			},
			"primaryType": "RandomAmbireTypeStruct",
			"message": {
				"identity": "0x0000000000000000000000000000000000000000",
				"rewards": 0
			}
		}`,
	}
	valid, err := client.Verify("0xee0d9f9e63fa7183bea2ca2e614cf539464a4c120c8dfc1d5ccc367f242a2c5939d7f59ec2ab413b8a9047de5de2f1e5e97da4eba2ef0d6a89136464f992dae11c")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result: ", valid)
}
