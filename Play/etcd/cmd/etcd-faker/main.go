package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type valueType string

const (
	valueTypeText   valueType = "text"
	valueTypeJSON   valueType = "json"
	valueTypeBinary valueType = "binary"
)

func main() {
	var (
		endpointList   string
		totalKeys      int
		concurrency    int
		keyPrefix      string
		requestTimeout time.Duration
		progressEvery  int
		binaryMin      int
		binaryMax      int
		includeText    bool
		includeJSON    bool
		includeBinary  bool
		globalSeed     int64
	)

	flag.StringVar(&endpointList, "endpoints", "localhost:2379", "comma-separated etcd endpoints")
	flag.IntVar(&totalKeys, "n", 1000, "total number of keys to create")
	flag.IntVar(&concurrency, "c", 16, "number of concurrent writer goroutines")
	flag.StringVar(&keyPrefix, "prefix", "datasets/loadtest", "base key prefix for generated keys")
	flag.DurationVar(&requestTimeout, "timeout", 5*time.Second, "per-request timeout duration")
	flag.IntVar(&progressEvery, "progress", 1000, "log progress every N successful writes (0 disables progress logs)")
	flag.IntVar(&binaryMin, "binary-min", 32, "minimum bytes for generated binary payloads")
	flag.IntVar(&binaryMax, "binary-max", 256, "maximum bytes for generated binary payloads")
	flag.BoolVar(&includeText, "text", true, "include plain text values")
	flag.BoolVar(&includeJSON, "json", true, "include JSON values")
	flag.BoolVar(&includeBinary, "binary", true, "include binary values")
	flag.Int64Var(&globalSeed, "seed", time.Now().UnixNano(), "seed for deterministic runs")
	flag.Parse()

	if totalKeys <= 0 {
		log.Fatalf("invalid -n value: must be > 0")
	}
	if concurrency <= 0 {
		log.Fatalf("invalid -c value: must be > 0")
	}
	if binaryMin <= 0 {
		log.Fatalf("invalid -binary-min value: must be > 0")
	}
	if binaryMax < binaryMin {
		log.Printf("adjusting binary-max (%d) to binary-min (%d)", binaryMax, binaryMin)
		binaryMax = binaryMin
	}

	valueTypes := buildValueTypeList(includeText, includeJSON, includeBinary)
	if len(valueTypes) == 0 {
		log.Println("no value types selected, defaulting to text")
		valueTypes = []valueType{valueTypeText}
	}

	endpoints := parseEndpointList(endpointList)
	if len(endpoints) == 0 {
		log.Fatalf("no endpoints provided")
	}

	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("failed to create etcd client: %v", err)
	}
	defer client.Close()

	start := time.Now()
	ctx := context.Background()
	var (
		completed int64
		failures  int64
	)

	jobs := make(chan int)
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for workerID := 0; workerID < concurrency; workerID++ {
		workerSeed := globalSeed + int64(workerID+1)*1_000_003
		go func(id int, seed int64) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(seed))
			faker := gofakeit.New(uint64(seed))
			for idx := range jobs {
				vt := valueTypes[rng.Intn(len(valueTypes))]
				key, value := generateKeyValue(faker, rng, keyPrefix, vt, idx, binaryMin, binaryMax)

				writeCtx, cancel := context.WithTimeout(ctx, requestTimeout)
				_, putErr := client.Put(writeCtx, key, string(value))
				cancel()

				if putErr != nil {
					atomic.AddInt64(&failures, 1)
					log.Printf("[worker %d] failed to put key %q: %v", id, key, putErr)
					continue
				}

				newCount := atomic.AddInt64(&completed, 1)
				if progressEvery > 0 && newCount%int64(progressEvery) == 0 {
					log.Printf("progress: %d/%d keys written", newCount, totalKeys)
				}
			}
		}(workerID, workerSeed)
	}

	for i := 0; i < totalKeys; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	elapsed := time.Since(start)

	log.Printf("load complete in %s", elapsed.Round(time.Millisecond))
	log.Printf("successful writes: %d, failures: %d", completed, failures)
	if elapsed > 0 {
		rate := float64(completed) / elapsed.Seconds()
		log.Printf("throughput: %.2f writes/sec", rate)
	}
}

func buildValueTypeList(text, jsonEnabled, binary bool) []valueType {
	var types []valueType
	if text {
		types = append(types, valueTypeText)
	}
	if jsonEnabled {
		types = append(types, valueTypeJSON)
	}
	if binary {
		types = append(types, valueTypeBinary)
	}
	return types
}

func parseEndpointList(raw string) []string {
	parts := strings.Split(raw, ",")
	var endpoints []string
	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			endpoints = append(endpoints, trimmed)
		}
	}
	return endpoints
}

func generateKeyValue(faker *gofakeit.Faker, rng *rand.Rand, prefix string, vt valueType, idx int, binaryMin, binaryMax int) (string, []byte) {
	suffix := fmt.Sprintf("%016x", rng.Uint64())
	switch vt {
	case valueTypeJSON:
		key := fmt.Sprintf("%s/json/%09d-%s", prefix, idx, suffix)
		price := faker.Price(10, 5000)
		payload := map[string]interface{}{
			"id":          fmt.Sprintf("order-%d", idx),
			"customer":    faker.Name(),
			"email":       faker.Email(),
			"product":     faker.AppName(),
			"price":       price,
			"purchasedAt": time.Now().Add(-time.Duration(rng.Intn(86_400)) * time.Second).UTC().Format(time.RFC3339),
			"tags":        []string{faker.BuzzWord(), faker.Company()},
		}
		b, err := json.Marshal(payload)
		if err != nil {
			log.Printf("json marshal failed for idx=%d: %v", idx, err)
			return key, []byte(faker.Sentence(12))
		}
		return key, b

	case valueTypeBinary:
		key := fmt.Sprintf("%s/blob/%09d-%s", prefix, idx, suffix)
		size := binaryMin
		if binaryMax > binaryMin {
			size += rng.Intn(binaryMax - binaryMin + 1)
		}
		buf := make([]byte, size)
		if _, err := rng.Read(buf); err != nil {
			log.Printf("rng read failed for idx=%d: %v", idx, err)
		}
		return key, buf

	default:
		key := fmt.Sprintf("%s/text/%09d-%s", prefix, idx, suffix)
		value := faker.Sentence(16)
		return key, []byte(value)
	}
}
