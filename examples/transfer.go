package main

import (
	"fmt"
	"log"

	"github.com/tgpxdev/cryptobot-sdk-golang/cryptobot"
)

func transfer(client *cryptobot.Client) {
	transfer, err := client.Transfer(cryptobot.TransferRequest{
		UserID:                  1,
		Asset:                   cryptobot.TON,
		Amount:                  "10.5",
		SpendID:                 "",
		Comment:                 "Debt",
		DisableSendNotification: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(
		"ID - %v, UserID - %s, Status - %s, Amount - %s, Asset - %s, Comment - %s, CompletedAt - %s \n",
		transfer.ID,
		transfer.UserID,
		transfer.Status,
		transfer.Amount,
		transfer.Asset,
		transfer.Comment,
		transfer.CompletedAt,
	)
}
