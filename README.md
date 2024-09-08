# How to Set Up Prometheus and Grafana in a Golang HTTP Server

## 1. Integrating Prometheus into a Go Application

Prometheus uses a metric format that can be integrated into a Go application using the `prometheus/client_golang` library.

### **1.1. Installing the library**:

```bash
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
```

### **1.2. Exporting metrics**:

Add the following code to your application to export metrics. For example:

```go
package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "handler"},
    )
)

func init() {
    prometheus.MustRegister(requestCount)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        requestCount.With(prometheus.Labels{"method": r.Method, "handler": "/"}).Inc()
        w.Write([]byte("Hello, Prometheus!"))
    })

    // Register the metrics handler
    http.Handle("/metrics", promhttp.Handler())

    http.ListenAndServe(":8080", nil)
}
```

In this example, metrics are available at the `/metrics` path.

## 2. Installing and Setting Up Prometheus

### **2.1. Download Prometheus**:

Go to the [official Prometheus website](https://prometheus.io/download/) and download the latest version.

### **2.2. Configuring Prometheus**:

Create a `prometheus.yml` configuration file:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: "golang_app"
    static_configs:
      - targets: ["localhost:8080"]
```

This configuration tells Prometheus to scrape metrics from your Go application running on port 8080.

### **2.3. Running Prometheus**:

Start Prometheus with the following command:

```bash
./prometheus --config.file=prometheus.yml
```

By default, Prometheus will be accessible at `http://localhost:9090`.

## 3. Installing and Setting Up Grafana

### **3.1. Download and Install Grafana**:

Go to the [official Grafana website](https://grafana.com/grafana/download) and follow the instructions to install Grafana for your operating system.

### **3.2. Running Grafana**:

Start Grafana:

```bash
sudo systemctl start grafana-server
```

Once started, the Grafana UI will be available at `http://localhost:3000`. The default login is `admin`, and the default password is `admin`.

### **3.3. Setting Up a Data Source**:

- Log in to Grafana.
- Go to `Configuration` -> `Data Sources`.
- Click `Add data source` and select `Prometheus`.
- Enter the Prometheus URL (`http://localhost:9090`) and save.

### **3.4. Creating a Dashboard**:

- Go to `Create` -> `Dashboard`.
- Click `Add new panel`, and select the metric you want to visualize (e.g., `http_requests_total`).
- Customize the graph display and save the panel.

## 4. Monitoring and Visualization

Now, you can observe your metrics through the Grafana dashboard. Based on the data exported from your Go application, you can create various visualizations and dashboards to monitor your application's performance.

## 5. Setting Up Alerts

You can configure alerts in both Prometheus and Grafana to notify you of issues with your application. In Prometheus, alerts are managed via Alertmanager, while Grafana allows integration with various messaging and email services for sending notifications.