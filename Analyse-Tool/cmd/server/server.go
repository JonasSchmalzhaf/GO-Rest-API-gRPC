package main

import (
	"DBs-Micro/gRPC"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type metrics struct {
	availableDBsCount *prometheus.GaugeVec
}

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	recordMetrics(m)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	if err := http.ListenAndServe(":2112", nil); err != nil {
		log.Fatal(err)
	}
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		availableDBsCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "Available DBs Count",
			Help: "Current number of available DBs.",
		},
			[]string{"names"}),
	}
	reg.MustRegister(m.availableDBsCount)
	return m
}

func recordMetrics(m *metrics) {
	go func() {
		for {
			availableDBs := getAvailableDBs()
			fmt.Println(time.Now().String() + ": Aktuell " + strconv.Itoa(len(availableDBs)) + " DBs {" + strings.Join(availableDBs, ", ") + "}")
			m.availableDBsCount.With(prometheus.Labels{"names": "{" + strings.Join(availableDBs, ", ") + "}"}).Set(float64(len(availableDBs)))
			time.Sleep(2 * time.Second)
		}
	}()
}

func getAvailableDBs() []string {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	defer conn.Close()

	client := gRPC.NewDatabaseServiceClient(conn)

	response, err := client.GetMultipleDBs(context.Background(), &gRPC.GetRequest{})
	if err != nil {
		log.Fatalf("failed to create gRPC connection: %v", err)
	}

	return response.Names
}
