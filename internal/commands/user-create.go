package commands

import (
	"errors"
	"fmt"
	_apiClient "github.com/openinfradev/tks-api/pkg/api-client"
	"github.com/openinfradev/tks-api/pkg/domain"
	"github.com/openinfradev/tks-client/internal/helper"
	"github.com/spf13/cobra"
)

func NewUserCreateCommand(globalOpts *GlobalOptions) *cobra.Command {
	var (
		accountId   string
		name        string
		email       string
		role        string
		department  string
		description string
		password    string
	)

	var command = &cobra.Command{
		Use:   "create",
		Short: "Create users.",
		Long: `Create users.
	
	Example:
	tks user create --account-id <ACCOUNT_ID> --name <NAME> --email <EMAIL> --role <ROLE> --department <DEPARTMENT> --description <DESCRIPTION> --password <PASSWORD>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if globalOpts.CurrentOrganizationId == "" {
				return errors.New("current organization is not set")
			}

			input := domain.CreateUserRequest{
				AccountId:   accountId,
				Name:        name,
				Email:       email,
				Role:        role,
				Department:  department,
				Description: description,
				Password:    password,
			}

			apiClient, err := _apiClient.NewWithToken(globalOpts.ServerAddr, globalOpts.AuthToken)
			helper.CheckError(err)

			url := fmt.Sprintf("organizations/%v/users", globalOpts.CurrentOrganizationId)
			body, err := apiClient.Post(url, input)
			if err != nil {
				return err
			}

			var out = domain.CreateUserResponse{}
			helper.Transcode(body, &out)

			return nil
		},
	}

	command.Flags().StringVar(&accountId, "account-id", "", "[required, unique] user accountId")
	command.Flags().StringVar(&name, "name", "", "[required] name")
	command.Flags().StringVar(&email, "email", "", "[required, unique] email")
	command.Flags().StringVar(&role, "role", "", "[required] role( one of admin, user)")
	command.Flags().StringVar(&department, "department", "", "[optional] department")
	command.Flags().StringVar(&description, "description", "", "[optional] description")
	command.Flags().StringVar(&password, "password", "", "[required] password")

	return command
}
