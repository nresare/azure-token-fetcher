package main

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	"os"
)

func main() {
	// the endpoint queried, follows the pattern "https://login.microsoftonline.com/<Tenant ID>"
	authority := "https://login.microsoftonline.com/1b5955bf-2426-4ced-8412-9fe81bb8bca4"
	// the client id
	clientId := "a0a86205-89d9-4f66-80af-9b5c35dd228a"
	// the target scope to encode as aud into the requested token
	scope := "ee7470ec-7849-42ed-85c5-1a22cf1a6774/.default"
	result, err := fetchToken(
		authority,
		clientId,
		"cert.pem",
		"key.pem",
		scope)
	if err != nil {
		fmt.Printf("Found an error: %s", err)

	}
	fmt.Printf("The token: '%s'\n", result.AccessToken)
}

func fetchToken(authority, clientId, certPath, keyPath, scope string) (*confidential.AuthResult, error) {
	certs, err := getCerts(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read cert from '%s': %w", certPath, err)
	}

	privateKey, err := getPrivateKey(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key from '%s': %w", keyPath, err)
	}

	cred, err := confidential.NewCredFromCert(certs, privateKey)
	if err != nil {
		return nil, err
	}

	client, err := confidential.New(authority, clientId, cred)
	if err != nil {
		return nil, err
	}

	result, err := client.AcquireTokenByCredential(context.Background(), []string{scope})
	return &result, err
}

func getCerts(path string) ([]*x509.Certificate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return []*x509.Certificate{cert}, nil
}

func getPrivateKey(path string) (crypto.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
