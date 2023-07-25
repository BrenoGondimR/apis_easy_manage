package configure

import (
	"example/web-service-gin/db"
	"example/web-service-gin/utils"
)

// ConfigureDB realiza a configuração da conexão com o banco de dados MongoDB.
func ConfigureDB() {
	utils.DB = db.ConnectDB()
}
