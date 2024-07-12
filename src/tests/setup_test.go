package tests

import (
	"api/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupTestDB(t *testing.T) {
	db := GetTestDBConnection()

	// Vérifier que la connexion à la base de données est établie
	assert.NotNil(t, db, "La base de données ne doit pas être nil")

	sqlDB, err := db.DB()

	assert.NoError(t, err, "Erreur lors de l'obtention de la connexion DB")

	err = sqlDB.Ping()

	assert.NoError(t, err, "Erreur lors de la connexion à la base de données")

	// Vérifier qu'une migration s'effectue correctement
	var count int64

	db.AutoMigrate(&models.User{})

	err = db.Model(&models.User{}).Count(&count).Error

	assert.NoError(t, err, "Erreur lors de la vérification de la table users")
}
