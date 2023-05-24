# GO signature mint

## How it works

The server is written in Go and uses the Gin framework. It has a single endpoint, `/generate`, which takes a JSON body with a `address` field and returns a signed payload for signature minting an ERC-721 token. We then use this server in a vite react application to signature mint NFTs using the thirdweb React SDK.

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

1. cd into the [backend](/backend/) folder.
2. Run `go run main.go` to start the server.
3. The app will be running on `localhost:8080`.
4. You can test the app by running the vite app in the frontend folder.

## Learn More

To learn more about thirdweb, react and Go, take a look at the following resources:

- [thirdweb go Documentation](https://portal.thirdweb.com/go) - learn about our Go SDK.
- [thirdweb React Documentation](https://portal.thirdweb.com/react) - learn about our React SDK.
- [thirdweb Portal](https://portal.thirdweb.com) - check our guides and development resources.

You can check out [the thirdweb GitHub organization](https://github.com/thirdweb-dev) - your feedback and contributions are welcome!

## Join our Discord!

For any questions, suggestions, join our discord at [https://discord.gg/thirdweb](https://discord.gg/thirdweb).
