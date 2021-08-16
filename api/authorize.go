package main

import (
	"encoding/json"
	"log"
	"time"
	"encoding/base64"
	"strings"
	"net/http"
	"github.com/gbrlsnchs/jwt/v3"
	"os"
)

func Authorize(writer http.ResponseWriter, request *http.Request) bool{

	// Check if Authorization header is available
	header := request.Header.Get("Authorization")
	if header == "" {
		log.Println("No Access Token found!")
		writer.WriteHeader(401)
		writer.Write([]byte("No Access Token found!"))
		return false
	}

	headerParts := strings.Split(header, " ")
	jwtToken := headerParts[1]

	//Create Algorithm interface to use for JWT verification
	hs := jwt.NewRS256(jwt.RSAPublicKey(getKey(os.Getenv("JWKS"))))
	
	var pl jwt.Payload

	//Validating the alg field in header
	if _, err := jwt.Verify([]byte(jwtToken), hs, &pl, jwt.ValidateHeader); 
	err != nil {
		log.Println("Error validating alg in header: ", err.Error())
		writer.WriteHeader(401)
		return false
	}

	//Verifies JWT signature
	if _, err := jwt.Verify([]byte(jwtToken), hs, &pl);
	err != nil {
		log.Println("Error verifying JWT signature: ", err.Error())
		writer.WriteHeader(401)
		return false
	}

	aud := jwt.Audience{os.Getenv("AUD")}
	iss := os.Getenv("ISS")
	now := time.Now()
	
	nbfValidator := jwt.NotBeforeValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now)
	issValidator := jwt.IssuerValidator(iss)	
	audValidator := jwt.AudienceValidator(aud)	
	
	validatePayload := jwt.ValidatePayload(&pl, nbfValidator, expValidator, issValidator, audValidator)

	//Validate nbf, exp, iss and aud
	if _, err := jwt.Verify([]byte(jwtToken), hs, &pl, validatePayload);
	err != nil {
		log.Println("Error validating JWT: ", err.Error())
		writer.WriteHeader(401)
		return false
	}

	//If verification and validation passed, decode the token and verify the scope claim
	tokenParts := strings.Split(jwtToken, ".")
	encodedPayload := tokenParts[1]

	payload, err := base64.RawURLEncoding.DecodeString(encodedPayload)

	if err != nil {
		log.Println("Error decoding JWT body: ", err.Error())
		writer.WriteHeader(500)
		return false
	}

	type Payload struct {
		Scope string `json:"scope,omitempty"`
	}

	var rawPayload Payload

	err = json.Unmarshal([]byte(payload), &rawPayload)
	if err != nil {
		log.Println("Error parsing payload: ", err.Error())
		writer.WriteHeader(500)
	}

	//Check if configured scope values are missing
	if(!strings.Contains(rawPayload.Scope, os.Getenv("SCOPE"))){	
		writer.WriteHeader(401)
		writer.Write([]byte("Missing required scope(s)"))
		return false
	}

	//All checks passed
	return true
}