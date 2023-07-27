package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Tratamento struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	NomePiscineiro string             `bson:"nome_piscineiro" json:"nome_piscineiro"`
	NomeEmpresa    string             `bson:"nome_empresa" json:"nome_empresa"`
	Cloro          float64            `bson:"cloro" json:"cloro"`
	Pha            float64            `bson:"pha" json:"pha"`
	Alcalinidade   float64            `bson:"alcalinidade" json:"alcalinidade"`
	Acidez         float64            `bson:"acidez" json:"acidez"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
}

// InsertTratamento insere um tratamento na collection de Piscina.
func (f *Tratamento) InsertTratamento(db *mongo.Database) error {
	f.ID = primitive.NewObjectID()
	piscinaCollection := db.Collection("piscina")

	// Verifica se já existe algum tratamento no dia correspondente.
	startOfDay := time.Date(f.CreatedAt.Year(), f.CreatedAt.Month(), f.CreatedAt.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := piscinaCollection.FindOne(context.TODO(), bson.M{
		"created_at": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}).Decode(&Tratamento{})
	if err != nil && err != mongo.ErrNoDocuments {
		// criar um arquivo de log, etc.
		return err
	}
	if err == nil {
		return errors.New("já houve um tratamento no dia de hoje")
	}

	// Criando o Funcionário.
	f.CreatedAt = time.Now()
	_, err = piscinaCollection.InsertOne(context.TODO(), f)
	if err != nil {
		return err
	}
	return nil
}
