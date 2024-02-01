package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/defang-io/defang/src/pkg"
)

func TenantFromAccessToken(at string) (pkg.TenantID, error) {
	parts := strings.Split(at, ".")
	if len(parts) != 3 {
		return "", errors.New("not a JWT")
	}
	var claims struct {
		Sub string `json:"sub"`
	}
	bytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bytes, &claims)
	return pkg.TenantID(claims.Sub), err
}