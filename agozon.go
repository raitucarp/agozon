package agozon

import "os"

type Config struct {
	AssociateTag    string
	AWSAccessKeyId  string
	SecretAccessKey string
	Locale          string
}

// Create new Agozon request.
func NewRequest(c *Config) (r Request) {
	var (
		AssociateTag    string
		AWSAccessKeyId  string
		SecretAccessKey string
		Locale          string
	)

	// Get AssociateTag from environment variable.
	// If not set, then trying to get it via config.
	// It will Panic if both is empty
	if AssociateTag = os.Getenv("ASSOCIATE_TAG"); AssociateTag == "" {
		AssociateTag = c.AssociateTag
		if AssociateTag == "" {
			panic("AssociateTag is empty. Please set ASSOCIATE_TAG in your environment variable, or set it via config")
		}
	}

	// Get AssociateTag from environment variable.
	// If not set, then trying to get it via config.
	// It will Panic if both is empty.
	if AWSAccessKeyId = os.Getenv("ACCESS_KEY_ID"); AWSAccessKeyId == "" {
		AWSAccessKeyId = c.AWSAccessKeyId
		if AWSAccessKeyId == "" {
			panic("AWSAccessKeyId is empty. Please set ACCESS_KEY_ID in your environment variable, or set it via config")
		}
	}

	// Get SecretAccessKey from environment variable.
	// If not set, then trying to get it via config.
	// It will Panic if both is empty.
	if SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY"); SecretAccessKey == "" {
		SecretAccessKey = c.SecretAccessKey
		if SecretAccessKey == "" {
			panic(`SecretAccessKey is empty.
				Please set SECRET_ACCESS_KEY in your environment variable, or set it via config`)
		}
	}

	// Get locale configuration from Environment Variable,
	// If not found in env, then set locale from config.
	// But if locale in config is empty, set to default US
	if Locale = os.Getenv("Locale"); Locale == "" {
		Locale = c.Locale
		if Locale == "" {
			Locale = "US"
		}
	}

	// Set initial setup
	r.SetLocale(Locale)
	r.create(AssociateTag, AWSAccessKeyId, SecretAccessKey)

	return
}
