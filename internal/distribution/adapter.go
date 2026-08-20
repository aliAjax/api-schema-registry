package distribution

type Signer interface {
	Sign([]byte) string
	Verify([]byte, string) bool
}
