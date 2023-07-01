## JustMySocks Exporter

### 项目简介

`justmysocks_exporter` 是一个用 Go 语言编写的 Prometheus Exporter，用于监控 JustMySocks 服务的带宽使用情况。通过调用 JustMySocks 提供的 API，获取服务的月度带宽限制、已使用带宽以及带宽重置日期等信息，并将这些信息以 Prometheus 指标的形式暴露出来。

### 功能

- 获取 JustMySocks 服务的月度带宽限制（单位：字节）
- 获取 JustMySocks 服务的已使用带宽（单位：字节）
- 获取 JustMySocks 服务的带宽重置日期（每月的某一天）

### 如何使用

1. 克隆项目到本地：

```
git clone https://github.com/yourusername/justmysocks_exporter.git
```

2. 进入项目目录并编译：

```
cd justmysocks_exporter
go build
```

3. 运行编译好的程序，传入必要的参数：

```
./justmysocks_exporter --api-address "https://justmysocks5.net/members/getbwcounter.php" --service "your_service_number" --id "your_uuid"
```

4. 在 Prometheus 配置文件中添加 `justmysocks_exporter` 作为监控目标：

```
scrape_configs:
  - job_name: 'justmysocks_exporter'
    static_configs:
      - targets: ['localhost:10001']
```

5. 重启 Prometheus 服务以应用新的配置。

### 指标说明

- `justmysocks_monthly_bw_limit_bytes`：JustMySocks 服务的月度带宽限制（单位：字节）
- `justmysocks_bw_counter_bytes`：JustMySocks 服务的已使用带宽（单位：字节）
- `justmysocks_bw_reset_day_of_month`：JustMySocks 服务的带宽重置日期（每月的某一天）

### 注意事项

- 本项目仅适用于 JustMySocks 服务，不适用于其他服务提供商。
- 请确保在运行程序时提供正确的 API 地址、服务编号和 UUID。
- 默认情况下，程序会每隔 5 分钟更新一次指标数据。如需修改更新频率，请在源代码中修改 `time.Sleep()` 函数的参数。