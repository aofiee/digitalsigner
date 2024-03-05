# Build for Linux
GOOS=linux GOARCH=amd64 go build -o myprogram_linux

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o myprogram.exe

# Build for macOS (optional, as you're already on macOS)
GOOS=darwin GOARCH=amd64 go build -o myprogram_mac

# Run the program
./bin/osx/digitalsigner

```bash
digitalsigner is a CLI tool for signing and verifying digital signature.:

This application is a tool to sign and verify digital signature.
It is a CLI tool that can be used to sign and verify digital signature.

Usage:
  digitalsigner [command]

Available Commands:
  completion         Generate the autocompletion script for the specified shell
  generate-key-pair  A command to generate a key pair.
  generate-signature A command to generate a digital signature.
  help               Help about any command
  verifydata         A command to verify the integrity of a file.

Flags:
  -h, --help     help for digitalsigner
  -t, --toggle   Help message for toggle

Use "digitalsigner [command] --help" for more information about a command.
```

## Generate Key Pair

./bin/osx/digitalsigner generate-key-pair -f example
```bash
INFO[0000] Generating key pair... example               
INFO[0000] Private key saved to example_private_key.pem 
INFO[0000] Public key saved to example_public_key.pem   
```

## Generate Signature

./bin/osx/digitalsigner generate-signature -f testfile.txt -p example_private_key.pem -s example_signature
```bash
INFO[0000] Generating signature...                      
INFO[0000] Signature saved to example_signature
```

## Verify Signature

./bin/osx/digitalsigner verifydata -f testfile.txt -p example_public_key.pem -s example_signature 
```bash
INFO[0000] Verifying signature...                       
INFO[0000] Signature verified successfully
```