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


empty
```shell
# Memory
used_memory:871872
used_memory_human:851.44K
used_memory_rss:8708096
used_memory_rss_human:8.30M
used_memory_peak:1858592
used_memory_peak_human:1.77M
used_memory_peak_perc:46.91%
used_memory_overhead:830360
used_memory_startup:809856
used_memory_dataset:41512
used_memory_dataset_perc:66.94%
allocator_allocated:903976
allocator_active:1830912
allocator_resident:4661248
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.03
allocator_frag_bytes:926936
allocator_rss_ratio:2.55
allocator_rss_bytes:2830336
rss_overhead_ratio:1.87
rss_overhead_bytes:4046848
mem_fragmentation_ratio:10.48
mem_fragmentation_bytes:7877240
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

```

10bytes
平均每个 key 的占用内存空间: 84
```shell
# Memory
used_memory:1714944
used_memory_human:1.64M
used_memory_rss:8912896
used_memory_rss_human:8.50M
used_memory_peak:1858592
used_memory_peak_human:1.77M
used_memory_peak_perc:92.27%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:353512
used_memory_dataset_perc:39.06%
allocator_allocated:1730280
allocator_active:2007040
allocator_resident:4841472
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.16
allocator_frag_bytes:276760
allocator_rss_ratio:2.41
allocator_rss_bytes:2834432
rss_overhead_ratio:1.84
rss_overhead_bytes:4071424
mem_fragmentation_ratio:5.32
mem_fragmentation_bytes:7238968
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

```
20 bytes
平均每个 key 的占用内存空间: 92
```shell
# Memory
used_memory:1794944
used_memory_human:1.71M
used_memory_rss:9158656
used_memory_rss_human:8.73M
used_memory_peak:1938528
used_memory_peak_human:1.85M
used_memory_peak_perc:92.59%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:433512
used_memory_dataset_perc:44.01%
allocator_allocated:1931728
allocator_active:2232320
allocator_resident:5066752
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.16
allocator_frag_bytes:300592
allocator_rss_ratio:2.27
allocator_rss_bytes:2834432
rss_overhead_ratio:1.81
rss_overhead_bytes:4091904
mem_fragmentation_ratio:5.22
mem_fragmentation_bytes:7404728
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

```

50 bytes
平均每个 key 的占用内存空间: 124
```shell
# Memory
used_memory:2114944
used_memory_human:2.02M
used_memory_rss:9433088
used_memory_rss_human:9.00M
used_memory_peak:2135944
used_memory_peak_human:2.04M
used_memory_peak_perc:99.02%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:753512
used_memory_dataset_perc:57.74%
allocator_allocated:2141064
allocator_active:2416640
allocator_resident:5373952
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.13
allocator_frag_bytes:275576
allocator_rss_ratio:2.22
allocator_rss_bytes:2957312
rss_overhead_ratio:1.76
rss_overhead_bytes:4059136
mem_fragmentation_ratio:4.55
mem_fragmentation_bytes:7359160
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

```

100 bytes
平均每个 key 的占用内存空间: 180
```shell
# Memory
used_memory:2674944
used_memory_human:2.55M
used_memory_rss:9867264
used_memory_rss_human:9.41M
used_memory_peak:2695944
used_memory_peak_human:2.57M
used_memory_peak_perc:99.22%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:1313512
used_memory_dataset_perc:70.43%
allocator_allocated:2696104
allocator_active:2998272
allocator_resident:5898240
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.11
allocator_frag_bytes:302168
allocator_rss_ratio:1.97
allocator_rss_bytes:2899968
rss_overhead_ratio:1.67
rss_overhead_bytes:3969024
mem_fragmentation_ratio:3.75
mem_fragmentation_bytes:7233336
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

200 bytes
平均每个 key 的占用内存空间:  292
```shell
# Memory
used_memory:3794944
used_memory_human:3.62M
used_memory_rss:11137024
used_memory_rss_human:10.62M
used_memory_peak:3815944
used_memory_peak_human:3.64M
used_memory_peak_perc:99.45%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:2433512
used_memory_dataset_perc:81.52%
allocator_allocated:3821336
allocator_active:4145152
allocator_resident:7053312
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.08
allocator_frag_bytes:323816
allocator_rss_ratio:1.70
allocator_rss_bytes:2908160
rss_overhead_ratio:1.58
rss_overhead_bytes:4083712
mem_fragmentation_ratio:2.97
mem_fragmentation_bytes:7383096
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0


```

1000 bytes
平均每个 key 的占用内存空间: 1092
```shell
# Memory
used_memory:11794944
used_memory_human:11.25M
used_memory_rss:19337216
used_memory_rss_human:18.44M
used_memory_peak:11815944
used_memory_peak_human:11.27M
used_memory_peak_perc:99.82%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:10433512
used_memory_dataset_perc:94.98%
allocator_allocated:11846600
allocator_active:12181504
allocator_resident:15372288
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.03
allocator_frag_bytes:334904
allocator_rss_ratio:1.26
allocator_rss_bytes:3190784
rss_overhead_ratio:1.26
rss_overhead_bytes:3964928
mem_fragmentation_ratio:1.65
mem_fragmentation_bytes:7583288
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

```

5000 bytes
平均每个 key 的占用内存空间: 5188
```shell
# Memory
used_memory:52754944
used_memory_human:50.31M
used_memory_rss:70975488
used_memory_rss_human:67.69M
used_memory_peak:52775944
used_memory_peak_human:50.33M
used_memory_peak_perc:99.96%
used_memory_overhead:1361432
used_memory_startup:809856
used_memory_dataset:51393512
used_memory_dataset_perc:98.94%
allocator_allocated:52786208
allocator_active:53071872
allocator_resident:66826240
total_system_memory:33546264576
total_system_memory_human:31.24G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.01
allocator_frag_bytes:285664
allocator_rss_ratio:1.26
allocator_rss_bytes:13754368
rss_overhead_ratio:1.06
rss_overhead_bytes:4149248
mem_fragmentation_ratio:1.35
mem_fragmentation_bytes:18261560
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
 
```