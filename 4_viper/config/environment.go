package config

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
)

type Environment struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadEnv(path string) (env Environment, err error) {
	// first version of reading application.yaml
	// viper.AddConfigPath(path)
	// viper.SetConfigName("application")
	// viper.SetConfigType("yaml")

	// viper.AutomaticEnv() // will use system's environment variables

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return
	// }

	// err = viper.Unmarshal(&env)
	// println(fmt.Sprintf("checking env: %s", env.DBDriver))
	// return

	// second version of reading credential.yml.enc
	masterKeyPath := filepath.Join(path, "master.key")
	credentialsPath := filepath.Join(path, "credentials.yml.enc")

	// Read and decode master key
	keyHex, err := ioutil.ReadFile(masterKeyPath)
	if err != nil {
		return
	}
	masterKey, err := hex.DecodeString(string(keyHex))
	if err != nil {
		return
	}

	// Decrypt the credentials file
	decryptedContent, err := decryptConfigFile(credentialsPath, hex.EncodeToString(masterKey))
	if err != nil {
		return
	}

	// Set the config type and read the decrypted content with viper
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // will use system's environment variables
	if err = viper.ReadConfig(bytes.NewBuffer(decryptedContent)); err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}

func decryptConfigFile(filename, keyString string) ([]byte, error) {
	key, _ := hex.DecodeString(keyString)
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
