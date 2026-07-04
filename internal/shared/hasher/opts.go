package hasher

type OptionFunc func(*argon2Hasher)

func WithParams(p Params) OptionFunc { return func(h *argon2Hasher) { h.params = p } }

// WithConcurrency caps in-flight hashes at n. A value of n <= 0 is ignored and
// the default cap is kept — it does not mean "unbounded".
func WithConcurrency(n int) OptionFunc {
	return func(h *argon2Hasher) {
		if n > 0 {
			h.sem = make(chan struct{}, n)
		}
	}
}
