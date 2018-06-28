# Zenit

[Zenit](https://en.wikipedia.org/wiki/Zenit_(satellite)) Project is missing DBA knife tool. Zenit is a russian was spy satellite.

This tool collect stats data from:

- Linux OS (CentOS)
- MySQL
- Percona ToolKit
- ProxySQL

The numeric values has represent time has in microseconds.

## ProxySQL

### Configure

Allow remote access:

```bash
mysql -u admin -padmin -h 127.0.0.1 -P 6032
SET admin-admin_credentials = "admin:admin;radminuser:radminpass";
LOAD ADMIN VARIABLES TO RUNTIME;
```

## Install & Configure

```bash
chown root. zenit
mv zenit /usr/local/bin/
export DSN_MYSQL="monitor:monitor@tcp(10.9.35.40:3306)/"
```

## Prometheus

Integration for Prometheus, in this example is add the follow commands into cron:

```cron
* * * * * /usr/local/bin/zenit -collect="mysql" > /usr/local/prometheus/textfile_collector/zenit.prom
*/5 * * * * DSN_MYSQL="monitor:monitor@tcp(10.9.35.40:3306)/" /usr/local/bin/zenit -collect="mysql-tables,mysql-overflow" >> /usr/local/prometheus/textfile_collector/zenit.prom
```

## Development

Build, upload to docker container and run:

```bash
GOOS=linux go build -ldflags "-s -w" -o zenit main.go && \
docker cp zenit d1c86f2f36ff:/root && \
docker exec -i -t d1c86f2f36ff /root/zenit -collect-os
```

You only need update the ID container from last command.
