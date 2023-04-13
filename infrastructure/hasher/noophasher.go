package hasher

// noop does not hash.
//
// Yes, it is no-op.
type noop struct{}

func NewNoop() *noop {
	return &noop{}
}

func (h *noop) Hash(text []byte) ([]byte, error) {
	return text, nil
}
