package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// Load .env jika ada (opsional - environment variables bisa dari system/cloud)
	godotenv.Load(".env")

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan. Set via environment variable atau file .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database : %v", err)
	}

	DB = db
	log.Println("koneksi ke postgreSQL BERHASIL")

}
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum diinisialisasi")
	}
	return DB
}
