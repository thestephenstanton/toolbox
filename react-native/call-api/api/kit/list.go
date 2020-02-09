package kit

type List struct {
	UID     string   `json:"id"`
	Name    string   `json:"name"`
	TodoIDs []string `json:"todoIds"`
}
