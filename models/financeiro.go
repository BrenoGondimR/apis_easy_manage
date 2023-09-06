package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Financeiro struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Ganhos    float64            `bson:"ganhos" json:"ganhos"`
	Custos    float64            `bson:"custos" json:"custos"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

// CreateFinanceiro Insert.
func (m *Financeiro) CreateFinanceiro(db *mongo.Database) (primitive.ObjectID, error) {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	_, err := db.Collection("financeiro").InsertOne(context.Background(), m)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return m.ID, nil
}

// GetAllCustos busca todos os registros de custos no banco de dados e calcula a soma.
func GetAllCustos(db *mongo.Database) (float64, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"custos": bson.M{"$gt": 0}}, // Filtra apenas registros com custos maiores que 0.
		},
		{
			"$group": bson.M{
				"_id":   nil,
				"total": bson.M{"$sum": "$custos"},
			},
		},
	}

	var result []bson.M
	cursor, err := db.Collection("financeiro").Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &result); err != nil {
		return 0, err
	}

	if len(result) > 0 {
		totalCustos := result[0]["total"].(float64)
		return totalCustos, nil
	}

	return 0, nil // Retorna 0 se não houver custos no banco de dados.
}

// GetAllGanhos busca todos os registros de ganhos no banco de dados e calcula a soma.
func GetAllGanhos(db *mongo.Database) (float64, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"ganhos": bson.M{"$gt": 0}}, // Filtra apenas registros com ganhos maiores que 0.
		},
		{
			"$group": bson.M{
				"_id":   nil,
				"total": bson.M{"$sum": "$ganhos"},
			},
		},
	}

	var result []bson.M
	cursor, err := db.Collection("financeiro").Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &result); err != nil {
		return 0, err
	}

	if len(result) > 0 {
		totalGanhos := result[0]["total"].(float64)
		return totalGanhos, nil
	}

	return 0, nil // Retorna 0 se não houver ganhos no banco de dados.
}

// GetAllRenda calcula o total de ganhos menos o total de custos.
func GetAllRenda(db *mongo.Database) (float64, error) {
	// Obtenha o total de ganhos
	totalGanhos, err := GetAllGanhos(db)
	if err != nil {
		return 0, err
	}

	// Obtenha o total de custos
	totalCustos, err := GetAllCustos(db)
	if err != nil {
		return 0, err
	}

	// Calcule o total de renda (ganhos - custos)
	totalRenda := totalGanhos - totalCustos
	return totalRenda, nil
}
