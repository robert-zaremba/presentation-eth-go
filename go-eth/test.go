package main

import (
	"math/big"
	"os"

	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//start deploy OMIT
func ownerTx(pkFile string, password string) (res *bind.TransactOpts) {
	// note: remember to check the errors!
	kr, err := os.Open(pkFile)
	if err != nil {
		panic(err)
	}
	if res, err = bind.NewTransactor(kr, password); err != nil {
		panic(err)
	}
	return
}

func deploy() *Power {
	var teller, tech common.Address
	endPoint := "http://localhost:10001"
	client, _ := ethclient.Dial(endPoint)
	tx := ownerTx("owner.json", "password")
	_, _, power, _ := DeployPower(tx, client, teller, tech)
	fmt.Println("PowerContract", power)
	return power
}

//end deploy OMIT

//start bal OMIT
func balance(contract *Power, user common.Address) *big.Int {
	b, _ := contract.BalanceOf(nil, user)
	return b
}

//end bal OMIT

//start pay OMIT
func payBill(contract *Power, user common.Address, amount *big.Int) {
	tx, _ := contract.PayBill(ownerTx("teller.json", "password2"), user, amount)
	fmt.Println("tx hash : ", tx.Hash().Hex())
}

//end pay OMIT

func main() {
	deploy()
}
