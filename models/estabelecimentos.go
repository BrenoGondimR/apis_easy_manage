package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Estabelecimento struct {
	ID                 primitive.ObjectID `bson:"_id, omitempty"`
	EstabelecimentoId  primitive.ObjectID `bson:"estabelecimento_id, omitempty"`
	Codigo             int                `bson:"codigo" json:"codigo"`
	Apelido            string             `bson:"apelido" json:"apelido"`
	Cep                string             `bson:"cep" json:"cep"`
	Cidade             string             `bson:"cidade" json:"cidade"`
	CpfCnpj            string             `bson:"cpf_cnpj" json:"cpf_cnpj"`
	Estado             string             `bson:"estado" json:"estado"`
	Fantasia           string             `bson:"fantasia" json:"fantasia"`
	Instagram          string             `bson:"instagram" json:"instagram"`
	Logradouro         string             `bson:"logradouro" json:"logradouro"`
	Numero             int                `bson:"numero" json:"numero"`
	RazaoSocial        string             `bson:"razao_social" json:"razao_social"`
	Site               string             `bson:"site" json:"site"`
	Telefone           string             `bson:"telefone" json:"telefone"`
	Financeiro         []Financeiro       `bson:"financeiro" json:"financeiro,omitempty"`
	Fornecedores       []Fornecedor       `bson:"fornecedores" json:"fornecedores,omitempty"`
	Funcionarios       []Funcionario      `bson:"funcionarios" json:"funcionarios,omitempty"`
	Manutencoes        []Manutencoes      `bson:"manutencoes" json:"manutencoes,omitempty"`
	PiscinaTratamentos []Tratamento       `json:"piscina_tratamentos" bson:"piscina_tratamentos"`
	PiscinaManutencoes []Manutencao       `json:"piscina_manutencoes" bson:"piscina_manutencoes"`
	CreatedAt          time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at" json:"updated_at"`
}
