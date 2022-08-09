package crypto

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func loadPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	return privateKey
}

func GetPublicAddress() common.Address {
	privateKey := loadPrivateKey()
	publicKey := privateKey.Public()
	publicKeyEcdsa, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyEcdsa)
}

func GetNonce(client *ethclient.Client) uint64 {
	publicAddress := GetPublicAddress()
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatal(err)
	}

	return nonce
}

func GetGasLimit(client *ethclient.Client, to common.Address, data []byte) uint64 {
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &to,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	return gasLimit
}

func GetGasPrice(client *ethclient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

func GetChainID(client *ethclient.Client) *big.Int {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return chainID
}

func GetMethodID(function string) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(function))
	methodId := hash.Sum(nil)[:4]
	return methodId
}

func GetFunctionData(signature string, args ...[]byte) []byte {
	var data []byte
	methodID := GetMethodID(signature)
	data = append(data, methodID...)
	for _, arg := range args {
		padded_arg := common.LeftPadBytes(arg, 32)
		data = append(data, padded_arg...)
	}

	return data
}

func EncodeArgs(abi_str string, name string, args ...interface{}) []byte {
	abi_interface, err := abi.JSON(strings.NewReader(abi_str))
	if err != nil {
		log.Fatal(err)
	}
	encodedArgs, err := abi_interface.Pack(
		name, args...,
	)
	return encodedArgs
}

func GetOpts(client *ethclient.Client) *bind.TransactOpts {
	privateKey := loadPrivateKey()
	chainID := GetChainID(client)
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	opts.Nonce = big.NewInt(int64(GetNonce(client)))
	opts.Value = big.NewInt(0)
	opts.GasPrice = GetGasPrice(client)
	return opts
}

func GenerateTransaction(client *ethclient.Client, to common.Address, value *big.Int, gasPrice *big.Int, data []byte) *types.Transaction {
	nonce := GetNonce(client)
	gasLimit := GetGasLimit(client, to, data)
	if gasPrice == nil {
		gasPrice = GetGasPrice(client)
	}
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	return tx
}

func SignTransaction(client *ethclient.Client, tx *types.Transaction) *types.Transaction {
	chainID := GetChainID(client)
	privateKey := loadPrivateKey()
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx
}

func SendTransaction(client *ethclient.Client, signedTx *types.Transaction) {
	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
