package eastic

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v6"
)

// Config of Elasticsearch
type Config struct {
	Hosts   string
	Timeout time.Duration
}

// NewClient create Elasticsearch client
func NewClient(cfg *Config) (*elasticsearch.Client, func(), error) {
	var (
		ctx    = context.Background()
		cancel context.CancelFunc
	)

	if t := cfg.Timeout; t > 0 {
		ctx, cancel = context.WithTimeout(ctx, t)
		defer cancel()
	}

	conf := elasticsearch.Config{
		Addresses: strings.Split(cfg.Hosts, ","),
		Transport: &http.Transport{
			DialContext: (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	client, err := elasticsearch.NewClient(conf)
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	_, err = client.Ping()
	if err != nil {
		return nil, cleanFunc, err
	}
	return client, cleanFunc, nil
}

// CreateIndexes 创建索引
func CreateIndexes(ctx context.Context, client *elasticsearch.Client) error {
	return createIndexes(
		ctx,
		client,
	)
}

type indexer interface {
	CreateIndexes(ctx context.Context, client *elasticsearch.Client) error
}

func createIndexes(ctx context.Context, client *elasticsearch.Client, indexes ...indexer) error {
	for _, idx := range indexes {
		err := idx.CreateIndexes(ctx, client)
		if err != nil {
			return err
		}
	}
	return nil
}
