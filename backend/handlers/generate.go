package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

type Body struct {
	Address string `json:"address"`
}

func GenerateSignature(c *gin.Context) {
	// Get contract
	contract := getContract()

	// Get address
	var address Body
	c.BindJSON(&address)

	// the payload to sign
	payload := &thirdweb.Signature721PayloadInput{
		To:                   address.Address,
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &thirdweb.NFTMetadataInput{
			Name:        "My NFT",
			Description: "This is my cool NFT",
			Image:       "https://gateway.ipfscdn.io/ipfs/QmXXjx3aJCs9W9mN35Aade6etSoceqMk8ykkasbB87MaLt/0.png",
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
	}

	signedPayload, err := contract.Signature.Generate(context.Background(), payload)

	// Panic if error
	if err != nil {
		panic(err)
	}

	println(signedPayload)

	// return signed payload
	c.JSON(200, gin.H{
		"signedPayload": signedPayload,
	})
}
