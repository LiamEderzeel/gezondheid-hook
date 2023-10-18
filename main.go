package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type plugin struct {
	config *Config
}

type Config struct {
	Method            string `json:"method"`
	Url               string `json:"url"`
	StatusCodeMinimum int    `json:"statusCodeMinimum"`
}

func NewConfig() Config {
	return Config{StatusCodeMinimum: 100, Method: "POST"}
}

type Ctx struct {
	Proto      string `json:"proto"`
	StatusCode int    `json:"status"`
}

func (m *plugin) SetConfig(config []byte) {
	var c Config

	if err := json.Unmarshal([]byte(config), &c); err != nil {
		log.Fatalf("Error deserializing data: %v", err)
	}

	m.config = &c
}

func (m *plugin) Run(getCotext func() []byte, next func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recoverd %v", r)
		}
	}()

	var ctx Ctx

	if m.config == nil {
		c := NewConfig()
		m.config = &c
	}

	next()

	if err := json.Unmarshal([]byte(getCotext()), &ctx); err != nil {
		log.Fatalf("Error deserializing data: %v", err)
	}

	if ctx.StatusCode > m.config.StatusCodeMinimum {

		client := http.Client{}

		req, err := http.NewRequest(m.config.Method, m.config.Url, nil)

		if err != nil {
			panic(fmt.Sprintf("Failed to create request for %#v\n", m.config.Url))
			// log.Panicf("Failed to create request for %#v\n", m.config.Url)
		}

		res, err := client.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Received error: %v\n", err.Error()))
		}

		fmt.Println(res.StatusCode)

	}
}

var Plugin plugin
