# GO signature mint

## How it works

The server is written in Go and uses the Gin framework. It has a single endpoint, `/generate`, which takes a JSON body with a `address` field and returns a signed payload for [signature minting](https://portal.thirdweb.com/go/erc721_signature_minting) an ERC-721 token. We then use this server in a vite react application to signature mint NFTs using the thirdweb React SDK.

## How to set it up

1. Install Go if you haven't already. You can find instructions [here](https://golang.org/doc/install).
2. Create a new project by cloning this repository and renaming the folder.
3. Update the contract addresses in [`common.go`](/backend/handlers/common.go) and [`App.tsx`](/frontend/src/App.tsx).
4. Export your wallet private key from your wallet and add it to the .env file in the [backend](/backend/) folder.

```bash
PRIVATE_KEY=your_private_key
```

Using private keys as an env variable is vulnerable to attacks and is not best practice. We are doing it in this guide for the sake of brevity, but we strongly recommend using a [secret manager to store your private key](https://portal.thirdweb.com/typescript/sdk.thirdwebsdk.fromwallet).

## How to run it

1. Setup the project using the following command:

```bash
make setup
```

This command will install the dependencies for both the frontend and the backend.

2. Open a new terminal and run the frontend using the following command:

```bash
make run-frontend
```

You will now be able to see a connect wallet button in front of you if open up the site.

3. Run the backend server using the following command:

```bash
make run-backend
```

This command will start the backend server on port 8080.

You can test the server by making a request to the `/generate` endpoint with a JSON body containing an `address` field. On doing that you will receive a signed payload for minting an ERC-721 token like this:

![generate output](https://blog.thirdweb.com/content/images/size/w1600/2023/05/make-a-request-to-the-API-and-recieve-the-signature.png)

4. You will now need to replace the website origin in the [main.go](/backend/main.go) to allow the frontend to make requests to the backend. Replace the following line:

```go
		AllowOrigins: []string{"http://192.168.29.209:5173"},
```

With your frontend origin.

5. You can now go to the frontend and click on the connect wallet button. Once your wallet is connected click on mint NFT and you will see a transaction pop up. Once you confirm it your NFT will be minted.

![mint NFT](https://res.cloudinary.com/didkcszrq/image/upload/v1685820472/SCR-20230604-btqu_skvwbr.png)

## Learn More

To learn more about thirdweb, react and Go, take a look at the following resources:

- [thirdweb go Documentation](https://portal.thirdweb.com/go) - learn about our Go SDK.
- [thirdweb React Documentation](https://portal.thirdweb.com/react) - learn about our React SDK.
- [thirdweb Portal](https://portal.thirdweb.com) - check our guides and development resources.

You can check out [the thirdweb GitHub organization](https://github.com/thirdweb-dev) - your feedback and contributions are welcome!

## Join our Discord!

For any questions, suggestions, join our discord at [https://discord.gg/thirdweb](https://discord.gg/thirdweb).
