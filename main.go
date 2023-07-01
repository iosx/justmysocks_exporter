package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

type JustMySocksData struct {
	MonthlyBWLimitB   uint64 `json:"monthly_bw_limit_b"`
	BWCounterB        uint64 `json:"bw_counter_b"`
	BWResetDayOfMonth int    `json:"bw_reset_day_of_month"`
}

var (
	monthlyBWLimitB = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "justmysocks_monthly_bw_limit_bytes",
			Help: "Monthly bandwidth limit in bytes",
		},
		[]string{"service"})
	bwCounterB = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "justmysocks_bw_counter_bytes",
			Help: "Bandwidth bandwidth in bytes",
		}, []string{"service"})
	bwResetDayOfMonth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "justmysocks_bw_reset_day_of_month",
			Help: "Day of month for bandwidth reset",
		}, []string{"service"})
)

func fetchJustMySocksAPIData(apiAddress string, service string, id string) (*JustMySocksData, error) {
	resp, err := http.Get(fmt.Sprintf("%s?service=%s&id=%s", apiAddress, service, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data JustMySocksData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func updateMetrics(apiAddress string, service string, id string) {
	data, err := fetchJustMySocksAPIData(apiAddress, service, id)
	if err != nil {
		log.Println("Error fetching data: ", err)
	} else {
		log.Println("Fetched data: ", *data)
	}

	monthlyBWLimitB.WithLabelValues(service).Set(float64(data.MonthlyBWLimitB))
	bwCounterB.WithLabelValues(service).Set(float64(data.BWCounterB))
	bwResetDayOfMonth.WithLabelValues(service).Set(float64(data.BWResetDayOfMonth))
}

var (
	// Default port allocation https://github.com/prometheus/prometheus/wiki/Default-port-allocations
	listenAddress = flag.String("web.listen-address", ":10001", "Address to listen on for web interface and telemetry.")
	apiAddress    = flag.String("api-address", "https://justmysocks5.net/members/getbwcounter.php", "Address of JustMySocks API")
	service       = flag.String("service", "", "JustMySocks service number")
	id            = flag.String("id", "", "JustMySocks UUID")
)

func main() {
	flag.Parse()
	log.SetOutput(os.Stdout)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		for {
			updateMetrics(*apiAddress, *service, *id)
			// Update every 5 minutes,Becasue the data is updated every 5 minutes by justmysocks
			time.Sleep(5 * time.Minute)
		}
	}()
	prometheus.MustRegister(monthlyBWLimitB, bwCounterB, bwResetDayOfMonth)
	log.Printf("Listening on %s...", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
