package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type hasilPrediksi struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Tanggal    string `json:"tanggal"`
	Penyakit   string `json:"penyakit"`
	Status     string `json:"status"`
	Presentase int    `json:"presentase"`
}

type penyakit struct {
	Id       int    `json:"id"`
	Penyakit string `json:"penyakit"`
	DNA      string `json:"dna"`
}

func main() {
	// Connect to Postgres
	connStr := "user=cnvpeuhhrqssab dbname=dcfnvjg2a29rfj password=ff5dd6b6ee20aff980ce149bccd0a782ee6dbb13347aa5fc431b86d4cae60c5d host=ec2-3-217-251-77.compute-1.amazonaws.com port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/result", func(c echo.Context) error {
		//validate search query
		//parse search query
		//benerin format tanggal
		var hasil_prediksi []hasilPrediksi
		rows, err := db.Query("SELECT * FROM hasil_prediksi")
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			hasil := new(hasilPrediksi)
			err := rows.Scan(&hasil.Id, &hasil.Nama, &hasil.Tanggal, &hasil.Penyakit, &hasil.Status, &hasil.Presentase)
			if err != nil {
				return err
			}
			hasil_prediksi = append(hasil_prediksi, *hasil)
		}
		return c.JSON(http.StatusOK, hasil_prediksi)
	})
	// e.POST("/result", insertQuery)
	// e.GET("/test", testResult)
	// e.POST("/test", insertTest)
	// e.POST("/disease", insertDisease)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
