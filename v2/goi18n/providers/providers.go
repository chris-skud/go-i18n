package providers

type Provider interface {
	Marshal(v interface{}, fileFormat string) ([]byte, error)
}

func GetProvider(p string) Provider {
	switch p {
	case "smartling":
		return &smartling{}
	}

	return nil
}
