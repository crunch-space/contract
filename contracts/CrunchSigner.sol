// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// @creator yanghao@ohdat.io
contract CrunchSigner {
    error SignatureLengthNotEnough65();
    address internal _signer;

    constructor(address signer_) {
        _setSigner(signer_);
    }

    function _setSigner(address signer_) internal {
        _signer = signer_;
    }

    function _getSigner() internal view returns (address) {
        return _signer;
    }
}
