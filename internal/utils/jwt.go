package utils

//引用 jwt-go 這個套件
import (
	"componentmod/internal/api/errorcode"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/utils/log"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var secret, jwtTokenTime, jwtRefTokenTime string

//jwt secret setting
var JwtConfig = []cli.Flag{
	&cli.StringFlag{
		Name:        "jwt-secret",
		Usage:       "jwt secret",
		Value:       "f9946c78-f48a-435d-acc4-4bf469ef2680",
		Destination: &secret,
		EnvVars:     []string{"jwt_secret"},
	},
	&cli.StringFlag{
		Name:        "jwt-token-time",
		Usage:       "jwt token time",
		Value:       "480",
		Destination: &jwtTokenTime,
		EnvVars:     []string{"jwt_token_time"},
	},
	&cli.StringFlag{
		Name:        "jwt-ref-token-time",
		Usage:       "jwt ref token time",
		Value:       "1440",
		Destination: &jwtRefTokenTime,
		EnvVars:     []string{"jwt_ref_token_time"},
	},
}

const (
	jwtToken     = "token"
	refreshToken = "refreshToken"
)

//
func GenJwt(userId int, name string) (string, error) {
	intJwtTokenTime, err := strconv.Atoi(jwtTokenTime)
	if err != nil {
		return "", err
	}
	return generateJwtToken(jwtToken, intJwtTokenTime, userId, name)
}

func GenRefJwt(userId int, name string) (string, error) {
	intJwtRefTokenTime, err := strconv.Atoi(jwtRefTokenTime)
	if err != nil {
		return "", err
	}
	return generateJwtToken(refreshToken, intJwtRefTokenTime, userId, name)
}

func generateJwtToken(tokenType string, timeMinute int, userId int, name string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = tokenType
	claims["id"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(int64(timeMinute))).Unix()
	claims["nbf"] = time.Now().Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)

	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func ValidateAndRefreshTokenCheck(token string) (*backstagedto.JwtUserInfoDTO, error) {
	check, jwtInfoDTO, tokenType := validateJwtToke(token)
	if check == false || jwtInfoDTO == nil || tokenType != refreshToken {
		return nil, CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.REFRESH_AUTHORIZED_ERROR)
	}
	return jwtInfoDTO, nil
}

func ValidateAndTokenCheck(token string) (*backstagedto.JwtUserInfoDTO, error) {
	check, jwtInfoDTO, tokenType := validateJwtToke(token)
	if check == false || jwtInfoDTO == nil || tokenType != jwtToken {
		return nil, CreateApiErr(errorcode.UNAUTHORIZED_ERROR_CODE, errorcode.UNAUTHORIZED_ERROR)
	}
	return jwtInfoDTO, nil
}

// This is the api to refresh tokens
// Most of the code is taken from the jwt-go package's sample codes
// https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func validateJwtToke(jwtToken string) (bool, *backstagedto.JwtUserInfoDTO, string) {

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.UNAUTHORIZED_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return false, nil, ""
	}

	if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, _ := strconv.Atoi(fmt.Sprint(Claims["id"]))
		name := fmt.Sprint(Claims["name"])
		tokeynType := fmt.Sprint(Claims["sub"])
		jwtInfo := &backstagedto.JwtUserInfoDTO{
			Id:   id,
			Name: name,
		}
		return true, jwtInfo, tokeynType //"token"	or	"refreshToken"
	}

	return false, nil, ""
}
