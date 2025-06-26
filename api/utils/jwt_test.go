package utils

import "testing"

func TestJWT(t *testing.T) {
	token, _ := GenerateJWT(1, "admin")
	claims, err := DecodeJWT(token)
	if err != nil {
		t.Fatal(err)
	}
	dt, _ := claims.GetExpirationTime()
	t.Log("all good, claims are: ", dt)
}
