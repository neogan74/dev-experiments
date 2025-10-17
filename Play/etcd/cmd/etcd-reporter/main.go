package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type valueType string

const (
	valueTypeText   valueType = "text"
	valueTypeJSON   valueType = "json"
	valueTypeBinary valueType = "binary"
	valueTypeOther  valueType = "other"
)

type summary struct {
	TotalKeys int64
	Stats     map[valueType]*typeStats
}

type typeStats struct {
	Count        int64
	TotalBytes   int64
	JSONPriceSum float64
	JSONPriceMin float64
	JSONPriceMax float64
}

func newTypeStats() *typeStats {
	return &typeStats{
		JSONPriceMin: math.Inf(1),
		JSONPriceMax: math.Inf(-1),
	}
}

func main() {
	var (
		endpointList   string
		prefix         string
		batchSize      int
		requestTimeout time.Duration
		serializable   bool
	)

	flag.StringVar(&endpointList, "endpoints", "localhost:2379", "comma-separated etcd endpoints")
	flag.StringVar(&prefix, "prefix", "datasets/loadtest", "key prefix to scan")
	flag.IntVar(&batchSize, "batch", 1000, "number of keys to fetch per request")
	flag.DurationVar(&requestTimeout, "timeout", 10*time.Second, "per-request timeout duration")
	flag.BoolVar(&serializable, "serializable", false, "use serializable reads instead of linearizable")
	flag.Parse()

	if batchSize <= 0 {
		log.Fatalf("invalid -batch value: must be > 0")
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

	ctx := context.Background()
	report, err := collect(ctx, client, prefix, batchSize, requestTimeout, serializable)
	if err != nil {
		log.Fatalf("collection failed: %v", err)
	}

	printReport(prefix, report)
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

func collect(ctx context.Context, client *clientv3.Client, prefix string, batchSize int, timeout time.Duration, serializable bool) (*summary, error) {
	startKey := prefix
	rangeEnd := clientv3.GetPrefixRangeEnd(prefix)
	opts := []clientv3.OpOption{
		clientv3.WithRange(rangeEnd),
		clientv3.WithLimit(int64(batchSize)),
	}
	if serializable {
		opts = append(opts, clientv3.WithSerializable())
	}

	stats := &summary{
		Stats: map[valueType]*typeStats{
			valueTypeText:   newTypeStats(),
			valueTypeJSON:   newTypeStats(),
			valueTypeBinary: newTypeStats(),
			valueTypeOther:  newTypeStats(),
		},
	}

	seenKeys := int64(0)

	for {
		reqCtx, cancel := context.WithTimeout(ctx, timeout)
		resp, err := client.Get(reqCtx, startKey, opts...)
		cancel()
		if err != nil {
			return nil, fmt.Errorf("get failed for start %q: %w", startKey, err)
		}

		if len(resp.Kvs) == 0 {
			break
		}

		for _, kv := range resp.Kvs {
			seenKeys++
			processKV(stats, kv.Key, kv.Value)
		}

		if !resp.More {
			break
		}

		lastKey := resp.Kvs[len(resp.Kvs)-1].Key
		startKey = nextKey(lastKey)
	}

	stats.TotalKeys = seenKeys
	return stats, nil
}

func processKV(sum *summary, key, value []byte) {
	k := string(key)
	t := detectValueType(k)
	ts := sum.Stats[t]
	if ts == nil {
		ts = newTypeStats()
		sum.Stats[t] = ts
	}

	ts.Count++
	ts.TotalBytes += int64(len(value))

	if t == valueTypeJSON {
		var obj map[string]interface{}
		if err := json.Unmarshal(value, &obj); err != nil {
			return
		}
		if priceVal, ok := obj["price"]; ok {
			price, ok := toFloat(priceVal)
			if !ok {
				return
			}
			ts.JSONPriceSum += price
			if price < ts.JSONPriceMin {
				ts.JSONPriceMin = price
			}
			if price > ts.JSONPriceMax {
				ts.JSONPriceMax = price
			}
		}
	}
}

func detectValueType(key string) valueType {
	switch {
	case strings.Contains(key, "/json/"):
		return valueTypeJSON
	case strings.Contains(key, "/blob/"):
		return valueTypeBinary
	case strings.Contains(key, "/text/"):
		return valueTypeText
	default:
		return valueTypeOther
	}
}

func toFloat(v interface{}) (float64, bool) {
	switch num := v.(type) {
	case float64:
		return num, true
	case float32:
		return float64(num), true
	case int:
		return float64(num), true
	case int64:
		return float64(num), true
	case json.Number:
		f, err := num.Float64()
		if err != nil {
			return 0, false
		}
		return f, true
	default:
		return 0, false
	}
}

func nextKey(b []byte) string {
	k := append([]byte{}, b...)
	for i := len(k) - 1; i >= 0; i-- {
		if k[i] < 0xFF {
			k[i]++
			return string(k[:i+1])
		}
	}
	return string(append(k, 0x00))
}

func printReport(prefix string, s *summary) {
	fmt.Println("=== etcd dataset report ===")
	fmt.Printf("Prefix: %s\n", prefix)
	fmt.Printf("Total keys: %d\n", s.TotalKeys)

	fmt.Println("\nValue types:")
	for _, t := range []valueType{valueTypeText, valueTypeJSON, valueTypeBinary, valueTypeOther} {
		ts := s.Stats[t]
		if ts == nil {
			continue
		}
		fmt.Printf("  - %s: %d keys (%.2f KB total)\n", t, ts.Count, float64(ts.TotalBytes)/1024)
		if t == valueTypeJSON && ts.Count > 0 && !math.IsInf(ts.JSONPriceMin, 1) {
			avg := ts.JSONPriceSum / float64(ts.Count)
			min := ts.JSONPriceMin
			max := ts.JSONPriceMax
			if math.IsInf(min, 1) {
				min = 0
			}
			if math.IsInf(max, -1) {
				max = 0
			}
			fmt.Printf("      prices -> avg: %.2f, min: %.2f, max: %.2f\n", avg, min, max)
		}
	}
}
