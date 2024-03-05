/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// verifydataCmd represents the verifydata command
var verifydataCmd = &cobra.Command{
	Use:   "verifydata",
	Short: "A command to verify the integrity of a file.",
	Long: `A command to verify the integrity of a file.

This command is used to verify the integrity of a file. It will use the public key to verify the digital signature and check the integrity of the file. The digital signature can be used to verify the integrity of the file.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Println("Verifying signature...")
		VerifySignature(fileName, signatureFileName, publicKey)
	},
}

func init() {
	verifydataCmd.Flags().StringVarP(&fileName, "file", "f", "", "The name of the file to verify the integrity of.")
	verifydataCmd.Flags().StringVarP(&signatureFileName, "signature-file", "s", "", "The name of the file containing the digital signature.")
	verifydataCmd.Flags().StringVarP(&publicKey, "public-key", "p", "", "The public key to use to verify the digital signature.")
	rootCmd.AddCommand(verifydataCmd)
}

func VerifySignature(fileName, signatureFilename, publicKeyPEMFile string) {
	publicKeyPEM, err := os.ReadFile(publicKeyPEMFile)
	if err != nil {
		logrus.Error("Failed to read public key:", err)
		return
	}

	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		logrus.Error("Failed to decode PEM block containing public key")
		return
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		logrus.Error("Failed to read file:", err)
		return
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		logrus.Error("Failed to parse public key:", err)
		return
	}

	signature, err := os.ReadFile(signatureFilename)
	if err != nil {
		logrus.Error("Failed to read signature:", err)
		return
	}

	hash := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		logrus.Error("Failed to verify signature:", err)
		return
	}

	logrus.Println("Signature verified successfully")
}
