package tripleosarch

// An OSArch is a pair of GOOS and GOARCH values indicating a platform.
type OSArch struct {
	GOOS string
  GOARCH string
}

func ParseOSArch(s string) (*OSArch, error) {
  parts := strings.Split(s, "/")
  if len(parts) != 2 {
    return nil, errors.New("not <os>/<arch> format")
  }
  return &OSArch{
    GOOS: parts[0],
    GOARCH: parts[1],
  }, nil
}

func (p OSArch) String() string {
	return p.GOOS + "/" + p.GOARCH
}
