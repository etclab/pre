package samba

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/etclab/pre"
)

func GetPublicParams(proxyId InstanceId) (pp *pre.PublicParams) {
	resp, err := http.Get(string(proxyId) + "/getPublicParams")
	if err != nil {
		log.Fatalf("Failed to fetch public parameters: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GetPublicParams returned non-OK status: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&pp); err != nil {
		log.Fatalf("Failed to decode public parameters: %v", err)
	}
	return pp
}

func RegisterPublicKey(proxyId, id InstanceId, pk pre.PublicKey) {
	pkMsg := PublicKeyMessage{
		InstanceId: id,
		PublicKey:  pk,
	}

	body, err := json.Marshal(pkMsg)
	if err != nil {
		log.Fatalf("Failed to marshal public key message: %v", err)
	}
	resp, err := http.Post(
		string(proxyId)+"/registerPublicKey",
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		log.Fatalf("Failed to register public key: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("RegisterPublicKey returned non-OK status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
}

func RequestPublicKey(proxyId InstanceId, functionId FunctionId) pre.PublicKey {
	req := PublicKeyRequest{FunctionId: functionId}
	reqBody, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("Failed to Marshal public key request: %v", err)
	}

	resp, err := http.Post(string(proxyId)+"/getPublicKey", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf("Failed to send public key request: %v", err)
	}

	defer resp.Body.Close()
	var pkMsg PublicKeyMessage

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("RequestPublicKey returned non-OK status: %d, body: %s", resp.StatusCode, body)
	}

	if err := json.Unmarshal(body, &pkMsg); err != nil {
		log.Fatalf("Failed to unmarshal public key response: %v", err)
	}
	return pkMsg.PublicKey
}

func SendMessage[T SambaMessage](m T, destId InstanceId) (response *http.Response, err error) {
	reqBody, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(string(destId)+"/message", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
