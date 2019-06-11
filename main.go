package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

const dbName = "test"

type config struct {
	ApiKeys    []string `yaml:"apiKeys"`
	ListenPort int      `yaml:"listenPort"`
	MongoDb    struct {
		ConnectionString string `yaml:"connectionString"`
		TargetCollection string `yaml:"targetCollection"`
	} `yaml:"mongoDb"`
}

type inputData struct {
	ApiKey string `json:"apiKey"`
}

func main() {
	if err := func() error {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		var configPath string
		flag.StringVar(&configPath, "config", "./config.yaml", "path to a config file")
		flag.Parse()
		b, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Printf("error: %+v\n", err)
			return err
		}
		var conf config
		if err := yaml.Unmarshal(b, &conf); err != nil {
			log.Printf("error: %+v\n", err)
			return err
		}
		log.Printf("connecting to mongo %+v\n", conf.MongoDb.ConnectionString)
		client, err := mongo.NewClient(options.Client().ApplyURI(conf.MongoDb.ConnectionString))
		if err != nil {
			log.Printf("error: %+v\n", err)
			return err
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err = client.Connect(ctx); err != nil {
			log.Printf("error: %+v\n", err)
			return err
		}
		r := chi.NewRouter()
		r.Post("/auditlog", func(w http.ResponseWriter, r *http.Request) {
			log.Println("auditlog")
			var inpData inputData
			var b []byte
			if err := func() error {
				b, err = ioutil.ReadAll(r.Body)
				if err != nil {
					log.Printf("error: %+v\n", err)
					return err
				}
				if err := json.Unmarshal(b, &inpData); err != nil {
					log.Printf("error: %+v\n", err)
					return err
				}
				return nil
			}(); err != nil {
				log.Printf("error: %+v\n", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			var foundApiKey bool
			for _, apiKey := range conf.ApiKeys {
				if apiKey == inpData.ApiKey {
					foundApiKey = true
					break
				}
			}
			if !foundApiKey {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			if err := func() error {
				var data map[string]interface{}
				if err := json.Unmarshal(b, &data); err != nil {
					log.Printf("error: %+v\n", err)
					return err
				}
				if _, ok := data["apiKey"]; ok {
					delete(data, "apiKey")
				}
				collection := client.Database(dbName).Collection(conf.MongoDb.TargetCollection)

				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				if _, err := collection.InsertOne(ctx, data); err != nil {
					log.Printf("error: %+v\n", err)
					return err
				}

				return nil
			}(); err != nil {
				log.Printf("error: %+v\n", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			w.Write([]byte("ok"))

		})
		log.Printf("listening on %d\n", conf.ListenPort)
		return http.ListenAndServe(fmt.Sprintf(":%d", conf.ListenPort), r)
	}(); err != nil {
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}

}
