## JustMySocks Exporter

[中文README](https://github.com/wulabing/justmysocks_exporter/README_CN.MD)
### Project Introduction

`justmysocks_exporter` is a Prometheus Exporter written in Go language, designed to monitor the bandwidth usage of JustMySocks services. By calling the API provided by JustMySocks, it obtains information such as the monthly bandwidth limit, used bandwidth, and bandwidth reset date of the service, and exposes this information in the form of Prometheus metrics.

### Features

- Obtain the monthly bandwidth limit of JustMySocks service (unit: bytes)
- Obtain the used bandwidth of JustMySocks service (unit: bytes)
- Obtain the bandwidth reset date of JustMySocks service (a specific day of each month)

### How to Use

1. Clone the project to your local machine:

```
git clone https://github.com/yourusername/justmysocks_exporter.git
```

2. Enter the project directory and compile:

```
cd justmysocks_exporter
go build
```

3. Run the compiled program, passing in the necessary parameters:

```
./justmysocks_exporter --api-address "https://justmysocks5.net/members/getbwcounter.php" --service "your_service_number" --id "your_uuid"
```

4. Add `justmysocks_exporter` as a monitoring target in the Prometheus configuration file:

```
scrape_configs:
  - job_name: 'justmysocks_exporter'
    static_configs:
      - targets: ['localhost:10001']
```

5. Restart the Prometheus service to apply the new configuration.

### Metric Descriptions

- `justmysocks_monthly_bw_limit_bytes`: Monthly bandwidth limit of JustMySocks service (unit: bytes)
- `justmysocks_bw_counter_bytes`: Used bandwidth of JustMySocks service (unit: bytes)
- `justmysocks_bw_reset_day_of_month`: Bandwidth reset date of JustMySocks service (a specific day of each month)

### Precautions

- This project is only applicable to JustMySocks services and is not suitable for other service providers.
- Please make sure to provide the correct API address, service number, and UUID when running the program.
- By default, the program updates the metric data every 5 minutes. If you need to modify the update frequency, please change the parameter of the `time.Sleep()` function in the source code.