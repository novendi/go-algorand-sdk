package templates

import (
	"encoding/base64"
	"github.com/algorand/go-algorand-sdk/types"
)

type Split struct {
	address string
	program string
}

const referenceProgram = "ASAIAQUCBgcICQomAyCztwQn0+DycN+vsk+vJWcsoz/b7NDS6i33HOkvTpf+YiC3qUpIgHGWE8/1LPh9SGCalSN7IaITeeWSXbfsS5wsXyC4kBQ38Z8zcwWVAym4S8vpFB/c0XC6R4mnPi9EBADsPDEQIhIxASMMEDIEJBJAABoxCSgSMQcyAxIxCCUSMQIhBA0QEBAiQAAzSDMAADMBABIxCTIDEjMABykSMwEHKhIzAAgzAAgzAQgIIQULIQYKEjMACCEHDRAQEBAQEA=="

var referenceOffsets = []uint64{1, 2, 3, 4, 5, 6, 7} // TODO values

// GetAddress returns the contract address
func (contract Split) GetAddress() string {
	return contract.address
}

// GetProgram returns b64-encoded version of the program
func (contract Split) GetProgram() string {
	return contract.program
}

//GetSendFundsTransaction returns a group transactions array which transfer funds according to the contract's ratio
// amount: uint64 number of assets to be transferred
// precise: handles rounding error. When False, the amount will be divided as closely as possible but one account will get
// 			slightly more. When true, returns an error.
func (contract Split) GetSendFundsTransaction(amount uint64, precise bool) ([]types.Transaction, error) {
	return nil, nil
}

func MakeSplit(owner, receiverOne, receiverTwo string, ratn, ratd, expiryRound, minPay, maxFee uint64) (Split, error) {
	referenceAsBytes, err := base64.StdEncoding.DecodeString(referenceProgram)
	if err != nil {
		return Split{}, err
	}
	injectionVector := []interface{}{owner, receiverOne, receiverTwo, ratn, ratd, expiryRound, minPay, maxFee} // TODO ordering
	injectedBytes, err := inject(referenceAsBytes, referenceOffsets, injectionVector)
	if err != nil {
		return Split{}, err
	}
	injectedProgram := base64.StdEncoding.EncodeToString(injectedBytes)
	return Split{address: "", program: injectedProgram}, err
}