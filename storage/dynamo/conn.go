package dynamo

import (
	"context"
	"github.com/EdmundMartin/discrete/torrents"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Options struct {
	Region string
}

type DataStore struct {
	client *dynamodb.Client
}

func (d DataStore) ListInfo(ctx context.Context, infoHashes []string) ([]*torrents.TorrentInfo, error) {
	//TODO implement me
	panic("implement me")
}

func DialConnection(ctx context.Context, opts *Options) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		func(options *config.LoadOptions) error {
			options.Region = opts.Region
			return nil
		},
	)
	if err != nil {
		return nil, err
	}
	svc := dynamodb.NewFromConfig(cfg)
	return svc, nil
}

func NewDataStore(ctx context.Context, opts *Options) (*DataStore, error) {
	conn, err := DialConnection(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &DataStore{client: conn}, nil
}
