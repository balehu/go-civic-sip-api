package go_civic_sip_api

import (
	"./config"
	"./model"
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

func (cs *CivicSip) getPrivateKeyFromHex() {}

func createCivicExt(body string) string {
	return ""
}

func (cs *CivicSip) verifyAndDecrypt() (model.UserData, error) {
	return model.UserData{}, nil
}

func (cs *CivicSip) verify(data string) (string, error) {
	return "", nil
}

func (cs *CivicSip) getPublicKeyFromHexString() {

}

func (cs *CivicSip) decrypt(encodedData string) (string, error) {
	return "", nil
}

// const BASE_URL = `https://api.civic.com/sip`
// const AUTH_PATH = `scopeRequest/authCode`
