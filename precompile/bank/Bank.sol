// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant BANK_PRECOMPILE_ADDRESS = 0x1000000000000000000000000000000000000001;

IBank constant BANK_CONTRACT = IBank(
    BANK_PRECOMPILE_ADDRESS
);

interface IBank {
    // Transactions
    function send(
        address fromAddress,
        address toAddress,
        uint256 amount,
        string memory denom
    ) external returns (bool success);

    // Queries
    function balance(
        address addr,
        string memory denom
    ) external view returns (uint256 balance);

    function supplyOf(
        string memory denom
    ) external view returns (uint256 supply);
}
