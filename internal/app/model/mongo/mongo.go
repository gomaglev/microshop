package mongo

import (
	"context"

	"github.com/gomaglev/microshop/v1/internal/pkg/config"
	"github.com/gomaglev/microshop/v1/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient create Mongo client
func NewClient() (*mongo.Client, func(), error) {

	cfg := config.C

	// &imongo.Config{
	// 	URI:      cfg.URI,
	// 	Database: cfg.Database,
	// 	Timeout:  time.Duration(cfg.Timeout) * time.Second,
	// }

	var (
		ctx    = context.Background()
		cancel context.CancelFunc
	)

	if t := cfg.Mongo.Timeout; t > 0 {
		ctx, cancel = context.WithTimeout(ctx, t)
		defer cancel()
	}

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {
		err := cli.Disconnect(context.Background())
		if err != nil {
			logger.Errorf(ctx, "Mongo disconnect error: %s", err.Error())
		}
	}

	err = cli.Ping(context.Background(), nil)
	if err != nil {
		return nil, cleanFunc, err
	}
	return cli, cleanFunc, nil
}

// CreateIndexes 创建索引
func CreateIndexes(ctx context.Context, cli *mongo.Client) error {
	return createIndexes(
		ctx,
		cli,
	)
}

type indexer interface {
	CreateIndexes(ctx context.Context, cli *mongo.Client) error
}

func createIndexes(ctx context.Context, cli *mongo.Client, indexes ...indexer) error {
	for _, idx := range indexes {
		err := idx.CreateIndexes(ctx, cli)
		if err != nil {
			return err
		}
	}
	return nil
}
