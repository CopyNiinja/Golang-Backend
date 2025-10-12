package main

import "golang.org/x/crypto/bcrypt"

// HashPassword return hashed password and err
func HashPassword(password string)(string,error){

	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),14); //cost 14

	if err!=nil{
		return "",err
	}

	return string(hashedPassword),nil;
}
//return true if password matched otherwise false
func CheckPasswordHash(password, hash string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))

	return err==nil;
}

func main() {
	//bcrypt

	password:="super_secret";

	hashedPass,_:=HashPassword(password);

	//compare
	isValid:= CheckPasswordHash(password,hashedPass);

	if isValid {
		//password matched
	}
	
}