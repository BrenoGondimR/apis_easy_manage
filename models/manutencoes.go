package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Manutencoes struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	Local          string             `bson:"local" json:"local"`
	Tipo           string             `bson:"tipo" json:"tipo"`
	DataOcorrencia time.Time          `bson:"data_ocorrencia" json:"data_ocorrencia"`
	DataResolucao  time.Time          `bson:"data_resolucao" json:"data_resolucao"`
	Servico        string             `bson:"servico" json:"servico"`
	Descricao      string             `bson:"descricao" json:"descricao"`
	Status         string             `bson:"status" json:"status"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
}

// CreateManutencao insere uma nova manutenção na coleção de Manutencoes.
func (m *Manutencoes) CreateManutencao(db *mongo.Database) (primitive.ObjectID, error) {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	_, err := db.Collection("manutencoes").InsertOne(context.Background(), m)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return m.ID, nil
}

// FindAll consulta todas as manutenções na coleção "manutencoes".
func (m *Manutencoes) FindAll(db *mongo.Database) ([]Manutencoes, error) {
	var result []Manutencoes
	collection := db.Collection("manutencoes")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateManutencaoStatus(db *mongo.Database, id primitive.ObjectID, status string) error {
	piscinaCollection := db.Collection("manutencoes")

	// Verifica se a manutenção existe.
	existingManutencao := &Manutencoes{}
	err := piscinaCollection.FindOne(context.TODO(), bson.M{
		"_id": id,
	}).Decode(existingManutencao)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("manutenção não encontrada")
		}
		return err
	}

	// Obtém a data atual.
	currentTime := time.Now()

	// Atualiza o estado da manutenção e a data de resolução.
	_, err = piscinaCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"status":         status,
			"data_resolucao": currentTime,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
