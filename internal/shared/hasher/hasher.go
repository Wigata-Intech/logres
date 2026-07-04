package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Params are the argon2id cost parameters. Memory is in KiB.
type Params struct {
	Memory  uint32
	Time    uint32
	Threads uint8
	SaltLen uint32
	KeyLen  uint32
}

// DefaultParams: m=64 MiB, t=3, p=1 (RFC 9106 second recommended set, above the
// OWASP floor). Benchmark on target hardware and tune Memory/Time to ~250–500 ms.
var DefaultParams = Params{
	Memory:  64 * 1024,
	Time:    3,
	Threads: 1,
	SaltLen: 16,
	KeyLen:  32,
}

// defaultConcurrency bounds in-flight hashes so a login flood queues instead of
// exhausting RAM (each hash holds Params.Memory KiB). Operators tune via
// WithConcurrency once the RAM budget is known.
const defaultConcurrency = 4

// ErrInvalidHash is returned when a stored PHC string cannot be parsed.
var ErrInvalidHash = errors.New("hasher: invalid encoded hash")

// Hasher hashes and verifies passwords. Verify reports needsRehash when the
// stored parameters differ from the current policy, so callers can transparently
// upgrade the row on a successful login.
type Hasher interface {
	Hash(password string) (string, error)
	Verify(password, encoded string) (ok, needsRehash bool, err error)
}

type argon2Hasher struct {
	params Params
	sem    chan struct{}
}

func New(opts ...OptionFunc) Hasher {
	h := &argon2Hasher{
		params: DefaultParams,
		sem:    make(chan struct{}, defaultConcurrency),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func (h *argon2Hasher) Hash(password string) (string, error) {
	salt := make([]byte, h.params.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("hasher.Hash: salt: %w", err)
	}
	key := h.key(password, salt, h.params)
	return encode(h.params, salt, key), nil
}

func (h *argon2Hasher) Verify(password, encoded string) (ok, needsRehash bool, err error) {
	params, salt, want, err := decode(encoded)
	if err != nil {
		return false, false, err
	}
	got := h.key(password, salt, params)
	if subtle.ConstantTimeCompare(got, want) != 1 {
		return false, false, nil
	}
	return true, params != h.params, nil
}

// key runs argon2id under the concurrency semaphore.
func (h *argon2Hasher) key(password string, salt []byte, p Params) []byte {
	h.sem <- struct{}{}
	defer func() { <-h.sem }()
	return argon2.IDKey([]byte(password), salt, p.Time, p.Memory, p.Threads, p.KeyLen)
}

func encode(p Params, salt, key []byte) string {
	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Time, p.Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	)
}

// decode parses a PHC string of the form
// $argon2id$v=19$m=65536,t=3,p=1$<b64salt>$<b64hash>.
func decode(encoded string) (p Params, salt, key []byte, err error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[0] != "" || parts[1] != "argon2id" {
		return Params{}, nil, nil, ErrInvalidHash
	}

	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil || version != argon2.Version {
		return Params{}, nil, nil, ErrInvalidHash
	}

	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Time, &p.Threads); err != nil {
		return Params{}, nil, nil, ErrInvalidHash
	}
	// Sscanf ignores trailing input; re-encoding and comparing rejects a params
	// segment with garbage after the last field (e.g. "m=8,t=1,p=1;...").
	if parts[3] != fmt.Sprintf("m=%d,t=%d,p=%d", p.Memory, p.Time, p.Threads) {
		return Params{}, nil, nil, ErrInvalidHash
	}

	salt, err = base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return Params{}, nil, nil, ErrInvalidHash
	}
	key, err = base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return Params{}, nil, nil, ErrInvalidHash
	}

	p.SaltLen = uint32(len(salt))
	p.KeyLen = uint32(len(key))
	return p, salt, key, nil
}
