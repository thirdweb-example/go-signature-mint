package handler

import (
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"os"
)

func getContract() *thirdweb.NFTCollection {
	sdk := initSdk()
	collectionAddress := "0x9D6597681B67FF75034dA5CF4F5e01Ef64F478cF"

	contract, _ := sdk.GetNFTCollection(collectionAddress)

	return contract
}

func initSdk() *thirdweb.ThirdwebSDK {
	privateKey := os.Getenv("PRIVATE_KEY")

	sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
		PrivateKey: privateKey,
	})
	if err != nil {
		panic(err)
	}

	return sdk
}
