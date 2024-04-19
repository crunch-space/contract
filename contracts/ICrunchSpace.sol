// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICrunchSpace {
    //event
    event Invite(address indexed invater, address indexed invitee);
    event Commission(
        address indexed invater,
        address indexed invitee,
        uint256 amount,
        uint256 level
    );
    event DeployCrunchApp(
        address indexed creator,
        address indexed crunchApp,
        uint256 tokenID
    );

    //mint crowd funding
    function mint(uint256 tokenID, uint256 amount) external payable;

    // claim && burn
    function claim(uint256 tokenID, uint256 amount) external payable;

    // deploy crunch app
    function deployCrunchApp(
        uint256 tokenID,
        uint256 price,
        uint256 amount
    ) external;

    // 设置每层分润比例
    function setCommissionRate(uint256[] memory rates) external;

    function getCommissionRate() external view returns (uint256[] memory);

    //
    function recharge(uint256 tokenID) external payable;

    // withdraw
    function withdraw(uint256 tokenID, uint256 amount) external payable;

    function totalSupplyOfTokenID(
        uint256 tokenID
    ) external view returns (uint256);

    // read function
    //获取dappContract
    function dappContract(uint256 tokenID) external view returns (address);

    // 获取投资金额
    // function sales(uint256 tokenID) external view returns (uint256);

    // // 获取用户投资金额
    // function userSales(
    //     uint256 tokenID,
    //     address user
    // ) external view returns (uint256);
}
