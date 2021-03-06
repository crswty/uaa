package matchers

type UaaConfig struct {
	Issuer      Issuer     `yaml:"issuer"`
	Encryption  Encryption `yaml:"encryption"`
	Login       Login      `yaml:"login"`
	LoginSecret string     `yaml:"LOGIN_SECRET"`
	Jwt         Jwt        `yaml:"jwt"`
	Database    Database   `yaml:"database"`
}

type Issuer struct {
	Uri string `yaml:"uri"`
}

type Encryption struct {
	ActiveKeyLabel string `yaml:"active_key_label"`
	EncryptionKeys []struct {
		Label      string `yaml:"label"`
		Passphrase string `yaml:"passphrase"`
	} `yaml:"encryption_keys"`
}

type JwtTokenPolicySigningKey struct {
	SigningKey string `yaml:"signingKey"`
}

type JwtTokenPolicy struct {
	ActiveKeyId string                              `yaml:"activeKeyId"`
	Keys        map[string]JwtTokenPolicySigningKey `yaml:"keys"`
}

type JwtToken struct {
	Policy JwtTokenPolicy `yaml:"policy"`
}

type Jwt struct {
	Token JwtToken `yaml:"token"`
}

type Login struct {
	ServiceProviderKey         string `yaml:"serviceProviderKey"`
	ServiceProviderKeyPassword string `yaml:"serviceProviderKeyPassword"`
	ServiceProviderCertificate string `yaml:"serviceProviderKeyCertificate"`
}

type Database struct {
	MaxActive          int    `yaml:"maxactive"`
	MaxIdle            int    `yaml:"maxidle"`
	MinIdle            int    `yaml:"minidle"`
	RemoveAbandoned    bool   `yaml:"removeabandoned"`
	LogAbandoned       bool   `yaml:"logabandoned"`
	AbandonedTimeout   int    `yaml:"abandonedtimeout"`
	EvictionIntervalMs int    `yaml:"evictionintervalms"`
	DriverClassName    string `yaml:"driverClassName"`
	Url                string `yaml:"url"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
}
