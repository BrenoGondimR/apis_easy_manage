package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Usuarios struct {
	ID                primitive.ObjectID `bson:"_id, omitempty"`
	Nome              string             `bson:"nome" json:"nome"`
	Email             string             `bson:"email" json:"email"`
	Telefone          string             `bson:"telefone" json:"telefone"`
	Password          string             `bson:"password_digest" json:"password"`
	Remember          string             `bson:"remember_digest" json:"remember"`
	Ativo             bool               `bson:"ativo" json:"ativo"`
	Role              string             `bson:"role" json:"role"`
	EstabelecimentoId primitive.ObjectID `bson:"estabelecimento_id, omitempty"`
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
}

// CreateUser insere um novo usuário no banco de dados e verifica se o usuário já existe.
func (u *Usuarios) CreateUser(db *mongo.Database) (primitive.ObjectID, error) {
	// Verificar se o usuário já existe com base no email.
	existingUser := Usuarios{}
	err := db.Collection("usuarios").FindOne(context.Background(), bson.M{"email": u.Email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// O usuário não existe, então podemos criar um novo.
			u.ID = primitive.NewObjectID()
			u.CreatedAt = time.Now()
			u.UpdatedAt = time.Now()

			_, err := db.Collection("usuarios").InsertOne(context.Background(), u)
			if err != nil {
				return primitive.NilObjectID, err
			}

			return u.ID, nil
		}
		// Ocorreu um erro ao verificar a existência do usuário.
		return primitive.NilObjectID, err
	}

	// O usuário já existe com o mesmo email.
	return primitive.NilObjectID, errors.New("O usuário já existe com o mesmo email")
}
