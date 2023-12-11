package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	configDir       = "config"
	credentialsFile = "credentials.yml.enc"
	masterKeyFile   = "master.key"
	defaultEditor   = "vim" // Change this to your preferred text editor
	masterKeyLength = 32    // 32 bytes for AES-256
)

func main() {
	// Ensure the config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		panic(err)
	}

	credentialsPath := filepath.Join(configDir, credentialsFile)
	masterKeyPath := filepath.Join(configDir, masterKeyFile)

	// Read or create master key
	var masterKey []byte
	if _, err := os.Stat(masterKeyPath); os.IsNotExist(err) {
		// Generate new master key
		masterKey = make([]byte, masterKeyLength)
		if _, err = rand.Read(masterKey); err != nil {
			panic(err)
		}
		// Save master key as hex string
		if err := ioutil.WriteFile(masterKeyPath, []byte(hex.EncodeToString(masterKey)), 0600); err != nil {
			panic(err)
		}
		fmt.Println("New master key generated.")
	} else {
		// Read existing master key
		keyHex, err := ioutil.ReadFile(masterKeyPath)
		if err != nil {
			panic(err)
		}
		masterKey, err = hex.DecodeString(string(keyHex))
		if err != nil {
			panic(err)
		}
	}

	// Decrypt or create the credentials file
	var plaintext []byte
	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		plaintext = []byte("initial: data\n") // Initial data for the credentials file
	} else {
		plaintext, err = decryptConfigFile(credentialsPath, hex.EncodeToString(masterKey))
		if err != nil {
			panic(err)
		}
	}

	// Create a temporary file for editing
	tmpfile, err := ioutil.TempFile("", "credentials-*.yml")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	if _, err := tmpfile.Write(plaintext); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	// Open the temporary file in the editor
	cmd := exec.Command(defaultEditor, tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// Read the edited file
	editedText, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	// Re-encrypt and save the credentials file
	if err := encryptConfigFile(credentialsPath, hex.EncodeToString(masterKey), editedText); err != nil {
		panic(err)
	}

	fmt.Println("Credentials updated successfully.")
}

func encryptConfigFile(filename, keyString string, plaintext []byte) error {
	key, _ := hex.DecodeString(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ioutil.WriteFile(filename, ciphertext, 0644)
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
