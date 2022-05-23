package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/SermoDigital/jose"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wilzyang/go-api/pkg"
)

// JWTAUTHURL - URL for oAUTH for Box
const JWTAUTHURL string = "https://api.box.com/oauth2/token"

// JWTGRANTTYPE - mandatory parameter for box oAuth
const JWTGRANTTYPE string = "urn:ietf:params:oauth:grant-type:jwt-bearer"

// oAuthResponse holds decoded JSON response from Box
type oAuthResponse struct {
	AccessToken  string   `json:"access_token"`
	Expires      int      `json:"expires_in"`
	RestrictedTo []string `json:"restricted_to"`
	TokenType    string   `json:"token_type"`
}

type AppUserResponse struct {
	ID string
}

type BoxConfig struct {
	PublicKeyID  string
	ClientID     string
	Sub          string
	ClientSecret string
}

func NewBoxAuth(bc BoxConfig) *BoxConfig {
	return &BoxConfig{
		PublicKeyID:  bc.PublicKeyID,
		ClientID:     bc.ClientID,
		Sub:          bc.Sub,
		ClientSecret: bc.ClientSecret,
	}
}

// CreateJWTAssertion - build up the JSON Web Token for oAuth
func (bc *BoxConfig) CreateJWTAssertion() (JWToken string, err error) {

	var msg string

	// Generate random bytes
	random, err := pkg.GenerateRandomBytes(32)
	if err != nil {
		msg = "Fail to generate random bytes"
		return msg, err
	}

	signingKey, err := ioutil.ReadFile("./config/box.pem")

	if err != nil {
		msg = "Unable to read signing key. Please ensure your private signing key is in the ./config/ directory"
		return msg, err
	}

	jti := jose.Base64Encode(random)

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	claims := make(jwt.MapClaims)
	token.Header["alg"] = "RS256"
	token.Header["typ"] = "JWT"
	token.Header["kid"] = bc.PublicKeyID
	claims["iss"] = bc.ClientID
	claims["aud"] = JWTAUTHURL
	claims["jti"] = jti
	claims["exp"] = time.Now().Add(time.Second * 45).Unix()
	claims["box_sub_type"] = "enterprise"
	claims["sub"] = bc.Sub

	token.Claims = claims

	key, err := jwt.ParseRSAPrivateKeyFromPEM(signingKey)

	if err != nil {
		msg = "error parsing RSA private key"
		return msg, err
	}

	// Sign the JWT
	JWToken, err = token.SignedString(key)

	if err != nil {
		msg = "Unable to sign token, please check that you have a signing key in ./keys/"
		return msg, err
	}

	return JWToken, err
}

// SendOAuthRequest - Sends a POST to authenticate against Box using JWT Assertion
func (bc *BoxConfig) SendOAuthRequest(JWToken string) (string, error) {

	var err error
	var msg string
	var decodedResponse *oAuthResponse

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("grant_type", JWTGRANTTYPE)
	_ = writer.WriteField("client_id", bc.ClientID)
	_ = writer.WriteField("client_secret", bc.ClientSecret)
	_ = writer.WriteField("assertion", JWToken)
	err = writer.Close()

	if err != nil {
		msg = "Writer error"
		return msg, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", JWTAUTHURL, payload)

	if err != nil {
		msg = "http error"
		return msg, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		msg = "Error submitting request to Box API"
		return msg, err
	}

	//
	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
	//

	err = json.NewDecoder(resp.Body).Decode(&decodedResponse)

	fmt.Println(err)

	if err != nil {
		msg = "Error decoding OAuthResponse"
	} else {
		// We only need the Access Token
		msg = decodedResponse.AccessToken
	}

	return msg, err
}

func (bc *BoxConfig) GenerateBoxJWT() (string, error) {
	JWTtoken, err := bc.CreateJWTAssertion()

	if err != nil {
		e := fmt.Sprintf("[GenerateBoxJWT] JWT Assertion : %v", err)
		return e, err
	}

	AuthToken, err := bc.SendOAuthRequest(JWTtoken)

	if err != nil {
		e := fmt.Sprintf("[GenerateBoxJWT] JWT Auth Token : %v", err)
		return e, err
	}

	return AuthToken, err
}
