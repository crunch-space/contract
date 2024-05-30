// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICrunchProtocol {
    //event
    //Deployment contract notice
    event DeployCrunchVendor(
        address indexed creator,
        address indexed crunchApp,
        uint256 tokenID
    );

    //mint
    function mint(uint256 tokenID, uint256 amount) external payable;

    // claim && burn
    function burnToClaim(uint256 tokenID, uint256 amount) external payable;

    // deploy crunch Vendor
    function deployCrunchVendor(
        uint256 tokenID,
        uint256 price, //NFT price
        uint256 creatorCommissionRate_, // Creator share ratio 1-100
        uint256 amount, // NFT quantity
        bytes memory signature_
    ) external;

    // Set the moisture ratio for each layer
    // example: [10, 5, 3] 10%, 5%, 3%
    function setCommissionRate(uint256[] memory rates) external;

    // Get the proportion of share
    function getCommissionRate() external view returns (uint256[] memory);

    //recharge
    // crunch app Contract profits are channeled into this contract
    // Proceeds from other chains also feed into this contract
    function revenuePool(uint256 tokenID) external payable;

    // withdraw
    // The creator reflects the investment amount
    function withdraw(
        uint256 tokenID,
        uint256 amount,
        bytes memory signature,
        string memory nonce
    ) external payable;

    // read function

    //Get the contract address of tokenid
    function contentContract(uint256 tokenID) external view returns (address);

    // tokenID
    // The current number of mint
    function tokenMinted(uint256 tokenID) external view returns (uint256);

    // Get the price at which the token can be exchanged
    function tokenValue(uint256 tokenID) external view returns (uint256);
}
