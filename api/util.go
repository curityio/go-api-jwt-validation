package main

import (
	"encoding/base64"
	"math/big"
	"log"
	"crypto/tls"
	"io/ioutil"
	"crypto/rsa"
	"crypto/ed25519"
	"encoding/json"
	"net/http"
	"strings"
)

func decodeBase64BigInt(s string) *big.Int {
	buffer, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	if err != nil {
		log.Fatalf("failed to decode base64: %v", err)
	}

	return big.NewInt(0).SetBytes(buffer)
}

// Get Key from JWKS endpoint
func getKey(jwksEndpoint string) *rsa.PublicKey{
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //For demo only to handle self-sign certs
	client:= &http.Client{Transport: customTransport}
    resp, err := client.Get(jwksEndpoint)

    if err != nil {
        log.Println(err)
    }

    defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	
	type Key struct {
		N string `json:"n"`
		E string `json:"e"`
	}

	var Keys struct {
		Key []Key `json:"keys"`
	}

	if err := json.Unmarshal(body, &Keys);
	err != nil{
		log.Println(err)
	}

	//TODO: Handle multiple keys
	N := decodeBase64BigInt(Keys.Key[0].N)
	E := int(decodeBase64BigInt(Keys.Key[0].E).Int64())

	return &rsa.PublicKey{
		N: N,
		E: E,
	}
}

// Get EdDSAKey from JWKS endpoint
func getEdDSAKey(jwksEndpoint string) ed25519.PublicKey {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //For demo only to handle self-sign certs
	client:= &http.Client{Transport: customTransport}
    resp, err := client.Get(jwksEndpoint)

    if err != nil {
        log.Println(err)
    }

    defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	
	type Key struct {
		X string `json:"x"`
	}

	var Keys struct {
		Key []Key `json:"keys"`
	}

	if err := json.Unmarshal(body, &Keys);
	err != nil{
		log.Println(err)
	}

	//TODO: Handle multiple keys
	X, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(Keys.Key[0].X);

	if err != nil {
		log.Fatalf("failed to decode base64: %v", err)
	}

	return ed25519.PublicKey(X)
}

func checkScopes(requiredScopes []string, providedScopes string) bool {
	
	for _, value := range requiredScopes {
		if(!strings.Contains(providedScopes, value)){
			return false
		}
	}
	
	return true
}