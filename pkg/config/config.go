package config

import (
	"coinpe/pkg/constants"
	"coinpe/pkg/logger"
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// we want to use the current working directory as the config directory
func getPwd() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(path.Dir(path.Dir(b))))
}

func parseFlags() {
	flag.String("job", "", "Name of the job to be executed. On completion of the job the server exits.")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

// we want to use the current working directory as the config directory
func getRootDir() string {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		p, err := os.Getwd()
		if err != nil {
			logger.Fatal("error in getting pwd ", err)
		}
		return p
	}
	return filepath.Clean(strings.TrimSpace(string(path)))
}

type viperLookuper struct {
	viper *viper.Viper
}

func (l viperLookuper) Lookup(key string) (string, bool) {
	v := l.viper.Get(key)
	if v == nil {
		return "", false
	}
	return v.(string), true
}

// Load configuration
func (c *Config) Load(cfg interface{}) error {
	parseFlags()
	v := viper.New()
	v.SetConfigType("env")
	v.SetConfigName(".env")
	v.AddConfigPath(getRootDir())
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		logger.Warn(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Fatal("Error reading config file: ", err)
		}
	}

	envconfig.ProcessWith(context.Background(), &envconfig.Config{
		Target:   cfg,
		Lookuper: viperLookuper{viper: v},
	})

	return nil
}

func (cfg *Config) GetServerPublicURL(c *gin.Context) string {
	if cfg.Server.PublicURL != "" {
		return cfg.Server.PublicURL
	} else if cfg.Environment != constants.EnvLocal {
		return fmt.Sprintf("http://%s", c.Request.Host)
	} else {
		return fmt.Sprintf("https://%s", c.Request.Host)
	}
}

func (c *Config) ShouldMock() bool {
	return c.Debug || c.Environment != constants.EnvProduction
}
