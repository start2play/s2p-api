package services

import "github.com/dgrijalva/jwt-go"

type jwtService struct{
	secretKey string,
	isr string,
}

func NewJWTService() *jwtService{
	return &jwtService{
		secretKey: "s2play-awl",
		isr: "s2p-api",
	}
}

type Claim struct{
	sum string `bson:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint)(string, error){
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: s.isr,
			IssuedAt: time.Now().Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(s.secretKey))

	if(err!=nil){
		return "",err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool{
	_, err := jwt.Parse(token, func(t, jwt.Token)(interface{},error){
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid{
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}


