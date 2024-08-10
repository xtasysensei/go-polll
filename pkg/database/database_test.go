package database

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/xtasysensei/go-poll/internal/config"
)

func TestInit(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "initialized database variable",
			args: args{cfg: &config.Config{
				Postgres: config.PostgresConfig{
					Server:   "localhost",
					Port:     "5432",
					DBName:   "go_app",
					User:     "sensei",
					Password: "12345",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.cfg)
			if DB == nil {
				t.Errorf("Init() did not initialize the data variable")
			}

		})
	}
}

func TestConnect(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "connected successfully",
			args: args{cfg: &config.Config{
				Postgres: config.PostgresConfig{
					Server:   "localhost",
					Port:     "5432",
					DBName:   "go_app",
					User:     "sensei",
					Password: "12345",
				},
			}},
			wantErr: false,
		},
		{
			name: "connected successfully",
			args: args{cfg: &config.Config{
				Postgres: config.PostgresConfig{
					Server:   "localhost",
					Port:     "5432",
					DBName:   "test_db",
					User:     "test_user",
					Password: "test_password",
				},
			}},
			wantErr: true,
		},
		{
			name: "Empty config, connection failed",
			args: args{cfg: &config.Config{
				Postgres: config.PostgresConfig{
					Server:   "",
					Port:     "",
					DBName:   "",
					User:     "",
					Password: "",
				},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantErr {
				t.Error("Connect() =nil want non-nil *sql.DB")
				return
			}
			if !tt.wantErr {
				if err := got.Ping(); err != nil {
					t.Errorf("Connect() returned invalid *sql.DB: %v", err)
				}
			}
		})
	}
}
