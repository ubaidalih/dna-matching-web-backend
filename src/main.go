package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/ubaidalih/Tubes3_13520061/algorithm"
)

type hasilPrediksi struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Tanggal    string `json:"tanggal"`
	Penyakit   string `json:"penyakit"`
	Status     string `json:"status"`
	Persentase int    `json:"persentase"`
}

type penyakit struct {
	Id       int    `json:"id"`
	Penyakit string `json:"penyakit"`
	DNA      string `json:"dna"`
}

type message struct {
	Message string `json:"message"`
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	e.GET("/", hello)
	e.POST("/result", func(c echo.Context) error {
		jsonBody := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
		if err != nil {
			log.Error("empty json body")
			return err
		}
		query := jsonBody["query"].(string)
		format := algorithm.ValidateQuery(query)
		if format == -1 {
			return c.JSON(http.StatusPartialContent, message{"Format query tidak valid"})
		}
		newQuery := algorithm.ParseQuery(query, format)
		tanggal := newQuery[0]
		penyakit := newQuery[1]
		//benerin format tanggal
		var hasil_prediksi []hasilPrediksi
		var rows *sql.Rows
		if tanggal == "" {
			rows, err = db.Query("SELECT * FROM hasil_prediksi WHERE penyakit = $1", penyakit)
		} else if penyakit == "" {
			rows, err = db.Query("SELECT * FROM hasil_prediksi WHERE tanggal = $1", tanggal)
		} else {
			rows, err = db.Query("SELECT * FROM hasil_prediksi WHERE tanggal = $1 AND penyakit = $2", tanggal, penyakit)
		}
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			hasil := new(hasilPrediksi)
			err := rows.Scan(&hasil.Id, &hasil.Nama, &hasil.Tanggal, &hasil.Penyakit, &hasil.Status, &hasil.Persentase)
			if err != nil {
				return err
			}
			hasil.Tanggal = hasil.Tanggal[:10]
			hasil_prediksi = append(hasil_prediksi, *hasil)
		}
		if hasil_prediksi == nil {
			return c.JSON(http.StatusPartialContent, message{"Tidak ada history yang cocok"})
		}
		return c.JSON(http.StatusOK, hasil_prediksi)
	})
	e.POST("/test", func(c echo.Context) error {
		jsonBody := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
		if err != nil {
			log.Error("empty json body")
			return err
		}
		nama := jsonBody["nama"].(string)
		dna := jsonBody["dna"].(string)
		penyakit := jsonBody["penyakit"].(string)
		algo := jsonBody["algo"].(string)
		var status string
		var persentase int
		curTime := time.Now()
		tanggal := curTime.Format("2006-01-02")

		if !algorithm.ValidateInput(dna) {
			return c.JSON(http.StatusPartialContent, message{"DNA tidak valid"})
		}

		rows, err := db.Query("SELECT rantai_dna FROM penyakit WHERE nama_penyakit = $1", penyakit)
		if err != nil {
			return err
		}
		defer rows.Close()

		var dna_penyakit string
		rows.Next()
		err = rows.Scan(&dna_penyakit)
		if err != nil {
			return c.JSON(http.StatusPartialContent, message{"Penyakit tidak ditemukan"})
		}

		if algo == "true" {
			if algorithm.KMP(dna, dna_penyakit) != -1 {
				status = "True"
				persentase = 100
			} else {
				persentase = algorithm.HammingDistance(dna, dna_penyakit)
				if persentase >= 80 {
					status = "True"
				} else {
					status = "False"
				}
			}
		} else {
			if algorithm.BoyerMoore(dna, dna_penyakit) != -1 {
				status = "True"
				persentase = 100
			} else {
				persentase = algorithm.HammingDistance(dna, dna_penyakit)
				if persentase >= 80 {
					status = "True"
				} else {
					status = "False"
				}
			}
		}
		db.Query("INSERT INTO hasil_prediksi (nama_pasien, tanggal, penyakit, status, persentase) VALUES ($1, $2, $3, $4, $5)", nama, tanggal, penyakit, status, persentase)

		return c.JSON(http.StatusOK, hasilPrediksi{0, nama, tanggal, penyakit, status, persentase})
	})
	e.POST("/disease", func(c echo.Context) error {
		jsonBody := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
		if err != nil {
			log.Error("empty json body")
			return err
		}
		penyakit := jsonBody["penyakit"].(string)
		dna := jsonBody["dna"].(string)

		if !algorithm.ValidateInput(dna) {
			return c.JSON(http.StatusOK, message{"DNA tidak valid"})
		}

		var nama string
		err = db.QueryRow("SELECT nama_penyakit FROM penyakit WHERE nama_penyakit = $1", penyakit).Scan(&nama)
		if err == sql.ErrNoRows {
			db.Query("INSERT INTO penyakit (nama_penyakit, rantai_dna) VALUES ($1, $2)", penyakit, dna)
			return c.JSON(http.StatusOK, message{"Penyakit berhasil ditambahkan"})
		} else if err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, message{"Penyakit sudah ada"})
		}
	})

	// Start server
	// port, _ := os.LookupEnv("PORT")
	// e.Logger.Fatal(e.Start(":" + port))
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
