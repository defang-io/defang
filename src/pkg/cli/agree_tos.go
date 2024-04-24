package cli

import (
	"context"
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/defang-io/defang/src/pkg/cli/client"
	"github.com/defang-io/defang/src/pkg/term"
)

var ErrTermsNotAgreed = errors.New("You must agree to the Defang terms of service to use this tool")

func InteractiveAgreeToS(ctx context.Context, client client.Client) error {
	term.Println(term.Nop, "Our latest terms of service can be found at https://defang.io/terms-service.html")

	var agreeToS bool
	err := survey.AskOne(&survey.Confirm{
		Message: "Do you agree to the Defang terms of service?",
	}, &agreeToS, nil)
	if err != nil {
		return err
	}

	if !agreeToS {
		return ErrTermsNotAgreed
	}

	return NonInteractiveAgreeToS(ctx, client)
}

func NonInteractiveAgreeToS(ctx context.Context, client client.Client) error {
	if DoDryRun {
		return ErrDryRun
	}

	if err := client.AgreeToS(ctx); err != nil {
		return err
	}
	term.Info(" * You have agreed to the Defang terms of service")
	return nil
}
