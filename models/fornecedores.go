package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Fornecedor struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
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

func UpdateFornecedorByID(db *mongo.Database, id primitive.ObjectID, updatedFornecedor *Fornecedor) error {
	collection := db.Collection("fornecedores")

	// Crie um filtro com base no ID fornecido.
	filter := bson.M{"_id": id}

	// Crie uma atualização para definir os novos valores.
	update := bson.M{
		"$set": bson.M{
			"fornecedor": updatedFornecedor.Fornecedor,
			"cidade":     updatedFornecedor.Cidade,
			"contato":    updatedFornecedor.Contato,
			"nota":       updatedFornecedor.Nota,
		},
	}

	// Execute a atualização no banco de dados.
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("fornecedor não encontrada")
		}
		return err
	}

	return nil
}

func GetFornecedorByID(db *mongo.Database, id primitive.ObjectID) (*Fornecedor, error) {
	collection := db.Collection("fornecedores")

	// Crie um filtro para encontrar a manutenção com o ID especificado.
	filter := bson.M{"_id": id}

	// Execute a consulta no banco de dados.
	var fornecedor Fornecedor
	err := collection.FindOne(context.TODO(), filter).Decode(&fornecedor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("treinamento não encontrado")
		}
		return nil, err
	}

	return &fornecedor, nil
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
