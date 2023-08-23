// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package zksnarkverifyinput

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/allowlist"
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
	// This contract also uses AllowList precompile.
	// You should also increase gas costs of functions that read from AllowList storage.
	VerifyPublicInputGasCost uint64 = 1 /* SET A GAS COST HERE */ + allowlist.ReadAllowListGasCost
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
	ErrCannotVerifyPublicInput = errors.New("non-enabled cannot call verifyPublicInput")

	// ZksnarkVerifyInputRawABI contains the raw ABI of ZksnarkVerifyInput contract.
	//go:embed contract.abi
	ZksnarkVerifyInputRawABI string

	ZksnarkVerifyInputABI = contract.ParseABI(ZksnarkVerifyInputRawABI)

	ZksnarkVerifyInputPrecompile = createZksnarkVerifyInputPrecompile()
)

// GetZksnarkVerifyInputAllowListStatus returns the role of [address] for the ZksnarkVerifyInput list.
func GetZksnarkVerifyInputAllowListStatus(stateDB contract.StateDB, address common.Address) allowlist.Role {
	return allowlist.GetAllowListStatus(stateDB, ContractAddress, address)
}

// SetZksnarkVerifyInputAllowListStatus sets the permissions of [address] to [role] for the
// ZksnarkVerifyInput list. Assumes [role] has already been verified as valid.
// This stores the [role] in the contract storage with address [ContractAddress]
// and [address] hash. It means that any reusage of the [address] key for different value
// conflicts with the same slot [role] is stored.
// Precompile implementations must use a different key than [address] for their storage.
func SetZksnarkVerifyInputAllowListStatus(stateDB contract.StateDB, address common.Address, role allowlist.Role) {
	allowlist.SetAllowListRole(stateDB, ContractAddress, address, role)
}

// UnpackVerifyPublicInputInput attempts to unpack [input] into the []*big.Int type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackVerifyPublicInputInput(input []byte) ([]*big.Int, error) {
	res, err := ZksnarkVerifyInputABI.UnpackInput("verifyPublicInput", input)
	if err != nil {
		return nil, err
	}
	unpacked := *abi.ConvertType(res[0], new([]*big.Int)).(*[]*big.Int)
	return unpacked, nil
}

// PackVerifyPublicInput packs [input] of type []*big.Int into the appropriate arguments for verifyPublicInput.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackVerifyPublicInput(input []*big.Int) ([]byte, error) {
	return ZksnarkVerifyInputABI.Pack("verifyPublicInput", input)
}

// PackVerifyPublicInputOutput attempts to pack given result of type bool
// to conform the ABI outputs.
func PackVerifyPublicInputOutput(result bool) ([]byte, error) {
	return ZksnarkVerifyInputABI.PackOutput("verifyPublicInput", result)
}

func verifyPublicInput(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, VerifyPublicInputGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	// attempts to unpack [input] into the arguments to the VerifyPublicInputInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackVerifyPublicInputInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// Allow list is enabled and VerifyPublicInput is a state-changer function.
	// This part of the code restricts the function to be called only by enabled/admin addresses in the allow list.
	// You can modify/delete this code if you don't want this function to be restricted by the allow list.
	stateDB := accessibleState.GetStateDB()
	// Verify that the caller is in the allow list and therefore has the right to call this function.
	callerStatus := allowlist.GetAllowListStatus(stateDB, ContractAddress, caller)
	if !callerStatus.IsEnabled() {
		return nil, remainingGas, fmt.Errorf("%w: %s", ErrCannotVerifyPublicInput, caller)
	}
	// allow list code ends here.

	// CUSTOM CODE STARTS HERE
	var output bool = true
	P, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	for _, publicInput := range inputStruct{
		comp := P.Cmp(publicInput)  // p > publicInput -> 1 ;p = publicInput -> 0 ;p < publicInput -> -1 ;
		if comp <= 0{
			output = false
			break
		}
	}

	packedOutput, err := PackVerifyPublicInputOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createZksnarkVerifyInputPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.
// Access to the getters/setters is controlled by an allow list for ContractAddress.
func createZksnarkVerifyInputPrecompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction
	functions = append(functions, allowlist.CreateAllowListFunctions(ContractAddress)...)

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"verifyPublicInput": verifyPublicInput,
	}

	for name, function := range abiFunctionMap {
		method, ok := ZksnarkVerifyInputABI.Methods[name]
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