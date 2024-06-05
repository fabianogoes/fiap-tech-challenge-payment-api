package repository

import (
	"context"
	"github.com/fabianogoes/fiap-payment/domain/entities"
	"github.com/fabianogoes/fiap-payment/frameworks/repository/dbo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database) *PaymentRepository {
	return &PaymentRepository{db, db.Collection("payments")}
}

func (p *PaymentRepository) GetPaymentById(id string) (*entities.Payment, error) {
	var payment dbo.Payment

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	err = p.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return payment.ToEntity(), nil
}

func (p *PaymentRepository) CreatePayment(payment *entities.Payment) (*entities.Payment, error) {
	paymentCreate := dbo.ToPaymentDBO(payment)

	res, err := p.collection.InsertOne(context.Background(), paymentCreate)
	if err != nil {
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	paymentResponse, err := p.GetPaymentById(id.Hex())
	if err != nil {
		return nil, err
	}

	return paymentResponse, nil
}
