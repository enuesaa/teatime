package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewQueryProvider(ctx context.Context, db *mongo.Database) *QueryProvider {
	return &QueryProvider{
		ctx: ctx,
		db: db,
	}
} 

type QueryProvider struct {
	ctx context.Context
	db  *mongo.Database
}

func (p *QueryProvider) Logs() *LogQuery {
	query := &LogQuery{
		Query{
			collectionName: "logs",
			db:             p.db,
			sc:             p.ctx,
		},
	}
	return query
}

func (p *QueryProvider) Teapods() *TeapodQuery {
	query := &TeapodQuery{
		Query{
			collectionName: "teapods",
			db:             p.db,
			sc:             p.ctx,
		},
	}
	return query
}

func (p *QueryProvider) Teas(teapod string, teabox string) *TeaQuery {
	collectionName := fmt.Sprintf("%s-%s", teapod, teabox)
	query := &TeaQuery{
		Query{
			collectionName: collectionName,
			db:             p.db,
			sc:             p.ctx,
		},
	}
	return query
}
