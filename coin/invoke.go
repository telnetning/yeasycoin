/*
Copyright Hydrusio Labs Inc. 2016 All Rights Reserved.
Written by mint.zhao.chiu@gmail.com. github.com: https://www.github.com/mintzhao

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package coin

import (
	"fmt"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

func (coin *Yeasycoin) createBank(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ErrInvalidParams
	}

	bankName := args[0]

	// get max bank id now
	maxBankIdBytes, err := stub.GetState(max_bankId)
	if err != nil {
		return nil, err
	}
	maxBankId, err := strconv.ParseInt(string(maxBankIdBytes), 10, 64)
	if err != nil {
		return nil, err
	}
	bankId := maxBankId + 1

	bank := &Bank{
		Id:   bankId,
		Name: bankName,
	}

	// put bank into blockchain
	bankBytes, err := proto.Marshal(bank)
	if err != nil {
		return nil, err
	}
	if err := stub.PutState(fmt.Sprintf("bank_%v", bankId), bankBytes); err != nil {
		return nil, err
	}
	if err := stub.PutState(max_bankId, []byte(strconv.FormatInt(bankId, 10))); err != nil {
		return nil, err
	}

	return bankBytes, nil
}

func (coin *Yeasycoin) createCompany(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, ErrInvalidParams
	}

	companyName := args[0]

	// get max company id
	maxCompanyIdBytes, err := stub.GetState(max_companyId)
	if err != nil {
		return nil, err
	}
	maxCompanyId, err := strconv.ParseInt(string(maxCompanyIdBytes), 10, 64)
	if err != nil {
		return nil, err
	}
	companyId := maxCompanyId + 1

	company := &Company{
		Id:   companyId,
		Name: companyName,
	}

	companyBytes, err := proto.Marshal(company)
	if err != nil {
		return nil, err
	}
	if err := stub.PutState(fmt.Sprintf("company_%v", companyId), companyBytes); err != nil {
		return nil, err
	}
	if err := stub.PutState(max_companyId, []byte(strconv.FormatInt(companyId, 10))); err != nil {
		return nil, err
	}

	return companyBytes, nil
}

func (coin *Yeasycoin) issueCoin(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, ErrInvalidParams
	}

	coinNumber, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return nil, err
	}
	// consider: if using time.Now() at chaincode, maybe the tx can not be the same at other vps
	timestamp := args[1]

	tx := &Transaction{
		FromType:  Transaction_FROM_CENTERBANK,
		FromId:    0,
		ToType:    Transaction_TO_CENTERBANK,
		ToId:      0,
		Timestamp: timestamp,
		Number:    coinNumber,
	}
	if txHash, err := HashTx(tx); err != nil {
		return nil, err
	} else {
		tx.Id = txHash
	}

	// get centerbank
	cbankBytes, err := stub.GetState("bank_0")
	if err != nil {
		return nil, err
	}
	bank, err := ParseBank(cbankBytes)
	if err != nil {
		return nil, err
	}
	bank.TotalNumber += coinNumber
	bank.RestNumber += coinNumber

	cbankBytes, err = proto.Marshal(bank)
	if err != nil {
		return nil, err
	}
	if err := stub.PutState("bank_0", cbankBytes); err != nil {
		return nil, err
	}

	return proto.Marshal(tx)
}

func (coin *Yeasycoin) issueCoinToBank(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}

func (coin *Yeasycoin) issueCoinToCp(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}

func (coin *Yeasycoin) transfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}
