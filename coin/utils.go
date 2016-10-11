/*
Copyright Hydursio Labs Inc. 2016 All Rights Reserved.
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
	"crypto/sha256"
	"fmt"

	"github.com/golang/protobuf/proto"
)

// HashTx get tx's hash, using sha256
func HashTx(tx *Transaction) (string, error) {
	txBytes, err := proto.Marshal(tx)
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	hash.Write(txBytes)

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// ParseBank parse bank object
func ParseBank(bankBytes []byte) (*Bank, error)  {
	bank := &Bank{}

	if err := proto.Unmarshal(bankBytes, bank); err != nil {
		return nil, err
	}

	return bank, nil
}