package controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

// Generate a random key of 32 bytes for AES-256 encryption
var encryptionKey = []byte("b7198d325062f753f487db02e97ec025")
var key = []byte("b7198d325062f753f487db02e97ec0257d63a7c62ee1aa05e84764c36f746c66")

// EncryptBytes encrypts the input byte slice using AES-256 encryption
func EncryptBytes(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Add padding to the data if needed
	data = addPadding(data)

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

// DecryptBytes decrypts the input byte slice using AES-256 encryption
func DecryptBytes(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	// Remove padding from the decrypted data
	data := removePadding(ciphertext)

	return data, nil
}

// Padding functions to ensure the data length is a multiple of the block size
func addPadding(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	paddingBytes := make([]byte, padding)
	return append(data, paddingBytes...)
}

func removePadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

func generateAESKey() ([]byte, error) {
	inputKey := "b7198d325062f753f487db02e97ec025b7198d325062f753f487db02e97ec025"

	// Convert the input string to a byte slice
	key, err := hex.DecodeString(inputKey)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	// Check if the byte slice is exactly 32 bytes long
	if len(key) != 32 {
		log.Printf("%+v", len(key))
		log.Printf("%+v", "input key must decode to exactly 32 bytes")
		return nil, errors.New("input key must decode to exactly 32 bytes")
	}

	return key, nil
}

func EncryptAndSave(srcFile *multipart.FileHeader, outputFilePath string, fileName string) error {
	log.Printf("InsideEncryptSave")
	outputFileName := fmt.Sprintf("%s/%s", outputFilePath, fileName)
	src, err := srcFile.Open()
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	defer src.Close()
	log.Printf("generateRandomIV")
	iv, err := generateRandomIV()
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	log.Printf("outputFileName")
	if err := os.MkdirAll(outputFilePath, os.ModePerm); err != nil {
		// Handle the error if directory creation fails
		log.Printf("Error creating directory: %+v", err)
		return err
	}
	dst, err := os.Create(outputFileName)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	defer dst.Close()

	log.Printf("NewCipher")
	encryptionKey1, _ := generateAESKey()
	log.Printf("%+v", encryptionKey1)
	block, err := aes.NewCipher(encryptionKey1)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	log.Printf("NewCFBEncrypter")
	stream := cipher.NewCFBEncrypter(block, iv)
	_, err = dst.Write(iv)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	log.Printf("src.Read")
	buf := make([]byte, 4096)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("%+v", err)
				return err
			}
			break
		}

		stream.XORKeyStream(buf[:n], buf[:n])
		_, err = dst.Write(buf[:n])
		if err != nil {
			log.Printf("%+v", err)
			return err
		}
	}

	return nil
}

func DecryptFile(inputFileName string, outputFileName string) error {
	// Open the encrypted file
	src, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer src.Close()

	// Read the IV from the beginning of the file
	iv := make([]byte, aes.BlockSize)
	_, err = src.Read(iv)
	if err != nil {
		return err
	}

	decryptionKey, _ := generateAESKey()

	// Create the AES cipher block
	block, err := aes.NewCipher(decryptionKey)
	if err != nil {
		return err
	}

	// Create a stream for decryption
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create the output file for writing the decrypted data
	dst, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Read and decrypt the file content
	buf := make([]byte, 4096)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		// Decrypt the data
		stream.XORKeyStream(buf[:n], buf[:n])

		// Write the decrypted data to the output file
		_, err = dst.Write(buf[:n])
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDecryptFile(inputFileName string) ([]byte, error) {
	// Open the encrypted file
	src, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Read the IV from the beginning of the file
	iv := make([]byte, aes.BlockSize)
	_, err = src.Read(iv)
	if err != nil {
		return nil, err
	}

	decryptionKey, _ := generateAESKey()

	// Create the AES cipher block
	block, err := aes.NewCipher(decryptionKey)
	if err != nil {
		return nil, err
	}

	// Create a stream for decryption
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create a buffer to collect the decrypted data
	var decryptedData bytes.Buffer

	// Read and decrypt the file content
	buf := make([]byte, 4096)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		// Decrypt the data
		stream.XORKeyStream(buf[:n], buf[:n])

		// Write the decrypted data to the buffer
		_, err = decryptedData.Write(buf[:n])
		if err != nil {
			return nil, err
		}
	}

	return decryptedData.Bytes(), nil
}

func generateRandomIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	_, err := io.ReadFull(rand.Reader, iv)
	return iv, err
}
