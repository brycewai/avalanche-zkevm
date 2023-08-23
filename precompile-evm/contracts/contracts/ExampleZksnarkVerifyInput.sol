//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./interfaces/IZksnarkVerifyInput.sol";

address constant ZksnarkVerifyInput_ADDRESS = 0x0300000000000000000000000000000000000077;

// ExampleHelloWorld shows how the HelloWorld precompile can be used in a smart contract.
contract ExampleHelloWorld {
  IZksnarkVerifyInput zksnarkVerifyInput = IZksnarkVerifyInput(ZksnarkVerifyInput_ADDRESS);

  function verifyPublicInput(uint256[] memory input) external returns (bool result){
    returns zksnarkVerifyInput.verifyPublicInput(input);
  }
}