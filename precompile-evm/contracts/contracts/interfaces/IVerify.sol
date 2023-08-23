// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

interface IVerify {
  function verifyPublicInput(uint256[] calldata inputs, uint256 p) external returns (bool result);

  function verifyProof(
    uint256[] calldata proof,
    uint256[] calldata inputs,
    bytes32 verificationKeyId
  ) external returns (bool result);
}
