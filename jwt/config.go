package jwt

//Key public keys storage structure
/*
	{"keys": [{...key1...}, {...key2...}]}
*/
type JWKPublicKeysStorage struct {
	Keys []JWKpublic `json:"keys"`
}

//JSON web key public key part
type JWKpublic struct {
	Kty string `json:"kty"` // Key type
	E   string `json:"e"`   // Key exponent
	Kid string `json:"kid"` // Key identifier
	N   string `json:"n"`   // Key modulus
	//Alg string `json:"alg"` // Encryption algorithm
	Use string `json:"use"` // Key usage encryption\signature
}

//JSON web key full key
//https://tools.ietf.org/html/rfc7518#section-6.3.1
type JWKfull struct {
	JWKpublic
	P   string   `json:"p"`
	Q   string   `json:"q"`
	D   string   `json:"d"`
	Qi  string   `json:"qi"`
	Dp  string   `json:"dp"`
	Dq  string   `json:"dq"`
	X5C []string `json:"x5c"`
}

//Configuration for one client key
type ClientKeysSettings struct {
	Sid string
	JWKpublic
}

//Configuration of clients public keys (temporary here, should be in DB and loaded fro COMS)
type ClientsKeysSettings []ClientKeysSettings

//Configuration with public key
type KeysSettings []JWKfull

//Configuration
type Configuration struct {
	ClientsKeys    ClientsKeysSettings
	Keys           KeysSettings
	Filename       string
	PublicKeys     JWKPublicKeysStorage
}
