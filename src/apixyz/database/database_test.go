package database

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitAllConnection(t *testing.T) {
	// Inisialisasi Viper dengan konfigurasi yang dibutuhkan
	viper.Set("database.server", "localhost")
	viper.Set("database.user", "root")
	viper.Set("database.password", "password")
	viper.Set("database.schema", "mydatabase")

	// Panggil fungsi yang akan diuji
	err := InitAllConnection()

	// Periksa apakah ada error saat inisialisasi koneksi
	if err != nil {
		t.Errorf("Failed to initialize all connections: %s", err.Error())
	}

	// Periksa apakah koneksi dbSKS sudah terbuka
	_, err = dbSKS.GetConnection()
	if err != nil {
		t.Errorf("Failed to get SKS connection: %s", err.Error())
	}

	// Cleanup setelah unit test selesai
	CloseAllConnection()
}

func TestGetConnectionSKS(t *testing.T) {
	// Inisialisasi Viper dengan konfigurasi yang dibutuhkan
	viper.Set("database.server", "localhost")
	viper.Set("database.user", "root")
	viper.Set("database.password", "password")
	viper.Set("database.schema", "mydatabase")

	// Panggil fungsi yang akan diuji
	_, err := GetConnectionSKS()

	// Periksa apakah ada error saat mendapatkan koneksi SKS
	if err != nil {
		t.Errorf("Failed to get SKS connection: %s", err.Error())
	}

	// Cleanup setelah unit test selesai
	CloseConnectionSKS()
}

func TestCloseConnectionSKS(t *testing.T) {
	// Inisialisasi Viper dengan konfigurasi yang dibutuhkan
	viper.Set("database.server", "localhost")
	viper.Set("database.user", "root")
	viper.Set("database.password", "password")
	viper.Set("database.schema", "mydatabase")

	// Inisialisasi koneksi SKS
	err := InitConnectionSKS()
	if err != nil {
		t.Errorf("Failed to initialize SKS connection: %s", err.Error())
	}

	// Panggil fungsi yang akan diuji
	CloseConnectionSKS()

	// Periksa apakah koneksi dbSKS sudah ditutup
	_, err = dbSKS.GetConnection()
	if err == nil {
		t.Error("Failed to close SKS connection")
	}
}
