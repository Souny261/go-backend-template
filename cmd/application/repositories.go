package application

import (
	"backend/internal/adapters/secondary/mailer"
	"backend/internal/adapters/secondary/minio"
	"backend/internal/adapters/secondary/mysql"
	"backend/internal/adapters/secondary/redis"
	"backend/internal/config"
	"log"
)

// Repositories holds all data storage interfaces
type AppRepositories struct {
	// System repositories
	mysql  *mysql.MySQLRepository
	redis  *redis.RedisRepository
	minio  *minio.MinIORepository
	mailer *mailer.MailerRepository
	// Domain repositories
	user *mysql.UserRepository
}

// setupRepositories initializes all data storage connections
func SetupRepositories(cfg *config.Config) *AppRepositories {
	// MySQL setup
	mysqlRepo, err := mysql.NewMySQLRepository(cfg.Database)
	if err != nil {
		log.Fatalf("❌ Failed to initialize MySQL repository: %v", err)
	}

	// Redis setup
	redisRepo, err := redis.NewRedisRepository(cfg.Redis)
	if err != nil {
		log.Fatalf("❌ Failed to initialize Redis repository: %v", err)
	}

	// MinIO setup
	minioRepo, err := minio.NewMinIORepository(cfg.Minio)
	if err != nil {
		log.Fatalf("❌ Failed to initialize MinIO repository: %v", err)
	}
	// Mailer setup
	mailerRepo := mailer.NewMailerRepository(cfg.Mailer)

	// Domain repositories
	userRepo := mysql.NewUserRepository(mysqlRepo.GetDB())

	return &AppRepositories{
		mysql:  mysqlRepo,
		redis:  redisRepo,
		minio:  minioRepo,
		user:   userRepo,
		mailer: mailerRepo,
	}

}

// closeRepositories gracefully closes all connections
func CloseRepositories(repos *AppRepositories) {
	if repos.mysql != nil {
		repos.mysql.Close()
	}
	if repos.redis != nil {
		repos.redis.Close()
	}
}
