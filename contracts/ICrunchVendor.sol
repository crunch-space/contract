// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICrunchVendor {
    //event
    // Invitation success notification
    event Invite(address indexed inviter, address indexed invitee);
    // Distribution notice
    event Commission(
        address indexed inviter,
        address indexed invitee,
        uint256 amount,
        uint256 level
    );
    // User recharge success notification
    event TopUpSuccess(address indexed user, uint256 amount);

    // read function
    // Get the tokenID of the dapp
    function contentID() external view returns (uint256);

    // Get creator address
    function creator() external view returns (address);

    // total sales quantity
    function boxOffice() external view returns (uint256);

    //Get the number of user purchases
    function userBalance(address) external view returns (uint256);

    //Get inviter
    function inviter(address) external view returns (address);

    // write function
    // User recharge (purchase)
    function topUp(address invater_, uint256 amount) external payable;
}
