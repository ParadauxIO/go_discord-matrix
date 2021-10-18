package main

/**
 * /login API for getting an access token.
 */
func login(username string, password string) string {
	type Identifier struct {
		Type string `json:"type"`
		User string `json:"user"`
	}

	type LoginSchema struct {
		Type       string     `json:"type"`
		Identifier Identifier `json:"identifier"`
		Password   string     `json:"password"`
	}

	return ""
}
