package cli

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/defang-io/defang/src/pkg/cli/client"
)

func Logout(ctx context.Context, client client.Client) error {
	Debug(" - Logging out")
	err := client.RevokeToken(ctx)
	// Ignore unauthenticated errors, since we're logging out anyway
	if e, ok := err.(*connect.Error); !ok || e.Code() != connect.CodeUnauthenticated {
		return err
	}
	// TODO: remove the cached token file
	// tokenFile := getTokenFile(fabric)
	// if err := os.Remove(tokenFile); err != nil {
	// 	return err
	// }
	return nil
}
