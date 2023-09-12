package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Treinamentos struct {
	ID                        primitive.ObjectID `bson:"_id" json:"_id"`
	Treinamento               string             `bson:"treinamento" json:"treinamento"`
	CargaHoraria              string             `bson:"carga_horaria" json:"carga_horaria"`
	DataTreinamento           time.Time          `bson:"data_treinamento" json:"data_treinamento"`
	CaracteristicaTreinamento string             `bson:"caracteristica_treinamento" json:"caracteristica_treinamento"`
	Funcionarios              string             `bson:"funcionarios" json:"funcionarios"`
	Tipo                      string             `bson:"tipo" json:"tipo"`
	Observacoes               string             `bson:"observacoes" json:"observacoes"`
	Status                    string             `bson:"status" json:"status"`
	CreatedAt                 time.Time          `bson:"created_at" json:"createdAt"`
}

// CreateTreinamento insere um novo treinamento na coleção de Treinamentos.
func (t *Treinamentos) CreateTreinamento(db *mongo.Database) primitive.ObjectID {
	t.ID = primitive.NewObjectID()
	t.CreatedAt = time.Now()

	// Verifique se já existe um treinamento com o mesmo nome
	filter := bson.M{"treinamento": t.Treinamento}
	var existingTreinamento Treinamentos
	err := db.Collection("treinamentos").FindOne(context.Background(), filter).Decode(&existingTreinamento)
	if err == nil {
		// Já existe um treinamento com o mesmo nome, não crie um novo
		return existingTreinamento.ID
	} else if err != mongo.ErrNoDocuments {
		// Ocorreu um erro diferente de "nenhum documento encontrado"
		return primitive.NilObjectID
	}

	// Não existe um treinamento com o mesmo nome, crie um novo
	_, err = db.Collection("treinamentos").InsertOne(context.Background(), t)
	if err != nil {
		return primitive.NilObjectID
	}

	return t.ID
}

// FindAll consulta todos os fornecedores na coleção "fornecedores".
func (t *Treinamentos) FindAll(db *mongo.Database) ([]Treinamentos, error) {
	var result []Treinamentos
	collection := db.Collection("treinamentos")

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

func UpdateTreinamentoStatus(db *mongo.Database, id primitive.ObjectID, status string) error {
	trenamentoCollection := db.Collection("treinamentos")

	// Verifica se a manutenção existe.
	existingManutencao := &Manutencoes{}
	err := trenamentoCollection.FindOne(context.TODO(), bson.M{
		"_id": id,
	}).Decode(existingManutencao)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("manutenção não encontrada")
		}
		return err
	}

	// Atualiza o estado da manutenção e a data de resolução.
	_, err = trenamentoCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"status": status,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
