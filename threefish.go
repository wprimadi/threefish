package threefish

import (
	"encoding/binary"
	"errors"
	"math/bits"
)

const (
	Threefish256  = 256
	Threefish512  = 512
	Threefish1024 = 1024
)

type Threefish struct {
	blockSize  int
	key        []uint64
	tweak      [3]uint64
	usePadding bool
}

func NewThreefish(size int, key []byte, tweak []byte, usePadding bool) (*Threefish, error) {
	if size != Threefish256 && size != Threefish512 && size != Threefish1024 {
		return nil, errors.New("invalid Threefish block size")
	}

	keyWords := size / 64
	if len(key) != keyWords*8 {
		return nil, errors.New("invalid key length")
	}

	var k = make([]uint64, keyWords+1)
	for i := 0; i < keyWords; i++ {
		k[i] = binary.LittleEndian.Uint64(key[i*8 : (i+1)*8])
	}

	var parity uint64
	for _, word := range k[:keyWords] {
		parity ^= word
	}
	k[keyWords] = parity ^ 0x1BD11BDAA9FC1A22

	var t [3]uint64
	if len(tweak) == 16 {
		t[0] = binary.LittleEndian.Uint64(tweak[:8])
		t[1] = binary.LittleEndian.Uint64(tweak[8:])
		t[2] = t[0] ^ t[1]
	} else {
		return nil, errors.New("invalid tweak length")
	}

	return &Threefish{
		blockSize:  size,
		key:        k,
		tweak:      t,
		usePadding: usePadding,
	}, nil
}

func (tf *Threefish) padToBlockSize(input []byte) []byte {
	padLength := tf.blockSize/8 - len(input)%tf.blockSize/8
	if padLength == 0 {
		return input
	}
	paddedInput := append(input, make([]byte, padLength)...)
	paddedInput[len(input)] = byte(padLength) // Store the padding length
	return paddedInput
}

func (tf *Threefish) unpad(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("input is empty")
	}

	// Get padding length from last byte
	padLength := int(input[len(input)-1])
	if padLength > len(input) {
		return nil, errors.New("invalid padding length")
	}

	return input[:len(input)-padLength], nil
}

func (tf *Threefish) EncryptBlock(input []byte) ([]byte, error) {
	if len(input) != tf.blockSize/8 && !tf.usePadding {
		return nil, errors.New("invalid input length")
	}

	if tf.usePadding {
		// Pad input to block size if necessary
		input = tf.padToBlockSize(input)
	}

	blockWords := tf.blockSize / 64
	plaintext := make([]uint64, blockWords)
	for i := 0; i < blockWords; i++ {
		plaintext[i] = binary.LittleEndian.Uint64(input[i*8 : (i+1)*8])
	}

	ciphertext := tf.encrypt(plaintext)
	output := make([]byte, len(input))
	for i := 0; i < blockWords; i++ {
		binary.LittleEndian.PutUint64(output[i*8:], ciphertext[i])
	}
	return output, nil
}

func (tf *Threefish) DecryptBlock(input []byte) ([]byte, error) {
	if len(input) != tf.blockSize/8 && !tf.usePadding {
		return nil, errors.New("invalid input length")
	}

	blockWords := tf.blockSize / 64
	ciphertext := make([]uint64, blockWords)
	for i := 0; i < blockWords; i++ {
		ciphertext[i] = binary.LittleEndian.Uint64(input[i*8 : (i+1)*8])
	}

	plaintext := tf.decrypt(ciphertext)
	output := make([]byte, len(input))
	for i := 0; i < blockWords; i++ {
		binary.LittleEndian.PutUint64(output[i*8:], plaintext[i])
	}

	if tf.usePadding {
		// Remove padding if used
		plaintext, err := tf.unpad(output)
		if err != nil {
			return nil, err
		}
		return plaintext, nil
	}

	return output, nil
}

func (tf *Threefish) encrypt(block []uint64) []uint64 {
	numRounds := 72 // Standard for Threefish
	for round := 0; round < numRounds; round += 4 {
		for i := 0; i < len(block); i++ {
			block[i] ^= tf.key[i%len(tf.key)]
			block[i] = bits.RotateLeft64(block[i], int(tf.tweak[round%3]%64))
		}
	}
	return block
}

func (tf *Threefish) decrypt(block []uint64) []uint64 {
	numRounds := 72 // Standard for Threefish
	for round := numRounds - 4; round >= 0; round -= 4 {
		for i := 0; i < len(block); i++ {
			block[i] = bits.RotateLeft64(block[i], -int(tf.tweak[round%3]%64))
			block[i] ^= tf.key[i%len(tf.key)]
		}
	}
	return block
}
