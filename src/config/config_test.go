package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestEnvironment(t *testing.T) {
	// Inisialisasi Viper
	viper.Set("environment", "production")

	// Panggil fungsi yang akan diuji
	result := Environment()

	// Periksa apakah hasil sesuai dengan yang diharapkan
	expected := "production"
	if result != expected {
		t.Errorf("Expected environment to be %s, but got %s", expected, result)
	}
}

func TestHostname(t *testing.T) {
	// Inisialisasi Viper
	viper.Set("host", "example.com")

	// Panggil fungsi yang akan diuji
	result := Hostname()

	// Periksa apakah hasil sesuai dengan yang diharapkan
	expected := "example.com"
	if result != expected {
		t.Errorf("Expected hostname to be %s, but got %s", expected, result)
	}
}

func TestDetermineListenAddress(t *testing.T) {
	// Inisialisasi Viper
	viper.Set("port", "8080")

	// Panggil fungsi yang akan diuji
	result := DetermineListenAddress()

	// Periksa apakah hasil sesuai dengan yang diharapkan
	expected := ":8080"
	if result != expected {
		t.Errorf("Expected listen address to be %s, but got %s", expected, result)
	}
}

func TestRealm(t *testing.T) {
	// Inisialisasi Viper
	viper.Set("realm", "myrealm")

	// Panggil fungsi yang akan diuji
	result := Realm()

	// Periksa apakah hasil sesuai dengan yang diharapkan
	expected := "myrealm"
	if result != expected {
		t.Errorf("Expected realm to be %s, but got %s", expected, result)
	}
}

func TestSuperSecretPassword(t *testing.T) {
	// Inisialisasi Viper
	viper.Set("token.supersecretpassword", "encryptedpassword")

	// Panggil fungsi yang akan diuji
	result := SuperSecretPassword()

	// Periksa apakah hasil sesuai dengan yang diharapkan
	expected := "decryptedpassword"
	if result != expected {
		t.Errorf("Expected super secret password to be %s, but got %s", expected, result)
	}
}

func TestTimeout(t *testing.T) {
	// Inisialisasi Viper
	viper.Set("timeout", 10)

	// Panggil fungsi yang akan diuji
	result := Timeout()

	// Periksa apakah hasil sesuai dengan yang diharapkan
	expected := 10
	if result != expected {
		t.Errorf("Expected timeout to be %d, but got %d", expected, result)
	}
}
