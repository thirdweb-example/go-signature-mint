package handler

import (
	"os"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func getContract() *thirdweb.NFTCollection {
	sdk := initSdk()
	collectionAddress := "0x9D6597681B67FF75034dA5CF4F5e01Ef64F478cF"

	contract, _ := sdk.GetNFTCollection(collectionAddress)

	return contract
}

func initSdk() *thirdweb.ThirdwebSDK {
	privateKey := os.Getenv("WALLET_PRIVATE_KEY")
	secretKey := os.Getenv("THIRDWEB_SECRET_KEY")

	sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
		PrivateKey: privateKey,
		SecretKey:  secretKey,
	})
	if err != nil {
		panic(err)
	}

	return sdk
}
