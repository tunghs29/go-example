package component

import "go.mongodb.org/mongo-driver/mongo"

type AppContext interface {
	GetMainDBConnection() *mongo.Client
	SecretKey() string
}

type appCtx struct {
	client    *mongo.Client
	secretKey string
}

func NewAppContext(client *mongo.Client, secretKey string) *appCtx {
	return &appCtx{client: client, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *mongo.Client {
	return ctx.client
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
