package temporal

import (
	"context"
	"log"
	"time"

	"github.com/Ratchaphon1412/assistant-llm/configs"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
)

func AIWorkflow(ctx context.Context, cfg *configs.Config, chatID uint, redisChanel string, workflowID string, input string) error {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort:  cfg.TEMPORAL_HOST + ":" + cfg.TEMPORAL_PORT,
		Namespace: cfg.TEMPORAL_NAMESPACE,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
		return err
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: cfg.TEMPORAL_TASK_QUEUE,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 100,
			MaximumAttempts:    5,
		},
	}

	we, err := c.ExecuteWorkflow(ctx, workflowOptions, cfg.TEMPORAL_WORKFLOW_NAME, chatID, redisChanel, input) // Replace with your actual workflow function name and parameters
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
		return err
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	return nil
}
