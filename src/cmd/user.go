package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newUserCmd())
}

func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage User resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newUserShowCmd(),
	)

	return cmd
}

func newUserShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <UserID>",
		Short: "Show User",
		RunE:  runUserShowCmd,
	}

	return cmd
}

func runUserShowCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	if len(args) == 0 {
		return errors.New("UserID is required")
	}

	UserID := args[0]

	req := UserShowRequest{
		ID: UserID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.UserShow(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "UserShow was failed: req = %+v, res = %+v", req, res)
	}

	User := res.User
	fmt.Printf(
		"id: %d, name: %s, created_at: %v, updated_at: %v\n",
		User.ID, User.Name, User.CreatedAt, User.UpdatedAt,
	)

	return nil
}

func (client *Client) UserShow(ctx context.Context, apiRequest UserShowRequest) (*UserShowResponse, error) {
	subPath := fmt.Sprintf("/api/v1/user/%d", apiRequest.ID)
	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse UserShowResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
