package repositories

import (
	"context"
	"errors"
	"log"
	"re-crm/dtos"
	"re-crm/entities"
	"re-crm/errs"
	"re-crm/utils"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepo(t *testing.T) {
	userRepo := setupTestDB()

	t.Run("nick exists and pass matches", func(t *testing.T) {
		_, err := userRepo.GetByNickPass(context.Background(), dtos.LoginDTO{
			Nick: "testNick",
			Pass: "testPass",
		})
		if err != nil {
			t.Fatal("want err = nil, got err != nil")
		}
	})
	t.Run("nick exists but pass doesnt match", func(t *testing.T) {
		_, err := userRepo.GetByNickPass(context.Background(), dtos.LoginDTO{
			Nick: "testNick",
			Pass: "wrongPass",
		})
		if err == nil || !errors.Is(err, errs.InvalidCreds) {
			t.Fatalf("want err = errs.invalidCreds, got %v", err)
		}
	})
	t.Run("nick doesnt exist", func(t *testing.T) {
		_, err := userRepo.GetByNickPass(context.Background(), dtos.LoginDTO{
			Nick: "wrongNick",
			Pass: "anyPass",
		})
		if err == nil || !errors.Is(err, errs.NotFound) {
			t.Fatalf("want err = errs.NotFound, got %v", err)
		}
	})

}

// makes db connection and creates user {TestNick, bc(TestPass), admin}
func setupTestDB() *UserRepoPg {
	test_dsn := "host=localhost user=postgres password=admin dbname=re-crm-test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(test_dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("😡 err connecting to test db: ", err)
	}
	sqlDb, _ := db.DB()
	if _, err := sqlDb.Query("DROP TABLE IF EXISTS users "); err != nil {
		log.Fatal("😡 err dropping table users", err)
	}
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatal("😡 err migrating user schema: ", err)
	}

	if _, err = sqlDb.Exec(`
    	INSERT INTO users (nick, pass, role) VALUES ($1, $2, $3)`, "testNick", utils.Bcryptify("testPass"), "admin"); err != nil {
		log.Fatalf("😡 failed to insert user: %v", err)
	}

	return NewUserRepoPg(db)
}
