import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import * as dotenv from "dotenv";
dotenv.config();

const config: HardhatUserConfig = {
  solidity: "0.8.24",
  networks: {
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.INFURA_API_KEY}`,
      accounts: [String(process.env.PRIVATE_KEY)],
    },
    mainnet: {
      url: `https://mainnet.infura.io/v3/${process.env.INFURA_API_KEY}`,
      accounts: [String(process.env.PRIVATE_KEY)],
    },
    // for mainnet
    "base-mainnet": {
      url: "https://mainnet.base.org",
      accounts: [String(process.env.PRIVATE_KEY)],
    },
    // for testnet
    "base-sepolia": {
      url: "https://sepolia.base.org",
      accounts: [String(process.env.PRIVATE_KEY)],
    },
    // for Sepolia testnet
    blast_sepolia: {
      url: "https://sepolia.blast.io",
      accounts: [String(process.env.PRIVATE_KEY)],
    },
  },
  etherscan: {
    // Your API key for Etherscan
    // Obtain one at https://etherscan.io/
    apiKey: {
      mainnet: `${process.env.ETHERSCAN_API_KEY}`, // process.env.ETHERSCAN_API_KEY,
      goerli: `${process.env.ETHERSCAN_API_KEY}`,
      blast_sepolia: "blast_sepolia", // apiKey is not required, just set a placeholder
      "base-sepolia": `${process.env.BASE_ETHERSCAN_API_KEY}`,
      "base-mainnet": `${process.env.BASE_ETHERSCAN_API_KEY}`,
    },
    customChains: [
      {
        network: "base-mainnet",
        chainId: 8453,
        urls: {
          apiURL: "https://api.basescan.org/api",
          browserURL: "https://basescan.org",
        },
      },
      {
        network: "base-sepolia",
        chainId: 84532,
        urls: {
          apiURL: "https://api-sepolia.basescan.org/api",
          browserURL: "https://sepolia.basescan.org",
        },
      },
      {
        network: "blast_sepolia",
        chainId: 168587773,
        urls: {
          apiURL:
            "https://api.routescan.io/v2/network/testnet/evm/168587773/etherscan",
          browserURL: "https://testnet.blastscan.io",
        },
      },
    ],
  },
  sourcify: {
    // Disabled by default
    // Doesn't need an API key
    enabled: false,
  },
};
export default config;
