# dashboard 

can add extra dashboard here ~

with database notice with : `prometheus-db`

> please refer `go-prometheus.json` template.


# refer:
- https://grafana.com/grafana/dashboards/6671
- https://github.com/iamgoangle/go-grafana-prometheus

# 關於增加新的監控指標

增加新的量測指標
1. 打開http://127.0.0.1:9090
2. 選擇要監測的指標後執行`execute`
3. 觀察語法e.g.
4. 將下面的模板適當修改後貼至go-client-default.json
```json
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "datasource": "prometheus-db",
      "editable": true,
      "error": false,
      "fill": 1,
      "grid": {},
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 0,
        "y": 22
      },
      "id": 9,
      "isNew": true,
      "legend": {
        "alignAsTable": true,
        "avg": true,
        "current": true,
        "max": true,
        "min": false,
        "show": true,
        "total": false,
        "values": true
      },
      "lines": true,
      "linewidth": 2,
      "links": [],
      "nullPointMode": "connected",
      "percentage": false,
      "pointradius": 5,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "go_threads{namespace=~\"^($namespace)$\",pod=~\"^($pod)$\"}",
          "format": "time_series",
          "intervalFactor": 2,
          "legendFormat": "{{pod}}",
          "metric": "go_threads",
          "refId": "A",
          "step": 4
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeShift": null,
      "title": "Go Threads",
      "tooltip": {
        "msResolution": false,
        "shared": true,
        "sort": 0,
        "value_type": "cumulative"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "datasource": "prometheus-db",
      "editable": true,
      "error": false,
      "fill": 1,
      "grid": {},
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 22
      },
      "id": 10,
      "isNew": true,
      "legend": {
        "alignAsTable": true,
        "avg": true,
        "current": true,
        "max": true,
        "min": false,
        "show": true,
        "total": false,
        "values": true
      },
      "lines": true,
      "linewidth": 2,
      "links": [],
      "nullPointMode": "connected",
      "percentage": false,
      "pointradius": 5,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "rate(go_memstats_frees_total{namespace=~\"^($namespace)$\",pod=~\"($pod)$\"}[30s])",
          "format": "time_series",
          "intervalFactor": 2,
          "legendFormat": "{{pod}} - alloc free",
          "metric": "go_gc_duration_seconds",
          "refId": "A",
          "step": 4
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeShift": null,
      "title": "go free_total rate",
      "tooltip": {
        "msResolution": false,
        "shared": true,
        "sort": 0,
        "value_type": "cumulative"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "s",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
```

5. 重啟docker-compose
