package unittest

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

// PostgresSuite is unit test for postgres
type PostgresSuite struct {
	MigrationPath string

	config PostgresConfig

	hasContainer bool
	createdDB    bool

	dsn string

	once sync.Once

	Image            string
	containerStarted int32
	container        testcontainers.Container
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (s *PostgresSuite) Setup() {
	s.once.Do(func() {
		if err := s.startPostgres(); err != nil {
			panic(err)
		}
	})
}

func (s *PostgresSuite) TearDown() {
	if err := s.stopPostgres(); err != nil {
		panic(err)
	}
}

func (s *PostgresSuite) Dsn() string {
	return s.dsn
}

func (s *PostgresSuite) startPostgres() error {
	s.loadConfig()

	if s.hasContainer {
		if err := s.startContainer(); err != nil {
			return err
		}
	}

	s.dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.config.User, s.config.Password, s.config.Host, s.config.Port, s.config.DBName)

	// create db
	if !s.createdDB {
		err := CreateDBIfNotExist(s.dsn)
		if err != nil {
			panic(err)
		}
		fmt.Printf("created database %s\n", s.config.DBName)

		s.createdDB = true
	}

	// migration
	if s.MigrationPath != "" {
		// TODO:
	}

	return nil
}

func (s *PostgresSuite) startContainer() error {
	ctx := context.Background()

	// start container
	if atomic.LoadInt32(&s.containerStarted) == 1 {
		return nil
	} else {
		defer atomic.StoreInt32(&s.containerStarted, 1)
	}

	fmt.Println("using container for testing")

	nPort, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return err
	}
	waitFor := wait.ForSQL(nPort, "postgres", func(port nat.Port) string {
		url := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable",
			s.config.User, s.config.Password, s.config.Host, port.Port())
		return url
	})
	waitFor.Timeout(30 * time.Second)

	req := testcontainers.ContainerRequest{
		Image:        s.Image,
		ExposedPorts: []string{nPort.Port()},
		Env: map[string]string{
			"POSTGRES_PASSWORD": s.config.Password,
		},
		WaitingFor: waitFor,
		AutoRemove: true,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return err
	}
	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return err
	}
	port := mappedPort.Port()
	s.container = container
	s.config.Host = host
	s.config.Port = port

	return nil
}

func (s *PostgresSuite) stopPostgres() error {
	if s.createdDB {
		err := DropDBIfExist(s.dsn)
		if err != nil {
			panic(err)
		}
	}

	if s.hasContainer && s.container != nil {
		// stop container
		s.container.Terminate(context.Background())
	}

	return nil
}

func (s *PostgresSuite) loadConfig() error {
	// generate a rand database for testing
	s.config.DBName = randString(10)

	// load config from environment variable
	if os.Getenv("POSTGRES_USER") != "" {
		s.config.User = os.Getenv("POSTGRES_USER")
		s.config.Password = os.Getenv("POSTGRES_PASSWORD")
		if h := os.Getenv("POSTGRES_HOST"); h != "" {
			s.config.Host = h
		} else {
			s.config.Host = "localhost"
		}
		if port := os.Getenv("POSTGRES_PORT"); port != "" {
			s.config.Port = port
		} else {
			s.config.Port = "5432"
		}

	} else {
		s.config.User = "postgres"
		s.config.Password = randString(10)
		s.config.Host = "localhost"
		s.config.Port = "5432"
		s.hasContainer = true
		if s.Image == "" {
			s.Image = "postgres:13-alpine"
		}
	}

	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyz"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
