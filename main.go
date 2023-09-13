package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const namespace = "df"

var (
	usedSpace = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "used_space_kilobytes",
			Help:      "Used space in kilobytes.",
		},
		[]string{"filesystem"},
	)

	availableSpace = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "available_space_kilobytes",
			Help:      "Available space in kilobytes.",
		},
		[]string{"filesystem"},
	)
)

func init() {
	prometheus.MustRegister(usedSpace)
	prometheus.MustRegister(availableSpace)
}

func collectMetrics() {
	cmd := exec.Command("df", "-k")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Scan() // skip the header line

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 6 {
			continue
		}

		filesystem := fields[0]
		used, _ := strconv.ParseFloat(fields[2], 64)
		available, _ := strconv.ParseFloat(fields[3], 64)

		usedSpace.WithLabelValues(filesystem).Set(used)
		availableSpace.WithLabelValues(filesystem).Set(available)
	}

	cmd.Wait()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			collectMetrics()
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
