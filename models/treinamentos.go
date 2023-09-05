package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Treinamentos struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	Treinamento               string             `bson:"treinamento" json:"treinamento"`
	CargaHoraria              string             `bson:"carga_horaria" json:"carga_horaria"`
	DataTreinamento           time.Time          `bson:"data_treinamento" json:"data_treinamento"`
	CaracteristicaTreinamento string             `bson:"caracteristica_treinamento" json:"caracteristica_treinamento"`
	Funcionarios              string             `bson:"funcionarios" json:"funcionarios"`
	Tipo                      string             `bson:"tipo" json:"tipo"`
	Observacoes               string             `bson:"observacoes" json:"observacoes"`
	CreatedAt                 time.Time          `bson:"created_at" json:"createdAt"`
}

// CreateTreinamento insere um novo treinamento na coleção de Treinamentos.
func (t *Treinamentos) CreateTreinamento(db *mongo.Database) (primitive.ObjectID, error) {
	t.ID = primitive.NewObjectID()
	t.CreatedAt = time.Now()

	// Verifique se já existe um treinamento com o mesmo nome
	filter := bson.M{"treinamento": t.Treinamento}
	var existingTreinamento Treinamentos
	err := db.Collection("treinamentos").FindOne(context.Background(), filter).Decode(&existingTreinamento)
	if err == nil {
		// Já existe um treinamento com o mesmo nome, não crie um novo
		return existingTreinamento.ID, nil
	} else if err != mongo.ErrNoDocuments {
		// Ocorreu um erro diferente de "nenhum documento encontrado"
		return primitive.NilObjectID, err
	}

	// Não existe um treinamento com o mesmo nome, crie um novo
	_, err = db.Collection("treinamentos").InsertOne(context.Background(), t)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return t.ID, nil
}
