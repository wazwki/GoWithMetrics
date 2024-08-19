## How to create golang HTTP-server with metrics

### 1. Integrating Prometheus into a Go Application

Prometheus uses a metrics format that can be integrated into a Go application using the `prometheus/client_golang` library.

1. **Install the library**:
   ```bash
   go get github.com/prometheus/client_golang/prometheus
   go get github.com/prometheus/client_golang/prometheus/promhttp
   ```

2. **Export metrics**:
   Add code to your application to export metrics. For example:

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
           requestCount.With(prometheus.Labels{"method": r.Method, "handler": "/"})
           w.Write([]byte("Hello, Prometheus!"))
       })

       // Register the metrics handler
       http.Handle("/metrics", promhttp.Handler())

       http.ListenAndServe(":8080", nil)
   }
   ```

   In this example, metrics are available at the `/metrics` path.

### 2. Installing and Configuring Prometheus

1. **Download Prometheus**:
   Go to the [official Prometheus website](https://prometheus.io/download/) and download the latest version.

2. **Configure Prometheus**:
   Create a configuration file `prometheus.yml`:

   ```yaml
   global:
     scrape_interval: 15s

   scrape_configs:
     - job_name: "golang_app"
       static_configs:
         - targets: ["localhost:8080"]
   ```

   This configuration file instructs Prometheus to collect metrics from your Go application running on port 8080.

3. **Run Prometheus**:
   Start Prometheus with the following command:

   ```bash
   ./prometheus --config.file=prometheus.yml
   ```

   By default, the Prometheus interface is available at `http://localhost:9090`.

### 3. Installing and Configuring Grafana

1. **Download and Install Grafana**:
   Go to the [official Grafana website](https://grafana.com/grafana/download) and install Grafana following the instructions for your operating system.

2. **Start Grafana**:
   Start Grafana:

   ```bash
   sudo systemctl start grafana-server
   ```

   Once started, the Grafana interface will be available at `http://localhost:3000`. The default login is `admin`, and the password is `admin`.

3. **Set Up Data Source**:
   - Log in to the Grafana interface.
   - Go to `Configuration` -> `Data Sources`.
   - Click `Add data source` and select `Prometheus`.
   - Enter the Prometheus URL (`http://localhost:9090`) and save.

4. **Create a Dashboard**:
   - Go to `Create` -> `Dashboard`.
   - Click `Add new panel` and select the metric you want to visualize (e.g., `http_requests_total`).
   - Configure the graph display and save the panel.

### 4. Monitoring and Visualization

You can now monitor your metrics through the Grafana interface. Based on the data you export from your Go application, you can create various visualizations and dashboards to monitor your application's performance.

### 5. (Optional) Setting Up Alerts

You can set up alerts in Prometheus and Grafana to notify you of issues with your application. In Prometheus, alerts are configured through Alertmanager, and in Grafana, you can set up notifications through integrations with various messaging and email services.
