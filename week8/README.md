# Redis
```shell
sudo docker run --name my-redis -p 6379:6379 -d redis
sudo docker exec -it my-redis sh

redis-cli
redis-benchmark


```

redis-cli
```shell
MEMORY HELP
MEMORY USAGE <key>
MEMORY USAGE <key> SAMPLES <count>
INFO MEMORY
```

# Part1 
redis-benchmark
```shell
redis-benchmark -d 10
Summary:
  throughput summary: 97560.98 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.299     0.104     0.279     0.447     0.647     1.463


redis-benchmark -d 20
Summary:
  throughput summary: 89766.61 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.310     0.128     0.295     0.439     0.679     1.663

redis-benchmark -d 50
Summary:
  throughput summary: 82850.04 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.346     0.136     0.327     0.503     0.783     2.007


redis-benchmark -d 100
Summary:
  throughput summary: 83125.52 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.372     0.128     0.343     0.615     0.863     1.703
        
redis-benchmark -d 200
Summary:
  throughput summary: 65832.78 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.449     0.160     0.415     0.663     1.031     2.007

redis-benchmark -d 1000
Summary:
  throughput summary: 52826.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.549     0.160     0.463     1.039     1.447    13.615
        
redis-benchmark -d 5000
Summary:
  throughput summary: 26652.45 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.714     0.424     1.471     2.935     4.087    31.135

```

## Part2

empty
```shell
# Memory
used_memory:871872
```

10bytes
平均每个 key 的占用内存空间: 84
```shell
# Memory
used_memory:1714944
```
20 bytes
平均每个 key 的占用内存空间: 92
```shell
# Memory
used_memory:1794944
```

50 bytes
平均每个 key 的占用内存空间: 124
```shell
# Memory
used_memory:2114944
```

100 bytes
平均每个 key 的占用内存空间: 180
```shell
# Memory
used_memory:2674944
```

200 bytes
平均每个 key 的占用内存空间:  292
```shell
# Memory
used_memory:3794944
```

1000 bytes
平均每个 key 的占用内存空间: 1092
```shell
# Memory
used_memory:11794944
```

5000 bytes
平均每个 key 的占用内存空间: 5188
```shell
# Memory
used_memory:52754944
```