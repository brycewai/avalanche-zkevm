// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package zkverify

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/vmerrs"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	VerifyGasCost      uint64 = 1000000000 /* SET A GAS COST HERE */
	VerifyInputGasCost uint64 = 1000000000 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// ZkVerifyRawABI contains the raw ABI of ZkVerify contract.
	//go:embed contract.abi
	ZkVerifyRawABI string

	ZkVerifyABI = contract.ParseABI(ZkVerifyRawABI)

	ZkVerifyPrecompile = createZkVerifyPrecompile()
)

type VerifyInput struct {
	Proof             []*big.Int
	Inputs            []*big.Int
	VerificationKeyId [32]byte
}

type VerifyInputInput struct {
	Inputs []*big.Int
	P      *big.Int
}

// UnpackVerifyInput attempts to unpack [input] as VerifyInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyInput(input []byte) (VerifyInput, error) {
	inputStruct := VerifyInput{}
	err := ZkVerifyABI.UnpackInputIntoInterface(&inputStruct, "verify", input)

	return inputStruct, err
}

// PackVerify packs [inputStruct] of type VerifyInput into the appropriate arguments for verify.
func PackVerify(inputStruct VerifyInput) ([]byte, error) {
	return ZkVerifyABI.Pack("verify", inputStruct.Proof, inputStruct.Inputs, inputStruct.VerificationKeyId)
}

// PackVerifyOutput attempts to pack given result of type bool
// to conform the ABI outputs.
func PackVerifyOutput(result bool) ([]byte, error) {
	return ZkVerifyABI.PackOutput("verify", result)
}

func verify(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, VerifyGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	// attempts to unpack [input] into the arguments to the VerifyInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackVerifyInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	_ = inputStruct // CUSTOM CODE OPERATES ON INPUT

	var output bool // CUSTOM CODE FOR AN OUTPUT
	packedOutput, err := PackVerifyOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackVerifyInputInput attempts to unpack [input] as VerifyInputInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyInputInput(input []byte) (VerifyInputInput, error) {
	inputStruct := VerifyInputInput{}
	err := ZkVerifyABI.UnpackInputIntoInterface(&inputStruct, "verify_input", input)

	return inputStruct, err
}

// PackVerifyInput packs [inputStruct] of type VerifyInputInput into the appropriate arguments for verify_input.
func PackVerifyInput(inputStruct VerifyInputInput) ([]byte, error) {
	return ZkVerifyABI.Pack("verify_input", inputStruct.Inputs, inputStruct.P)
}

// PackVerifyInputOutput attempts to pack given result of type bool
// to conform the ABI outputs.
func PackVerifyInputOutput(result bool) ([]byte, error) {
	return ZkVerifyABI.PackOutput("verify_input", result)
}

func verifyInput(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, VerifyInputGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	
	// attempts to unpack [input] into the arguments to the VerifyInputInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackVerifyInputInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	_ = inputStruct // CUSTOM CODE OPERATES ON INPUT

	var output bool // CUSTOM CODE FOR AN OUTPUT
	packedOutput, err := PackVerifyInputOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createZkVerifyPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createZkVerifyPrecompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"verify":       verify,
		"verify_input": verifyInput,
	}

	for name, function := range abiFunctionMap {
		method, ok := ZkVerifyABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}
