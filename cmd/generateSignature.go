/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// generateSignatureCmd represents the generateSignature command
var generateSignatureCmd = &cobra.Command{
	Use:   "generate-signature",
	Short: "A command to generate a digital signature.",
	Long: `A command to generate a digital signature from a file.

This command is used to generate a digital signature from a file. It will use the private key to sign the file and generate a digital signature. The digital signature can be used to verify the integrity of the file.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Println("Generating signature...")
		GenerateSignature(fileName, privateKey, signatureFileName)
	},
}

func init() {
	generateSignatureCmd.Flags().StringVarP(&fileName, "file", "f", "", "The name of the file to generate a digital signature for.")
	generateSignatureCmd.Flags().StringVarP(&privateKey, "private-key", "p", "", "The private key to use to generate the digital signature.")
	generateSignatureCmd.Flags().StringVarP(&signatureFileName, "signature-file", "s", "", "The name of the file to save the digital signature to.")
	rootCmd.AddCommand(generateSignatureCmd)
}

func GenerateSignature(file string, privateKeyPEMFile string, signatureFileName string) {
	privateKeyPEM, err := os.ReadFile(privateKeyPEMFile)
	if err != nil {
		logrus.Error("Failed to read private key:", err)
		return
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		logrus.Error("Failed to decode PEM block containing private key")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logrus.Error("Failed to parse private key:", err)
		return
	}

	data, err := os.ReadFile(file)
	if err != nil {
		logrus.Error("Failed to read file:", err)
		return
	}

	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		logrus.Println("Failed to sign file:", err)
		return
	}

	signatureFile, err := os.Create(signatureFileName)
	if err != nil {
		logrus.Error("Failed to create signature file:", err)
		return
	}
	defer signatureFile.Close()

	_, err = signatureFile.Write(signature)
	if err != nil {
		logrus.Error("Failed to write signature to file:", err)
		return
	}

	logrus.Println(fmt.Sprintf("Signature saved to %s", signatureFileName))
}
