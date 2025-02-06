package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString gera uma string aleatória com o tamanho especificado.
func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < number; i++ {
		index, err := randomInt(k)
		if err != nil {
			panic("Failed to generate random number")
		}
		c := alphabet[index]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomEmail gera um email aleatório.
func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}

// randomInt retorna um número inteiro aleatório entre 0 e max usando crypto/rand.
func randomInt(max int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}
