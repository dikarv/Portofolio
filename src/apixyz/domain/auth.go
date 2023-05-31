package domain

type (
	TokenClaims struct {
		UserID       string
		TokenID      string
		IdentityUser string
		Phoneno      string
		Alg          string
		Iat          float64
		Exp          float64
	}
)
