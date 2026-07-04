package otp

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Generate returns a zero-padded numeric code of the given length from a CSPRNG.
// rand.Int draws uniformly over [0, 10^digits), so there is no modulo bias.
func Generate(digits int) (string, error) {
	if digits <= 0 {
		return "", fmt.Errorf("otp.Generate: digits must be positive, got %d", digits)
	}
	upper := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil)
	n, err := rand.Int(rand.Reader, upper)
	if err != nil {
		return "", fmt.Errorf("otp.Generate: %w", err)
	}
	return fmt.Sprintf("%0*d", digits, n), nil
}

// Hash is the at-rest representation of a code: hex-encoded SHA-256. A fast hash
// is acceptable here because the low-entropy code is protected by the service's
// attempt throttle and short expiry, not by hash cost.
func Hash(code string) string {
	sum := sha256.Sum256([]byte(code))
	return hex.EncodeToString(sum[:])
}
