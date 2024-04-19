// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICrunchApp {
    //event
    event Invite(address indexed invater, address indexed invitee);
    event Commission(
        address indexed invater,
        address indexed invitee,
        uint256 amount,
        uint256 level
    );

    // read function
    function dappID() external view returns (uint256);

    function creator() external view returns (address);

    function sales() external view returns (uint256);

    function userSales(address) external view returns (uint256);

    function invater(address) external view returns (address);

    // write function
    function recharge(address invater, uint256 amount) external payable;
}
