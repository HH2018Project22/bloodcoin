package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/HH2018Project22/bloodcoin/blockchain"
	"github.com/btcsuite/btcutil/base58"
)

var (
	prescriptionCommand = flag.NewFlagSet("prescription", flag.ExitOnError)
	data                string
	prescriptionQuiet   bool
)

func init() {
	prescriptionCommand.StringVar(&data, "data", data, "Prescription data")
	prescriptionCommand.BoolVar(&prescriptionQuiet, "quiet", prescriptionQuiet, "Quiet mode")
}

func doPrescription(args []string) {

	if err := prescriptionCommand.Parse(args); err != nil {
		panic(err)
	}

	bc := getBlockchain()

	if !prescriptionQuiet {
		log.Println("adding prescription")
	}

	prescription := &blockchain.Prescription{}
	if err := json.Unmarshal([]byte(data), prescription); err != nil {
		panic(err)
	}

	prescriptionEvent := blockchain.NewPrescriptionEvent(prescription)
	block, err := bc.AddEvent(prescriptionEvent)

	if err != nil {
		panic(err)
	}

	base58Hash := base58.Encode(block.Hash)
	if !prescriptionQuiet {
		log.Println("Block:", base58Hash)
		log.Println("saving blockchain")
	} else {
		fmt.Println(base58Hash)
	}

	if err := bc.Save(blockchainPath); err != nil {
		panic(err)
	}

}
