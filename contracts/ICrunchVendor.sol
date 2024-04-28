// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICrunchVendor {
    //event
    // 邀请成功通知
    event Invite(address indexed inviter, address indexed invitee);
    // 分润通知
    event Commission(
        address indexed inviter,
        address indexed invitee,
        uint256 amount,
        uint256 level
    );
    // 用户充值成功通知
    event TopUpSuccess(address indexed user, uint256 amount);

    // read function
    // 获取 dapp 的 tokenID
    function contentID() external view returns (uint256);

    // 获取创建者地址
    function creator() external view returns (address);

    // 总销售数量
    function boxOffice() external view returns (uint256);

    //获取用户购买的数量
    function userBalance(address) external view returns (uint256);

    //查询邀请人
    function inviter(address) external view returns (address);

    // write function
    // 用户充值(购买)
    function topUp(address invater_, uint256 amount) external payable;
}
