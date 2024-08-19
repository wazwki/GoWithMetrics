## Как создать golang HTTP-сервер с метриками

### 1. Интеграция Prometheus в приложение на Go

Prometheus использует формат метрик, который можно интегрировать в приложение на Go с помощью библиотеки `prometheus/client_golang`.

1. **Установка библиотеки**:
   ```bash
   go get github.com/prometheus/client_golang/prometheus
   go get github.com/prometheus/client_golang/prometheus/promhttp
   ```

2. **Экспорт метрик**:
   Добавьте в ваше приложение код для экспорта метрик. Например:

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

       // Регистрируем обработчик метрик
       http.Handle("/metrics", promhttp.Handler())

       http.ListenAndServe(":8080", nil)
   }
   ```

   В этом примере метрики доступны по пути `/metrics`.

### 2. Установка и настройка Prometheus

1. **Скачайте Prometheus**:
   Перейдите на [официальный сайт Prometheus](https://prometheus.io/download/) и скачайте последнюю версию.

2. **Настройка Prometheus**:
   Создайте файл конфигурации `prometheus.yml`:

   ```yaml
   global:
     scrape_interval: 15s

   scrape_configs:
     - job_name: "golang_app"
       static_configs:
         - targets: ["localhost:8080"]
   ```

   Этот конфигурационный файл указывает Prometheus собирать метрики с вашего приложения на Go, работающего на порту 8080.

3. **Запуск Prometheus**:
   Запустите Prometheus с помощью следующей команды:

   ```bash
   ./prometheus --config.file=prometheus.yml
   ```

   По умолчанию интерфейс Prometheus доступен по адресу `http://localhost:9090`.

### 3. Установка и настройка Grafana

1. **Скачайте и установите Grafana**:
   Перейдите на [официальный сайт Grafana](https://grafana.com/grafana/download) и установите Grafana, следуя инструкциям для вашей операционной системы.

2. **Запуск Grafana**:
   Запустите Grafana:

   ```bash
   sudo systemctl start grafana-server
   ```

   После запуска интерфейс Grafana будет доступен по адресу `http://localhost:3000`. Логин по умолчанию: `admin`, пароль: `admin`.

3. **Настройка источника данных**:
   - Зайдите в интерфейс Grafana и войдите в систему.
   - Перейдите в раздел `Configuration` -> `Data Sources`.
   - Нажмите `Add data source` и выберите `Prometheus`.
   - Введите URL Prometheus (`http://localhost:9090`) и сохраните.

4. **Создание дашборда**:
   - Перейдите в раздел `Create` -> `Dashboard`.
   - Нажмите `Add new panel` и выберите метрику, которую вы хотите визуализировать (например, `http_requests_total`).
   - Настройте отображение графика и сохраните панель.

### 4. Мониторинг и визуализация

Теперь вы можете наблюдать за своими метриками через интерфейс Grafana. На основе данных, которые вы экспортируете из вашего Go-приложения, вы можете создавать различные визуализации и дашборды для мониторинга производительности вашего приложения.

### 5. (Опционально) Настройка алертов

Вы можете настроить алерты в Prometheus и Grafana для уведомления о проблемах с вашим приложением. В Prometheus алерты настраиваются через Alertmanager, а в Grafana можно настроить оповещения через интеграции с различными мессенджерами и почтовыми сервисами.
