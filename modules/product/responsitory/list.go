package responsitory

import (
	"Test/modules/product/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ListProductStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*mongo.Cursor, error)
}

type listProductResp struct {
	collection *mongo.Collection
}

func NewListProductResp(collection *mongo.Collection) *listProductResp {
	return &listProductResp{collection: collection}
}

func (resp *listProductResp) ListProduct(
	ctx context.Context,
) ([]*models.Product, error) {
	var products []*models.Product

	cursor, err := resp.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(context.TODO()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return products, nil
}
