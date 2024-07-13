package database

import (
	"api/src/models"
	"api/src/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testing_database_is_working_as_expected(t *testing.T) {
	// On setup une base de donnée sqlite en mémoire dans le cadre du test
	db := tests.GetTestDBConnection()

	// Vérifier que la connexion à la base de données est établie
	assert.NotNil(t, db, "La base de données ne doit pas être nil")

	sqlDB, err := db.DB()

	assert.NoError(t, err, "Erreur lors de l'obtention de la connexion DB")

	err = sqlDB.Ping()

	assert.NoError(t, err, "Erreur lors de la connexion à la base de données")

	// Ici on vérifie que la table users a bien été créée en migrant le modèle User sur la base de données sqlite
	var count int64

	db.AutoMigrate(&models.User{})

	err = db.Model(&models.User{}).Count(&count).Error

	assert.NoError(t, err, "Erreur lors de la vérification de la table users")
}
