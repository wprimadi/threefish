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
go get -u github.com/wprimadi/threefish@v1.0.1
```
or 
```go
go get -u github.com/wprimadi/threefish@latest
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

	"github.com/wprimadi/threefish"
)

func main() {
	// Define a 512-bit key (64 bytes) and 16-byte tweak
	key := []byte("1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	tweak := []byte("abcdefghijklmnop")
	plaintext := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. 1234567")

	// Initialize Threefish-512
	cipher, err := threefish.NewThreefish(Threefish1024, key, tweak)
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

## Key and Tweak in Threefish
### Key
The key is the main secret value used to encrypt and decrypt data. Threefish supports key sizes of 256, 512, and 1024 bits, depending on the chosen block size. A strong and unpredictable key is essential to ensure the security of encrypted data.
- The key length must match the block size (e.g., 64 bytes for Threefish-512).
- If the key is weak or predictable, the encryption can be compromised.
- Secure key generation should use a cryptographic random number generator (e.g., crypto/rand in Golang).

### Tweak
The tweak is an additional input that modifies how encryption is applied to data, similar to a nonce or initialization vector (IV) in other encryption schemes. It prevents the same plaintext from always producing the same ciphertext, adding an extra layer of security.
- Threefish uses a 128-bit tweak (16 bytes), divided into two 64-bit words.
- The tweak should remain the same between encryption and decryption.
- Unlike IVs in AES, tweaks can be structured with metadata, counters, or block indices to support unique encryption behavior.

A correct key and tweak combination is critical to ensure that decryption produces the original data. If either value is incorrect or missing, decryption will fail and return unreadable data.

## Key Size and Threefish Mode
| Threefish Mode | Key Size (bytes)     | Tweak Size (bytes) |
|----------------|----------------------|--------------------|
| Threefish-256  | 32 bytes (256-bit)   | 16 bytes           |
| Threefish-512  | 64 bytes (512-bit)   | 16 bytes           |
| Threefish-1024 | 128 bytes (1024-bit) | 16 bytes           |

## License
This project is licensed under the MIT License.





