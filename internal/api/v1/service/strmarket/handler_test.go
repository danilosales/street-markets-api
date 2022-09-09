package strmarket

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/danilosales/api-street-markets/config"
	"github.com/danilosales/api-street-markets/config/logger"
	"github.com/danilosales/api-street-markets/internal/database"
	"github.com/danilosales/api-street-markets/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var (
	user     = "postgres"
	password = "postgres"
	db       = "markets_test"
	port     = "5433"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", db)
	os.Setenv("DB_PORT", port)

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_USER=" + user,
			"POSTGRES_DB=" + db,
			"listen_addresses = '*'",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	databaseUrl := fmt.Sprintf(dsn, user, password, port, db)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(60) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		db, err := sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	//Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

var ID int

func setup() (*gin.Engine, *StretMarketHandler) {
	appConf := config.AppConfig()
	logger := logger.New("info")
	database.ConnectDatabase(&appConf.Db, logger)
	database.DB.AutoMigrate(&model.StreetMarket{})

	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	handler := New(logger)
	return routes, handler
}

func getStreetMarketMock() model.StreetMarket {
	return model.StreetMarket{
		Long:       "-46550164",
		Lat:        "-23558733",
		Setcens:    "355030885000091",
		Areap:      "3550308005040",
		Coddist:    87,
		Distrito:   "VILA FORMOSA",
		Codsubpref: 26,
		Subprefe:   "ARICANDUVA-FORMOSA-CARRAO",
		Regiao5:    "Leste",
		Regiao8:    "Leste 1",
		NomeFeira:  "Test Street Market",
		Registro:   "4041-2",
		Logradouro: "RUA JOAO PADRE CARLOS",
		Numero:     "S/N",
		Bairro:     "VL FORMOSA",
		Referencia: "TV RUA PRETORIA",
	}
}

func createStreetMarketMock() {
	s := getStreetMarketMock()
	database.DB.Create(&s)
	ID = int(s.Id)
}

func removeStreetMarketMock() {
	database.DB.Delete(model.StreetMarket{}, ID)
}

func TestGetStreetMarket(t *testing.T) {
	r, h := setup()
	createStreetMarketMock()
	defer removeStreetMarketMock()

	r.GET("/api/v1/street-markets/:code", h.GetStreetMarket)

	req, _ := http.NewRequest("GET", "/api/v1/street-markets/4041-2", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var marketMock model.StreetMarketDto
	json.Unmarshal(rr.Body.Bytes(), &marketMock)

	assert.Equal(t, "Test Street Market", marketMock.NomeFeira)

}

func TestGetStreetMarketReturnsNotFound(t *testing.T) {
	r, h := setup()

	r.GET("/api/v1/street-markets/:code", h.GetStreetMarket)

	req, _ := http.NewRequest("GET", "/api/v1/street-markets/4041-0", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

}

func TestDeleteStreetMarket(t *testing.T) {
	r, h := setup()
	createStreetMarketMock()

	r.DELETE("/api/v1/street-markets/:code", h.DeleteStreetMarket)

	req, _ := http.NewRequest("DELETE", "/api/v1/street-markets/4041-2", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

}

func TestDeleteStreetMarketReturnsNotFound(t *testing.T) {
	r, h := setup()

	r.DELETE("/api/v1/street-markets/:code", h.DeleteStreetMarket)

	req, _ := http.NewRequest("DELETE", "/api/v1/street-markets/4041-2", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestCreateStreetMarket(t *testing.T) {
	r, h := setup()

	defer database.DB.Where(&model.StreetMarket{Registro: "4041-2"}).Delete(&model.StreetMarket{})

	r.POST("/api/v1/street-markets", h.CreateStreetMarket)

	m := getStreetMarketMock().ToDto()
	mJson, _ := json.Marshal(m)
	req, _ := http.NewRequest("POST", "/api/v1/street-markets", bytes.NewBuffer(mJson))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestCreateStreetMarketReturnsBadRequest(t *testing.T) {
	r, h := setup()

	r.POST("/api/v1/street-markets", h.CreateStreetMarket)

	m := getStreetMarketMock().ToDto()
	m.NomeFeira = ""
	mJson, _ := json.Marshal(m)
	req, _ := http.NewRequest("POST", "/api/v1/street-markets", bytes.NewBuffer(mJson))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateStreetMarketReturnsNoFound(t *testing.T) {
	r, h := setup()

	r.PUT("/api/v1/street-markets/:code", h.UpdateStreetMarket)

	req, _ := http.NewRequest("UPDATE", "/api/v1/street-markets/4041-2", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUpdateStreetMarketReturnsBadRequest(t *testing.T) {
	r, h := setup()
	createStreetMarketMock()
	defer removeStreetMarketMock()

	r.PUT("/api/v1/street-markets/:code", h.UpdateStreetMarket)

	m := getStreetMarketMock().ToDto()
	m.NomeFeira = ""
	mJson, _ := json.Marshal(m)

	req, _ := http.NewRequest("PUT", "/api/v1/street-markets/4041-2", bytes.NewBuffer(mJson))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateStreetMarket(t *testing.T) {
	r, h := setup()
	createStreetMarketMock()
	defer removeStreetMarketMock()

	r.PUT("/api/v1/street-markets/:code", h.UpdateStreetMarket)

	m := getStreetMarketMock().ToDto()
	m.NomeFeira = "Super market"
	mJson, _ := json.Marshal(m)

	req, _ := http.NewRequest("PUT", "/api/v1/street-markets/4041-2", bytes.NewBuffer(mJson))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestSearchStreetMarket(t *testing.T) {
	r, h := setup()
	createStreetMarketMock()
	defer removeStreetMarketMock()

	r.GET("/api/v1/street-markets", h.SearchStreetMarket)

	req, _ := http.NewRequest("GET", "/api/v1/street-markets?regiao5=Leste", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var marketMocks model.StreetMarketDtos
	json.Unmarshal(rr.Body.Bytes(), &marketMocks)

	assert.Equal(t, "Test Street Market", marketMocks[0].NomeFeira)

}
