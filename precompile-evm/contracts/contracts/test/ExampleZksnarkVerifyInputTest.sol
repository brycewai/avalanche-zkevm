//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../ExampleZksnarkVerifyInput.sol";
import "../interfaces/IZksnarkVerifyInput.sol";
import "@avalabs/subnet-evm-contracts/contracts/test/AllowListTest.sol";

contract ExampleZksnarkVerifyInputTest is AllowListTest {
  IZksnarkVerifyInput zksnarkVerifyInput = IZksnarkVerifyInput(ZksnarkVerifyInput_ADDRESS);

  //   function step_getDefaultzksnarkVerifyInput() public {
  //     ExampleZksnarkVerifyInput example = new ExampleZksnarkVerifyInput();
  //     address exampleAddress = address(example);

  //     assertRole(zksnarkVerifyInput.readAllowList(exampleAddress), AllowList.Role.None);
  //     assertEq(example.verifyPublicInput([10, 20]), "Hello World!");
  //   }

  //   function step_doesNotSetGreetingBeforeEnabled() public {
  //     ExampleZksnarkVerifyInput example = new ExampleZksnarkVerifyInput();
  //     address exampleAddress = address(example);

  //     assertRole(zksnarkVerifyInput.readAllowList(exampleAddress), AllowList.Role.None);

  //     try example.setGreeting("testing") {
  //       assertTrue(false, "setGreeting should fail");
  //     } catch {} // TODO should match on an error to make sure that this is failing in the way that's expected
  //   }

  //   function step_setAndGetGreeting() public {
  //     ExampleZksnarkVerifyInput example = new ExampleZksnarkVerifyInput();
  //     address exampleAddress = address(example);

  //     assertRole(zksnarkVerifyInput.readAllowList(exampleAddress), AllowList.Role.None);
  //     zksnarkVerifyInput.setEnabled(exampleAddress);
  //     assertRole(zksnarkVerifyInput.readAllowList(exampleAddress), AllowList.Role.Enabled);

  //     string memory greeting = "testgreeting";
  //     example.setGreeting(greeting);
  //     assertEq(example.sayHello(), greeting);
  //   }
}
