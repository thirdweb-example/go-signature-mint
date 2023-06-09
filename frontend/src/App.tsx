import { Web3Button, useAddress, useContract } from "@thirdweb-dev/react";
import "./styles/Home.css";

export default function Home() {
  const address = useAddress();
  const { contract: nftCollection } = useContract(
    "0x9D6597681B67FF75034dA5CF4F5e01Ef64F478cF",
    "nft-collection"
  );

  const mint = async () => {
    try {
      const request = await fetch("http://localhost:8080/generate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ address }),
      });

      const signature = await request.json();
      console.log(signature);

      const nft = await nftCollection?.erc721.mint(signature);
      console.log(nft);
      return nft;
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="container">
      <main className="main">
        <Web3Button
          action={mint}
          contractAddress="0x9D6597681B67FF75034dA5CF4F5e01Ef64F478cF"
        >
          Mint
        </Web3Button>
      </main>
    </div>
  );
}
