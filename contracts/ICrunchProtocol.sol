// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICrunchProtocol {
    //event
    //部署合约通知
    event DeployCrunchVendor(
        address indexed creator,
        address indexed crunchApp,
        uint256 tokenID
    );

    //mint (众筹)
    function mint(uint256 tokenID, uint256 amount) external payable;

    // claim && burn
    function burnToClaim(uint256 tokenID, uint256 amount) external payable;

    // deploy crunch Vendor
    function deployCrunchVendor(
        uint256 tokenID,
        uint256 price, //NFT 价格
        uint256 creatorCommissionRate_, // 创建者分润比例 1-100
        uint256 amount, // NFT 总量
        bytes memory signature_
    ) external;

    // 设置每层分润比例
    // example: [10, 5, 3] //三层分润 10%, 5%, 3%
    function setCommissionRate(uint256[] memory rates) external;

    // 获取分润比例
    function getCommissionRate() external view returns (uint256[] memory);

    //充值
    // crunch app 合约利润打入到这个合约
    // 从其他的链的收益也打入到这个合约
    function revenuePool(uint256 tokenID) external payable;

    // withdraw
    // 创作者 体现投资金额
    function withdraw(
        uint256 tokenID,
        uint256 amount,
        bytes memory signature,
        string memory nonce
    ) external payable;

    // read function

    //查询 tokenid 的 合约地址
    function contentContract(uint256 tokenID) external view returns (address);

    // tokenID 目前已经mint的数量
    function tokenMinted(uint256 tokenID) external view returns (uint256);

    // 查询 token 可以兑换的价格
    function tokenValue(uint256 tokenID) external view returns (uint256);
}
