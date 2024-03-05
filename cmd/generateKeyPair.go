/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// generateKeyPairCmd represents the generateKeyPair command
var generateKeyPairCmd = &cobra.Command{
	Use:   "generate-key-pair",
	Short: "A command to generate a key pair.",
	Long: `A command to generate a key pair.

This command is used to generate a key pair. It will generate a public and private key using the RSA algorithm. The key pair will be used to sign and verify digital signatures.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Println("Generating key pair...", fileName)
		GenerateKeyPair(fileName)
	},
}

func init() {

	generateKeyPairCmd.Flags().StringVarP(&fileName, "file", "f", "key_pair", "The name of the file to save the key pair to.")
	rootCmd.AddCommand(generateKeyPairCmd)
}

func GenerateKeyPair(fileKeyPairName string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.Error("Failed to generate RSA key pair:", err)
		return
	}

	privateKeyFile, err := os.Create(fmt.Sprintf("%s_private_key.pem", fileKeyPairName))
	if err != nil {
		logrus.Error("Failed to create private key file:", err)
		return
	}
	defer privateKeyFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		logrus.Error("Failed to write private key to file:", err)
		return
	}

	logrus.Println(fmt.Sprintf("Private key saved to %s_private_key.pem", fileKeyPairName))

	publicKeyFile, err := os.Create(fmt.Sprintf("%s_public_key.pem", fileKeyPairName))
	if err != nil {
		logrus.Error("Failed to create public key file:", err)
		return
	}
	defer publicKeyFile.Close()

	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		logrus.Error("Failed to write public key to file:", err)
		return
	}

	logrus.Println(fmt.Sprintf("Public key saved to %s_public_key.pem", fileKeyPairName))
}
