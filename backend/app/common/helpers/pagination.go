package helpers

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

func CreateCursor(id string, createdAt time.Time, pointsNext bool) Map {
	return Map{
		"id":          id,
		"created_at":  createdAt,
		"points_next": pointsNext,
	}
}

func EncodeCursor(cursor Map) string {
	if len(cursor) == 0 {
		return ""
	}
	serializedCursor, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	encodedCursor := base64.StdEncoding.EncodeToString(serializedCursor)
	return encodedCursor
}
func DecodeCursor(cursor string) (Map, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, err
	}

	var cur Map
	if err := json.Unmarshal(decodedCursor, &cur); err != nil {
		return nil, err
	}
	return cur, nil
}
