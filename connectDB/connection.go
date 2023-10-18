package connectdb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, err
	}

	authToken, err := auth.BuildAuthToken(
		context.TODO(),
		"mydb.123456789012.us-east-1.rds.amazonaws.com:3306", // Database Endpoint (With Port)
		"us-east-1", // AWS Region
		"myDb",      // Database Account
		cfg.Credentials,
	)

	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("user=asad dbname=mydb password=%s host=mydb.123456789012.us-east-1.rds.amazonaws.com port=5432 sslmode=disable", authToken)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
