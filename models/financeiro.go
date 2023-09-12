package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Financeiro struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Ganhos        float64            `bson:"ganhos" json:"ganhos"`
	Origem        string             `bson:"origem" json:"origem"`
	TipoTransacao string             `bson:"tipo_transacao" json:"tipo_transacao"`
	Custos        float64            `bson:"custos" json:"custos"`
	Status        string             `bson:"status" json:"status"`
	Data          time.Time          `bson:"data" json:"data"`
	CreatedAt     time.Time          `bson:"created_at" json:"createdAt"`
}

// TotaisMensais contém os totais de custos, ganhos e renda por mês.
type TotaisMensais struct {
	Custos []float64 `json:"custos"`
	Ganhos []float64 `json:"ganhos"`
	Renda  []float64 `json:"renda"`
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

// CalcularTotaisMensais calcula os totais mensais de custos, ganhos e renda.
func CalcularTotaisMensais(db *mongo.Database) (*TotaisMensais, error) {
	// Crie um mapa para armazenar os totais mensais para cada campo.
	totais := &TotaisMensais{
		Custos: make([]float64, 12), // Inicialize com 12 meses
		Ganhos: make([]float64, 12),
		Renda:  make([]float64, 12),
	}

	// Defina o período de tempo desejado, por exemplo, para o ano atual.
	agora := time.Now()
	anoAtual := agora.Year()

	// Crie um pipeline para filtrar registros relevantes (no ano atual).
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"data": bson.M{
					"$gte": time.Date(anoAtual, 1, 1, 0, 0, 0, 0, time.UTC),
					"$lte": time.Date(anoAtual, 12, 31, 23, 59, 59, 999999999, time.UTC),
				},
			},
		},
	}

	// Execute a agregação para obter os registros relevantes.
	cursor, err := db.Collection("financeiro").Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Percorra os registros e some os valores para cada mês.
	for cursor.Next(context.Background()) {
		var financeiro Financeiro
		if err := cursor.Decode(&financeiro); err != nil {
			return nil, err
		}

		mes := financeiro.Data.Month() // Obtém o mês
		// Atualize os totais mensais para cada campo.
		totais.Custos[mes-1] += financeiro.Custos // Subtraia 1 para obter o índice correto (janeiro é 1, fevereiro é 2, etc.)
		totais.Ganhos[mes-1] += financeiro.Ganhos
		totais.Renda[mes-1] = totais.Ganhos[mes-1] - totais.Custos[mes-1]
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return totais, nil
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

// FindAll consulta todas as manutenções na coleção "manutencoes".
func (m *Financeiro) FindAll(db *mongo.Database) ([]Financeiro, error) {
	var result []Financeiro
	collection := db.Collection("financeiro")

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
