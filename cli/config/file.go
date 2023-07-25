package config

type Config struct {
	Identity Identity `json:"identity"`
}

type Identity struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`

	IdentityToken string `json:"identity_token,omitempty"`
}
