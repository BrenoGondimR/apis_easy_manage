package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Fornecedor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Fornecedor string             `bson:"fornecedor" json:"fornecedor"`
	Cidade     string             `bson:"cidade" json:"cidade"`
	Contato    string             `bson:"contato" json:"contato"`
	Nota       float64            `bson:"nota" json:"nota"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
}

// InsertFornecedor insere um fornecedor na collection de Fornecedores.
func (f *Fornecedor) InsertFornecedor(db *mongo.Database) error {
	f.ID = primitive.NewObjectID()
	fornecedorsCollection := db.Collection("fornecedores")

	// Verifique se já existe um fornecedor com o mesmo nome
	filter := bson.M{"fornecedor": f.Fornecedor}
	var existingFornecedor Fornecedor
	err := fornecedorsCollection.FindOne(context.Background(), filter).Decode(&existingFornecedor)
	if err == nil {
		// Fornecedor já existe, não é necessário criar um novo
		return nil
	} else if err != mongo.ErrNoDocuments {
		// Ocorreu um erro diferente de "nenhum documento encontrado"
		return err
	}

	// Crie um novo fornecedor
	f.CreatedAt = time.Now()
	_, err = fornecedorsCollection.InsertOne(context.Background(), f)
	return err
}

// FindAll consulta todos os fornecedores na coleção "fornecedores".
func (f *Fornecedor) FindAll(db *mongo.Database) ([]Fornecedor, error) {
	var result []Fornecedor
	collection := db.Collection("fornecedores")

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
