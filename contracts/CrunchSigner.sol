// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract CrunchSigner {
    error SignatureLengthNotEnough65();
    error InvalidSignature();

    address internal _signer;

    constructor(address signer_) {
        _setSigner(signer_);
    }

    function getHash(
        bytes memory data
    ) internal pure virtual returns (bytes32) {
        bytes32 hash = keccak256(
            abi.encodePacked(
                "\x19Ethereum Signed Message:\n32",
                keccak256(abi.encodePacked(data))
            )
        );
        return hash;
    }

    function _setSigner(address signer_) internal {
        _signer = signer_;
    }

    function _getSigner() internal view returns (address) {
        return _signer;
    }

    function checkSign(
        bytes memory signature,
        bytes32 hash
    ) internal view virtual returns (bool) {
        if (signature.length != 65) {
            revert SignatureLengthNotEnough65();
        }
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := byte(0, mload(add(signature, 96)))
        }

        return ecrecover(hash, v, r, s) == _signer;
    }
}
