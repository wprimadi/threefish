# Threefish Golang Implementation

![Go Version](https://img.shields.io/github/go-mod/go-version/wprimadi/threefish) 
![License](https://img.shields.io/github/license/wprimadi/threefish) 
![Stars](https://img.shields.io/github/stars/wprimadi/threefish?style=social) 
![Last Commit](https://img.shields.io/github/last-commit/wprimadi/threefish) 
![Go Report Card](https://goreportcard.com/badge/github.com/wprimadi/threefish) 
![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=wprimadi_threefish&metric=alert_status) 
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows-blue) 

A pure Golang implementation of the **Threefish** block cipher, supporting **256-bit, 512-bit, and 1024-bit** key sizes. This package provides encryption and decryption functionality based on the Threefish block cipher, designed for high-speed performance and security.

## Features
✅ Supports **Threefish-256, Threefish-512, and Threefish-1024**  
✅ Implements **encryption and decryption** functions  
✅ Uses **64-bit word operations** for efficiency  
✅ Compatible with **SonarQube quality gates (no issues, code smells, or security hotspots)**  
✅ **Go Report A+ Grade** compliant  
✅ Easy-to-use API  

## Installation
```go
go get github.com/wprimadi/threefish
```

## Usage
### Import the package
```go
import "github.com/wprimadi/threefish"
```

### Encrypt & Decrypt Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/yourusername/threefish"
)

func main() {
	// Define a 512-bit key (64 bytes) and 16-byte tweak
	key := make([]byte, 64)  // Replace with secure key
	tweak := make([]byte, 16) // Replace with secure tweak
	plaintext := []byte("This is a secret message!")

	// Initialize Threefish-512
	cipher, err := threefish.NewThreefish(threefish.Threefish512, key, tweak)
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt
	encrypted, err := cipher.EncryptBlock(plaintext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", encrypted)

	// Decrypt
	decrypted, err := cipher.DecryptBlock(encrypted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", string(decrypted))
}
```

## License
This project is licensed under the MIT License.





