// Package kafka instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package kafka

import (
	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/describe"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/status"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/topics"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/use"
)

func NewKafkaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "Manage your Kafka instances",
		Long:  "Manage your Kafka instances",
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(),
		describe.NewDescribeCommand(),
		delete.NewDeleteCommand(),
		list.NewListCommand(),
		use.NewUseCommand(),
		status.NewStatusCommand(),
		topics.NewTopicsCommand(),
	)
	return cmd
}