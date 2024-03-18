package main

import (
	"crypto/rand"
	"eintrag/pkg/eintrag"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/jamesruan/sodium"
	"math/big"
	"os"
)

const KeySize = 64
const Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

func main() {
	configFile := flag.String("config", "config.json", "")
	doGenerate := flag.Bool("generate-key", false, "")
	updateMaster := flag.Bool("generate-master-key", false, "")
	flag.Parse()

	config := eintrag.Config{
		Listen:       "127.0.0.1",
		Port:         8080,
		DatabaseType: "postgres",
	}

	if _, err := os.Stat(*configFile); err == nil {
		if err := readConfigurationFile(*configFile, &config); nil != err {
			panic(err)
		}

		if *doGenerate || "" == config.SigningKey {
			config.SigningKey, _ = generateKey()
		}

		if *updateMaster || "" == config.MasterKey {
			config.MasterKey = generateMasterKey()
		}

	} else if errors.Is(err, os.ErrNotExist) {
		config.SigningKey, _ = generateKey()
		config.MasterKey = generateMasterKey()
	} else {
		panic(err)
	}

	content, err := json.MarshalIndent(config, "", "  ")
	if nil != err {
		panic(err)
	}

	err = os.WriteFile(*configFile, content, 0600)
	if nil != err {
		panic(err)
	}

	fmt.Printf("Wrote to file: %s", *configFile)
}

func readConfigurationFile(path string, config *eintrag.Config) error {
	content, err := os.ReadFile(path)
	if nil != err {
		return err
	}

	if err := json.Unmarshal(content, config); nil != err {
		return err
	}

	return nil
}

func generateKey() (string, error) {
	ret := make([]byte, KeySize)

	for i := 0; i < KeySize; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(Alphabet))))
		if nil != err {
			return "", err
		}

		ret[i] = Alphabet[num.Int64()]
	}

	return string(ret), nil
}

func generateMasterKey() string {
	masterKey := sodium.MakeMasterKey()

	output := make([]byte, base64.StdEncoding.EncodedLen(len(masterKey.Bytes)))
	base64.StdEncoding.Encode(output, masterKey.Bytes)

	return string(output)
}
