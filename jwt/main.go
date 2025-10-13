// JSON WEB TOKEN
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("secretKey")

//
func main() {
    token, err := GenerateToken("123", "faiyaz@g.com", "Faiyaz", "Ahmed", "admin", "refreshToken")
    if err != nil {
        panic(err)
    }
    fmt.Println("Token:", token)

    claims, err := VerifyToken(token)
    if err != nil {
        fmt.Println("Token invalid:", err)
    } else {
        fmt.Printf("Token valid! UserID: %s, Email: %s\n", claims.UserID, claims.Email)
    }
}


type CustomClaim struct{
	UserID string `json:"user_id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Role string `json:"role"`
	jwt.RegisteredClaims  
}

//GenerateToken func takes claim parameters and return token,err
func GenerateToken(userID,email,first_name,last_name,role string,typeOfToken string)(string ,error){
 //1.Generating Token 
   //check the type of token  (refresh or access token) and calculate expiry time
    var expiryTime time.Time; 
    if strings.ToLower(typeOfToken)=="refreshToken"{
       expiryTime= time.Now().Add(time.Hour * 24);
	}else{
		expiryTime=time.Now().Add(time.Hour *2);
	}
   //claim
   claim:= CustomClaim{
	UserID: userID,
	Email: email,
	FirstName :first_name,
	LastName :last_name,
	Role: role,
	RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiryTime),
		Issuer: "app",
	},
   }
   //generate token
   token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
 //2.Signing token with secret key
    signedToken,err:= token.SignedString(SecretKey);

  return signedToken,err
}

//VerifyToken func takes a jwt token and verifies it and return Claims and error(if any)
func VerifyToken(token string)(CustomClaim,error){

 //1.Parsing the token
  claims:=CustomClaim{}
   parsedToken,err:=jwt.ParseWithClaims(token,&claims,func(token *jwt.Token)(interface{},error){
        //verify the method
			_,ok:=token.Method.(*jwt.SigningMethodHMAC)
			if !ok{
        return nil,errors.New("unaccepted signing method")
			}
			return SecretKey,nil
	 })
	 if err!=nil{
		return CustomClaim{},err;
	 }
 //2.Verify the token 
  if !parsedToken.Valid{
		return CustomClaim{},errors.New("invalid token")
	}
 //3.Validates the claims
 if (claims.UserID=="" || claims.Email=="" || claims.FirstName =="" ||claims.LastName=="" || claims.Role == ""){
  return CustomClaim{},errors.New("invalid token:missing claims")
 }
 //4.Return
 return claims,nil;
}