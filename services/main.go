package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		// THIS SHOULD BE COMING FROM A CONFIG FILE
		APIKey: "WmF4NzFvNEJ5LVFaT0RPLVdNWDU6OUxENS1QLXdRTTJscjk0eDMxS3BNQQ==",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	infores, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	fmt.Println(infores)
	// This should be dynamic and coming from ../storage/....
	// We need a way of reading specific files or all files
	// Then parse it to new buffer string as bulk action
	// Havent found a good way of doing this
	// !!!! THERE ARE ALSO SOME CERT ISSUES
	buf := bytes.NewBufferString(`{"index":{"_index":"collections"}}
	{"symbol":"BTC-USD","price":70955.33}
	{"index":{"_index":"collections"}}
	{"symbol":"ETH-USD","price":3545.42}
	{"index":{"_index":"collections"}}
	{"symbol":"USDT-USD","price":1.0001}
	{"index":{"_index":"collections"}}
	{"symbol":"BNB-USD","price":617.57}
	{"index":{"_index":"collections"}}
	{"symbol":"SOL-USD","price":174.87}
	{"index":{"_index":"collections"}}
	{"symbol":"XRP-USD","price":0.612684}
	{"index":{"_index":"collections"}}
	{"symbol":"STETH-USD","price":3538.17}
	{"index":{"_index":"collections"}}
	{"symbol":"USDC-USD","price":1.0001}
	{"index":{"_index":"collections"}}
	{"symbol":"DOGE-USD","price":0.198185}
	{"index":{"_index":"collections"}}
	{"symbol":"TON11419-USD","price":7.2437}
	{"index":{"_index":"collections"}}
	{"symbol":"ADA-USD","price":0.589281}
	{"index":{"_index":"collections"}}
	{"symbol":"AVAX-USD","price":46.77}
	{"index":{"_index":"collections"}}
	{"symbol":"SHIB-USD","price":0.000028}
	{"index":{"_index":"collections"}}
	{"symbol":"DOT-USD","price":8.4616}
	{"index":{"_index":"collections"}}
	{"symbol":"BCH-USD","price":612.37}
	{"index":{"_index":"collections"}}
	{"symbol":"WBTC-USD","price":71000.55}
	{"index":{"_index":"collections"}}
	{"symbol":"TRX-USD","price":0.121630}
	{"index":{"_index":"collections"}}
	{"symbol":"WTRX-USD","price":0.121245}
	{"index":{"_index":"collections"}}
	{"symbol":"LINK-USD","price":17.81}
	{"index":{"_index":"collections"}}
	{"symbol":"MATIC-USD","price":0.888672}
	{"index":{"_index":"collections"}}
	{"symbol":"LTC-USD","price":98.62}
	{"index":{"_index":"collections"}}
	{"symbol":"NEAR-USD","price":6.8924}
	{"index":{"_index":"collections"}}
	{"symbol":"ICP-USD","price":15.75}
	{"index":{"_index":"collections"}}
	{"symbol":"UNI7083-USD","price":9.0659}
	{"index":{"_index":"collections"}}
	{"symbol":"LEO-USD","price":5.8029}
`)

	ingestResult, err := es.Bulk(
		bytes.NewReader(buf.Bytes()),
		es.Bulk.WithIndex("collections"),
		es.Bulk.WithPipeline("ent-search-generic-ingestion"),
	)

	fmt.Println(ingestResult, err)
}

// WmF4NzFvNEJ5LVFaT0RPLVdNWDU6OUxENS1QLXdRTTJscjk0eDMxS3BNQQ==
