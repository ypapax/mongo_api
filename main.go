package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"gopkg.in/yaml.v2"
	"os"
)

type config struct {
	ApiKeys []string `yaml:"apiKeys"`
	ListenPort int `yaml:"listenPort"`
	MongoDb struct {
		ConnectionString string `yaml:"connectionString"`
		TargetCollection string `yaml:"targetCollection"`
	} `yaml:"mongoDb"`
}

func main(){
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

		r := chi.NewRouter()
		r.Get("/auditlog", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
		log.Printf("listening on %d\n", conf.ListenPort)
		return http.ListenAndServe(fmt.Sprintf(":%d", conf.ListenPort), r)
	}(); err != nil {
		log.Printf("error: %+v\n", err)
		os.Exit(1)
	}

}

