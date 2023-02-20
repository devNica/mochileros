package configurations

type Argon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func NewArgonConfg() Argon2Config {
	argonConfig := Argon2Config{
		Memory:      64 * 24,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return argonConfig
}
