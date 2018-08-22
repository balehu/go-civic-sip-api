package go_civic_sip_api

import (
	"./config"
	"./model"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"net/http"
)

type CivicSip struct {
	baseURL  string
	authPath string
	config   *config.CivicConfig
	// private static final ObjectMapper MAPPER = new ObjectMapper();
}

func NewCivicSip(config *config.CivicConfig) *CivicSip {
	return &CivicSip{baseURL: `https://api.civic.com/sip`,
		authPath: `scopeRequest/authCode`,
		config:   config}

}

func (cs *CivicSip) ExchangeToken(jwtToken string) (model.UserData, error) {
	return model.UserData{}, nil
}

func (cs *CivicSip) createHttpClient() *http.Client {
	return &http.Client{}
}

func (cs *CivicSip) makeAUthorizationHeader(body string) (string, error) {
	return "", nil
}

func (cs *CivicSip) createToken() (string, error) {
	// now := time.Now().UnixNano() / int64(time.Microsecond)
	// till := now + 3 * 60000
	return "", nil
}

func (cs *CivicSip) createJwtBuilder() {

}

func (cs *CivicSip) getPrivateKeyFromHex() (*ecdsa.PrivateKey, error) {

	b, err := hex.DecodeString(cs.config.PrivateKey)
	if err != nil {
		return nil, errors.New("invalid hex string")
	}
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = elliptic.P256()

	/*
		if 8*len(b) != priv.Params().BitSize {
			return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
		}
	*/
	priv.D = new(big.Int).SetBytes(b)

	/*
		// The priv.D must < N
		if priv.D.Cmp(secp256k1N) >= 0 {
			return nil, fmt.Errorf("invalid private key, >=N")
		}
	*/

	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(b)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}

func createCivicExt(body string) string {
	return ""
}

func (cs *CivicSip) verifyAndDecrypt() (model.UserData, error) {
	return model.UserData{}, nil
}

func (cs *CivicSip) verify(data string) (string, error) {
	return "", nil
}

func (cs *CivicSip) getPublicKeyFromHexString() (*ecdsa.PublicKey, error) {

	keyHex, err := hex.DecodeString(cs.config.PublicKey)
	if err != nil {
		return nil, err
	}
	x, y := elliptic.Unmarshal(elliptic.P256(), keyHex)
	//log.Printf("x : %v, y: %v", x, y)
	key := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	return key, nil
}

func (cs *CivicSip) decrypt(encodedData string) (string, error) {

	pkBytes, err := hex.DecodeString(cs.config.ApplicationSecret)
	if err != nil {
		return "", err
	}

	var block cipher.Block

	if block, err = aes.NewCipher(pkBytes); err != nil {
		return "", err
	}

	iv := []byte(encodedData[0:32])
	log.Printf("iv : %v", iv)
	messagePart := encodedData[32:]
	log.Printf("messagePart : %v", messagePart)
	var encodedPart []byte
	len, err := base64.StdEncoding.Decode(encodedPart, []byte(messagePart))
	if err != nil {
		return "", err
	}
	log.Printf("len : %v", len)

	cbc := cipher.NewCBCDecrypter(block, iv)

	var decodedMessage []byte
	cbc.CryptBlocks(decodedMessage, encodedPart)
	log.Printf("decoded message : %v", string(decodedMessage[:]))
	return string(decodedMessage[:]), nil
}

// const BASE_URL = `https://api.civic.com/sip`
// const AUTH_PATH = `scopeRequest/authCode`
