package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

type fakeFlag struct {
	Key   string
	Value string
}
type fakeENVLoader struct {
	flags []fakeFlag
}

func (l *fakeENVLoader) Load(v viper.Viper) (*viper.Viper, error) {
	for idx := range l.flags {
		f := l.flags[idx]
		v.Set(f.Key, f.Value)
	}
	return &v, nil
}

func TestLoadConfig(t *testing.T) {
	fileLoader := NewFileLoader(".env.sample", "../..")
	errFileLoader := NewFileLoader(".env.err", "../..")
	type args struct {
		loaders []Loader
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Unable to load env",
			args: args{
				loaders: []Loader{
					errFileLoader,
					&fakeENVLoader{flags: []fakeFlag{}},
				},
			},
			want: Config{
				ServiceName: "",
			},
		},
		{
			name: "Load from env",
			args: args{
				loaders: []Loader{
					fileLoader,
					&fakeENVLoader{flags: []fakeFlag{}},
				},
			},
			want: Config{
				Port: "8100",
			},
		},
		{
			name: "Load from flag env",
			args: args{
				loaders: []Loader{
					fileLoader,
					&fakeENVLoader{flags: []fakeFlag{}},
					NewENVLoader(),
				},
			},
			want: Config{
				ServiceName:     "example-be",
				BaseURL:         "http://localhost:8100",
				Port:            "8100",
				Env:             "dev",
				AllowedOrigins:  "*",
				AccessTokenTTL:  600,
				RefreshTokenTTL: 7776000,
				DBHost:          "127.0.0.1",
				DBPort:          "5432",
				DBUser:          "user",
				DBName:          "dbname",
				DBPass:          "pass",
				DBSSLMode:       "disable",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadConfig(tt.args.loaders)
			if got.ServiceName != tt.want.ServiceName {
				t.Error("TestLoadConfig fail service name did not equal")
			}
		})
	}
}

func TestConfig_GetCORS(t *testing.T) {
	type fields struct {
		AllowedOrigins string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Get cors list",
			fields: fields{
				AllowedOrigins: "localhost:8020;",
			},
			want: []string{
				"localhost:8020",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				AllowedOrigins: tt.fields.AllowedOrigins,
			}
			if got := c.GetCORS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetCORS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultConfigLoaders(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Get default config with 2 item",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConfigLoaders(); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("DefaultConfigLoaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
