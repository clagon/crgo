package crgo

type Version struct {
	Build   int `json:"build,omitempty"`
	Major   int `json:"major,omitempty"`
	Content int `json:"content,omitempty"`
}

type JsonNode map[string]any

// List represents the untyped list referenced by Fingerprint.files in the
// upstream Swagger document. The document does not define this schema.
type List []any

type Fingerprint struct {
	Sha     string `json:"sha,omitempty"`
	Version string `json:"version,omitempty"`
	Files   List   `json:"files,omitempty"`
}

type JsonLocalizedName map[string]any

type Arena struct {
	Name     JsonLocalizedName `json:"name,omitempty"`
	Id       int               `json:"id,omitempty"`
	RawName  string            `json:"rawName,omitempty"`
	IconUrls map[string]any    `json:"iconUrls,omitempty"`
}

type GameMode struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type VerifyTokenRequest struct {
	Token string `json:"token,omitempty"`
}

type VerifyTokenResponse struct {
	Tag    string `json:"tag,omitempty"`
	Token  string `json:"token,omitempty"`
	Status string `json:"status,omitempty"`
}

type StringList []string

type String map[string]any

type IntegerList []int

type Integer map[string]any

type Float map[string]any

type TrailEventList []TrailEvent

type TrailEvent struct {
	Title       JsonLocalizedName `json:"title,omitempty"`
	Description JsonLocalizedName `json:"description,omitempty"`
	EventTag    string            `json:"eventTag,omitempty"`
}

type ClientError struct {
	Reason  string         `json:"reason,omitempty"`
	Message string         `json:"message,omitempty"`
	Type    string         `json:"type,omitempty"`
	Detail  map[string]any `json:"detail,omitempty"`
}
