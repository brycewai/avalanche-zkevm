// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;
interface IVerify {
    
    function verify_input(uint256[] calldata inputs,uint256 p) external returns (bool result);
    
    /// @notice Checks the arguments of Proof, through elliptic curve
    ///  pairing functions.
    /// @dev
    ///  MUST return `true` if Proof passes all checks (i.e. the Proof is
    ///  valid).
    ///  MUST return `false` if the Proof does not pass all checks (i.e. if the
    ///  Proof is invalid).
    /// @param proof A zk-SNARK.
    /// @param inputs Public inputs which accompany Proof.
    /// @param verificationKeyId A unique identifier (known to this verifier
    ///  contract) for the Verification Key to which Proof corresponds.
    /// @return result The result of the verification calculation. True
    ///  if Proof is valid; false otherwise.
    function verify(uint256[] calldata proof, uint256[] calldata inputs, bytes32 verificationKeyId) external returns (bool result); 

}