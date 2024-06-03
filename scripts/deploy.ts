import { ethers, run } from "hardhat";

async function main() {
  const params = ["0x1D3f1A700b1e6df89E9Ef56689B6Ad4B5b932A88"];
  const c = await ethers.deployContract("CrunchProtocol", [params[0]]);
  const caddress = await c.getAddress();
  // const caddress = `0x52806ae50D45A2c9f1836eCA255E3A5b2108ecC4`;
  console.log(`deployed to ${caddress}`);
  //Delay a little
  // await new Promise((resolve) => setTimeout(resolve, 20000));
  // Verify the contract
  await run("verify:verify", {
    address: caddress,
    constructorArguments: params,
    contract: "contracts/CrunchProtocol.sol:CrunchProtocol",
  });
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
