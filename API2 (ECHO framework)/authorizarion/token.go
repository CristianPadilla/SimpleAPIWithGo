package authorizarion

import (
	"errors"
	"time"

	"github.com/CristianPadilla/APIs/clase7-ECHO/model"
	"github.com/dgrijalva/jwt-go"
)

//GenerateToken .
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			//tiempo de expiracion del token
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			// editor: el nombre de la compañia
			Issuer: "Company Name",
		},
	}
	//SigningMethodRS256 es solo uno de los metodos mas utilizados, pero hay varios
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	//firmar el token con certificado privado
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//ValidateToken .
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("token no valido")
	}
	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("no se pudo obtener el claim")
	}
	return *claim, nil

}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
