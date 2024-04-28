// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
import "@openzeppelin/contracts/access/Ownable.sol";
import "./ICrunchVendor.sol";
import "./ICrunchProtocol.sol";

// @creator yanghao@ohdat.io
contract CrunchVendor is ICrunchVendor, Ownable {
    address _spaceAddress;
    uint256 _tokenID;
    uint256 _price;
    mapping(address => address) _inviter;
    uint256 _totalSales;
    mapping(address => uint256) _userSales;

    constructor(
        address initialOwner,
        address spaceAddress,
        uint256 tokenID
    ) Ownable(initialOwner) {
        _spaceAddress = spaceAddress;
        _tokenID = tokenID;
    }

    function contentID() public view returns (uint256) {
        return _tokenID;
    }

    function creator() public view returns (address) {
        return owner();
    }

    function setPrice(uint256 price_) public onlyOwner {
        _price = price_;
    }

    function getPrice() public view returns (uint256) {
        return _price;
    }

    function inviter(address user) public view returns (address) {
        return _inviter[user];
    }

    function topUp(address inviter_, uint256 amount) public payable {
        require(msg.value == _price * amount, "price not match");
        if (
            inviter_ != address(0) &&
            inviter_ != msg.sender &&
            _inviter[msg.sender] == address(0)
        ) {
            _inviter[msg.sender] = inviter_;
            emit Invite(inviter_, msg.sender);
        }
        // 三层分润
        uint256[] memory rates = ICrunchProtocol(_spaceAddress)
            .getCommissionRate();
        uint256 rechargeValue_ = msg.value;
        for (uint256 i = 0; i < rates.length; i++) {
            if (i > 0) {
                inviter_ = _inviter[inviter_];
            }
            if (inviter_ == address(0)) {
                break;
            }
            uint256 commission = (amount * rates[i]) / 100;
            payable(inviter_).transfer(commission);
            emit Commission(inviter_, msg.sender, commission, i + 1);
            rechargeValue_ -= commission;
        }
        _userSales[msg.sender] += amount;
        _totalSales += amount;
        ICrunchProtocol(_spaceAddress).revenuePool{value: rechargeValue_}(
            _tokenID
        );
        emit TopUpSuccess(msg.sender, amount);
    }

    function userBalance(address user_) public view returns (uint256) {
        return _userSales[user_];
    }

    function boxOffice() public view returns (uint256) {
        return _totalSales;
    }
}
