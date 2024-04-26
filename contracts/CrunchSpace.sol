// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol";
import "./ICrunchSpace.sol";
import "./CrunchApp.sol";
import "./CrunchSigner.sol";

// @creator yanghao@ohdat.io
contract CrunchSpace is
    ERC1155,
    Ownable,
    ERC1155Supply,
    ICrunchSpace,
    CrunchSigner
{
    constructor(
        address initialOwner
    ) ERC1155("") Ownable(initialOwner) CrunchSigner(initialOwner) {}

    //每个tokenID对应的总量
    mapping(uint256 => uint256) internal _totalSupply;
    uint256[] internal _layerCommissionRates; // app layer commission rates

    mapping(uint256 => TokenInfo) internal TokenMap;

    modifier checkBeforeDeploy(uint256 tokenID, bytes memory signature_) {
        //test network return
        _;
        return;
        bytes memory h = abi.encodePacked(tokenID, msg.sender, address(this));
        if (!checkSign(signature_, getHash(h))) {
            revert InvalidSignature();
        }
        _;
    }

    struct TokenInfo {
        uint256 price; // NFT 单价
        uint256 totalSupply; // NFT total supply
        uint256 balance; // 设计师可提现金额
        uint256 rechargeBalance; // 充值金额
        uint256 creatorCommissionRate; // 设计师分成比例
        address creator;
        address crunchApp;
    }

    function setURI(string memory newuri) public onlyOwner {
        _setURI(newuri);
    }

    //TODO: implement the following functions
    function mint(uint256 tokenID, uint256 amount) public payable override {
        // check tokenID exist
        require(TokenMap[tokenID].creator == address(0), "tokenID exist");
        _mint(msg.sender, tokenID, amount, "");
        uint256 totalPrice = amount * TokenMap[tokenID].price;
        if (msg.value > totalPrice) {
            payable(msg.sender).transfer(msg.value - totalPrice);
        }
        TokenMap[tokenID].balance += totalPrice;
    }

    function burnToClaim(
        uint256 tokenID,
        uint256 amount
    ) public payable override {
        require(balanceOf(msg.sender, tokenID) >= amount, "balance not enough");
        uint256 balance = tokenPrice(tokenID) * amount;
        payable(msg.sender).transfer(balance);
        TokenMap[tokenID].rechargeBalance =
            TokenMap[tokenID].rechargeBalance -
            balance;
        _burn(msg.sender, tokenID, amount);
    }

    function deployCrunchApp(
        uint256 tokenID,
        uint256 price,
        uint256 creatorCommissionRate_,
        uint256 amount,
        bytes memory signature_
    ) public override checkBeforeDeploy(tokenID, signature_) {
        // check tokenID exist
        require(TokenMap[tokenID].creator == address(0), "tokenID exist");
        require(price > 0, "price must > 0");
        require(creatorCommissionRate_ > 0, "creatorCommissionRate_ must > 0");
        require(
            creatorCommissionRate_ <= 100,
            "creatorCommissionRate_ must <= 100"
        );

        // deploy crunch app
        CrunchApp app = new CrunchApp(msg.sender, address(this), tokenID);
        TokenMap[tokenID] = TokenInfo(
            price,
            amount,
            0,
            0,
            creatorCommissionRate_,
            msg.sender,
            address(app)
        );
        // event
        emit DeployCrunchApp(msg.sender, address(app), tokenID);
    }

    function setCommissionRate(uint256[] memory rates) public override {
        _layerCommissionRates = rates;
    }

    function getCommissionRate() public view returns (uint256[] memory) {
        return _layerCommissionRates;
    }

    function setSigner(address signer_) public onlyOwner {
        _setSigner(signer_);
    }

    function withdraw(uint256 tokenID, uint256 amount) public payable override {
        //check token id
        require(TokenMap[tokenID].creator != address(0), "tokenID not exist");
        // check creator
        require(TokenMap[tokenID].creator == msg.sender, "only creator");
        // check balance
        require(TokenMap[tokenID].balance >= amount, "balance not enough");
        TokenMap[tokenID].balance = TokenMap[tokenID].balance - amount;
        payable(msg.sender).transfer(amount);
    }

    function recharge(uint256 tokenID) public payable override {
        //
        uint256 totalPrice = TokenMap[tokenID].price *
            totalSupplyOfTokenID(tokenID);
        uint256 rechargeBalance = TokenMap[tokenID].rechargeBalance;
        if (rechargeBalance >= totalPrice) {
            // creator 分成金额
            uint256 creatorAmount = (msg.value *
                TokenMap[tokenID].creatorCommissionRate) / 100;
            payable(TokenMap[tokenID].creator).transfer(creatorAmount);
            TokenMap[tokenID].rechargeBalance += msg.value - creatorAmount;
        } else if (rechargeBalance + msg.value > totalPrice) {
            // 分成金额
            uint256 balance = rechargeBalance + msg.value - totalPrice;
            // creator 分成金额
            uint256 creatorAmount = (balance *
                TokenMap[tokenID].creatorCommissionRate) / 100;
            payable(TokenMap[tokenID].creator).transfer(creatorAmount);
            TokenMap[tokenID].rechargeBalance += msg.value - creatorAmount;
        } else {
            TokenMap[tokenID].rechargeBalance += msg.value;
        }
    }

    function totalSupplyOfTokenID(
        uint256 tokenID
    ) public view override returns (uint256) {
        return _totalSupply[tokenID];
    }

    function tokenPrice(
        uint256 tokenID
    ) public view override returns (uint256) {
        return
            TokenMap[tokenID].rechargeBalance / totalSupplyOfTokenID(tokenID);
    }

    function dappContract(
        uint256 tokenID
    ) public view override returns (address) {
        return TokenMap[tokenID].crunchApp;
    }

    // The following functions are overrides required by Solidity.
    function _update(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory values
    ) internal override(ERC1155, ERC1155Supply) {
        super._update(from, to, ids, values);
        for (uint256 i = 0; i < ids.length; i++) {
            if (from == address(0)) {
                _totalSupply[ids[i]] += values[i];
            }
            if (to == address(0)) {
                _totalSupply[ids[i]] -= values[i];
            }
        }
    }
}
