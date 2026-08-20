package distribution

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type HMACSigner struct{ Secret []byte }

func (s HMACSigner) Sign(b []byte) string {
	h := hmac.New(sha256.New, s.Secret)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
func (s HMACSigner) Verify(b []byte, sig string) bool {
	got := s.Sign(b)
	return hmac.Equal([]byte(got), []byte(sig))
}
