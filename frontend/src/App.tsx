import { Web3Button, useAddress, useContract } from "@thirdweb-dev/react";
import "./styles/Home.css";
import axios from "axios";

export default function Home() {
  const address = useAddress();
  const { contract: nftCollection } = useContract(
    "0x9D6597681B67FF75034dA5CF4F5e01Ef64F478cF",
    "nft-collection"
  );

  const mint = async () => {
    try {
      const signature = await axios.post("http://localhost:8080/generate", {
        address,
      });

      const nft = nftCollection?.erc721.mint(signature.data);
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
