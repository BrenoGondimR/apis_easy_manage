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
	Type           string             `bson:"type" json:"type"`
	NomePiscineiro string             `bson:"nome_piscineiro" json:"nome_piscineiro"`
	NomeEmpresa    string             `bson:"nome_empresa" json:"nome_empresa"`
	Cloro          float64            `bson:"cloro" json:"cloro"`
	Pha            float64            `bson:"pha" json:"pha"`
	Alcalinidade   float64            `bson:"alcalinidade" json:"alcalinidade"`
	Acidez         float64            `bson:"acidez" json:"acidez"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
}

type Manutencao struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	Titulo         string             `bson:"titulo" json:"titulo"`
	Type           string             `bson:"type" json:"type"`
	Estado         string             `bson:"estado" json:"estado"`
	NomePiscineiro string             `bson:"nome_piscineiro" json:"nome_piscineiro"`
	NomeEmpresa    string             `bson:"nome_empresa" json:"nome_empresa"`
	Prioridade     string             `bson:"prioridade" json:"prioridade"`
	Descricao      string             `bson:"descricao" json:"descricao"`
	CreatedAt      time.Time          `bson:"created_at" json:"createdAt"`
}

// InsertTratamento insere um tratamento na collection de Piscina.
func (f *Tratamento) InsertTratamento(db *mongo.Database) error {
	f.ID = primitive.NewObjectID()
	piscinaCollection := db.Collection("piscina")

	// Verifica se já existe algum tratamento no dia correspondente com type igual a true.
	startOfDay := time.Date(f.CreatedAt.Year(), f.CreatedAt.Month(), f.CreatedAt.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	existingTratamento := &Tratamento{}
	err := piscinaCollection.FindOne(context.TODO(), bson.M{
		"created_at": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
		"type": "Tratamento",
	}).Decode(existingTratamento)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Criar um arquivo de log, etc.
			return err
		}
	} else {
		return errors.New("já houve um tratamento no dia de hoje com type igual a true")
	}

	// Criando o tratamento.
	f.CreatedAt = time.Now()
	_, err = piscinaCollection.InsertOne(context.TODO(), f)
	if err != nil {
		return err
	}
	return nil
}

// UpdateManutencao atualiza o estado de uma manutenção na collection de Piscina.
func UpdateManutencao(db *mongo.Database, id primitive.ObjectID, estado string) error {
	piscinaCollection := db.Collection("piscina")

	// Verifica se a manutenção existe.
	existingManutencao := &Manutencao{}
	err := piscinaCollection.FindOne(context.TODO(), bson.M{
		"_id":  id,
		"type": "Manutenção",
	}).Decode(existingManutencao)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("manutenção não encontrada")
		}
		return err
	}

	// Atualiza o estado da manutenção.
	_, err = piscinaCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{
		"$set": bson.M{"estado": estado},
	})
	if err != nil {
		return err
	}

	return nil
}

// InsertManutencao insere uma manutenção na collection de Piscina.
func (m *Manutencao) InsertManutencao(db *mongo.Database) error {
	m.ID = primitive.NewObjectID()
	piscinaCollection := db.Collection("piscina")

	// Verifica se já existe alguma manutenção com o mesmo título e type igual a false.
	existingManutencao := &Manutencao{}
	err := piscinaCollection.FindOne(context.TODO(), bson.M{
		"titulo": m.Titulo,
		"type":   "Manutenção",
	}).Decode(existingManutencao)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Criar um arquivo de log, etc.
			return err
		}
	} else {
		return errors.New("já existe uma manutenção com o mesmo título e type igual a false")
	}
	m.CreatedAt = time.Now()
	_, err = piscinaCollection.InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}
	return nil
}

// FindAll consulta todos os tratamentos na coleção "piscina".
func (f *Tratamento) FindAll(db *mongo.Database) ([]Tratamento, error) {
	var result []Tratamento
	collection := db.Collection("piscina")

	// Filtro para buscar apenas documentos com "type" igual a true (tratamento).
	filter := bson.M{"type": "Tratamento"}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// FindAllManutencao consulta todos os tratamentos na coleção "piscina".
func (m *Manutencao) FindAllManutencao(db *mongo.Database) ([]Manutencao, error) {
	var result []Manutencao
	collection := db.Collection("piscina")

	// Filtro para buscar apenas documentos com "type" igual a (Manutencao).
	filter := bson.M{"type": "Manutenção"}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}
