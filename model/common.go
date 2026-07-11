package model

type Version struct {
	Build   int `json:"build"`
	Major   int `json:"major"`
	Content int `json:"content"`
}

type JsonNode map[string]any

// Paging contains cursors for a paginated API response.
type Paging struct {
	Cursors PagingCursors `json:"cursors"`
}

// PagingCursors contains cursors before and after the current page.
type PagingCursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

// List represents the untyped list referenced by Fingerprint.files in the
// upstream Swagger document. The document does not define this schema.
type List []any

type Fingerprint struct {
	Sha     string `json:"sha"`
	Version string `json:"version"`
	Files   List   `json:"files"`
}

// Objectとして定義されているが、実際にはstringで返却される
// type JsonLocalizedName map[string]any
type JsonLocalizedName string

type Arena struct {
	Name     JsonLocalizedName `json:"name"`
	Id       int               `json:"id"`
	RawName  string            `json:"rawName"`
	IconUrls map[string]any    `json:"iconUrls"`
}

type GameMode struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type VerifyTokenRequest struct {
	Token string `json:"token,omitempty"`
}

type VerifyTokenResponse struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

type StringList []string

type String map[string]any

type IntegerList []int

type Integer map[string]any

type Float float64

type TrailEventList []TrailEvent

type TrailEvent struct {
	Title       JsonLocalizedName `json:"title"`
	Description JsonLocalizedName `json:"description"`
	EventTag    string            `json:"eventTag"`
}

type ClientError struct {
	Reason  string         `json:"reason"`
	Message string         `json:"message"`
	Type    string         `json:"type"`
	Detail  map[string]any `json:"detail"`
}
