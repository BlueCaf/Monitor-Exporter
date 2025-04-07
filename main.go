package main

import (
    "flag"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    log "github.com/sirupsen/logrus"
)

var (
    // 수동 버전 메트릭 등록
    exporterInfo = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "query_exporter_build_info",
            Help: "Build information for the query exporter",
        },
        []string{"version", "revision"},
    )

    // 버전 정보 - 필요시 go build -ldflags로 주입 가능
    version  = "1.0.0"
    revision = "dev"
)

func init() {
    exporterInfo.WithLabelValues(version, revision).Set(1)
    prometheus.MustRegister(exporterInfo)
}

func main() {
    var bind string
    flag.StringVar(&bind, "bind", "0.0.0.0:9104", "HTTP bind address")
    flag.Parse()

    http.Handle("/metrics", promhttp.Handler())

    log.Infof("Starting HTTP server on %s", bind)
    if err := http.ListenAndServe(bind, nil); err != nil {
        log.Fatalf("HTTP server failed: %v", err)
    }
}
