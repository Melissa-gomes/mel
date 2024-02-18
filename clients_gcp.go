package mel

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
)

type PubsubConfig struct {
	ProjectId      string
	TopicId        string
	SubscriptionId string
}

func Pubsub(ctx context.Context, conf PubsubConfig) (*pubsub.Client, error) {
	if conf.ProjectId == "" {
		return nil, errors.New("conf.ProjectId is required")
	}

	c, err := pubsub.NewClient(ctx, conf.ProjectId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error to init pubsub client: %s", err.Error()))
	}

	return c, nil
}

func PubsubTopic(ctx context.Context, conf PubsubConfig) (*pubsub.Topic, error) {
	if conf.ProjectId == "" || conf.TopicId == "" {
		if conf.ProjectId == "" {
			return nil, errors.New("conf.ProjectId is required")
		} else {
			return nil, errors.New("conf.TopicId is required")
		}
	}

	c, err := pubsub.NewClient(ctx, conf.ProjectId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error to init pubsub client: %s", err.Error()))
	}

	t := c.Topic(conf.TopicId)
	return t, nil
}

func PubsubSubscription(ctx context.Context, conf PubsubConfig) (*pubsub.Subscription, error) {
	if conf.ProjectId == "" || conf.SubscriptionId == "" {
		if conf.ProjectId == "" {
			return nil, errors.New("conf.ProjectId is required")
		} else {
			return nil, errors.New("conf.SubscriptionId is required")
		}
	}

	c, err := pubsub.NewClient(ctx, conf.ProjectId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error to init pubsub client: %s", err.Error()))
	}

	s := c.Subscription(conf.SubscriptionId)
	return s, nil
}

func BigQueryClient(ctx context.Context, projectId string) (*bigquery.Client, error) {
	bqC, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error to init bigquery client: %s", err.Error()))
	}

	return bqC, nil
}
