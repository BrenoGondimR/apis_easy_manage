package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Funcionario struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Nome      string             `bson:"nome" json:"nome"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Cargo     string             `bson:"cargo" json:"cargo"`
	Genero    string             `bson:"genero" json:"genero"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

// InsertFuncionario insere um funcionário na collection de Funcionarios.
func (f *Funcionario) InsertFuncionario(db *mongo.Database) error {
	f.ID = primitive.NewObjectID()
	funcionariosCollection := db.Collection("funcionarios")

	// Verifica se o funcionário já existe.
	err := funcionariosCollection.FindOne(context.TODO(), bson.M{"email": f.Email}).Decode(&Funcionario{})
	if err != nil && err != mongo.ErrNoDocuments {
		// Aqui você pode lidar com a gravação do erro, como criar um arquivo de log, etc.
		return err
	}
	if err == nil {
		return errors.New("o funcionário já existe com este email")
	}

	// Criando o Funcionário.
	f.CreatedAt = time.Now()
	_, err = funcionariosCollection.InsertOne(context.TODO(), f)
	if err != nil {
		return err
	}
	return nil
}

// FindAll consulta todos os funcionários na coleção "funcionarios"
func (f *Funcionario) FindAll(db *mongo.Database) ([]Funcionario, error) {
	var result []Funcionario
	collection := db.Collection("funcionarios")

	//TODO Depois aplicar filtros
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
