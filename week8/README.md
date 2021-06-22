# Week8 Homework
## Redis
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

## Part1 
redis-benchmark
```shell
redis-benchmark  -q -d 10
PING_INLINE: 90991.81 requests per second, p50=0.263 msec                   
PING_MBULK: 89686.10 requests per second, p50=0.263 msec                   
SET: 90252.70 requests per second, p50=0.263 msec                   
GET: 88417.33 requests per second, p50=0.263 msec                   
INCR: 88888.89 requests per second, p50=0.263 msec                   
LPUSH: 89047.20 requests per second, p50=0.263 msec                   
RPUSH: 90991.81 requests per second, p50=0.263 msec                   
LPOP: 92336.11 requests per second, p50=0.263 msec                   
RPOP: 91240.88 requests per second, p50=0.263 msec                   
SADD: 90334.23 requests per second, p50=0.263 msec                    
HSET: 89605.73 requests per second, p50=0.263 msec                   
SPOP: 91491.30 requests per second, p50=0.255 msec                   
ZADD: 93109.87 requests per second, p50=0.255 msec                   
ZPOPMIN: 91996.32 requests per second, p50=0.255 msec                   
LPUSH (needed to benchmark LRANGE): 92764.38 requests per second, p50=0.255 msec                   
LRANGE_100 (first 100 elements): 40683.48 requests per second, p50=0.591 msec                   
LRANGE_300 (first 300 elements): 16409.58 requests per second, p50=1.455 msec                   
LRANGE_500 (first 500 elements): 10485.48 requests per second, p50=2.319 msec                   
LRANGE_600 (first 600 elements): 8726.00 requests per second, p50=2.735 msec                  
MSET (10 keys): 100806.45 requests per second, p50=0.239 msec    

redis-benchmark  -q -d 20
PING_INLINE: 93283.58 requests per second, p50=0.255 msec                   
PING_MBULK: 91324.20 requests per second, p50=0.263 msec                   
SET: 90009.01 requests per second, p50=0.263 msec                   
GET: 92592.59 requests per second, p50=0.255 msec                   
INCR: 91911.76 requests per second, p50=0.263 msec                   
LPUSH: 90090.09 requests per second, p50=0.263 msec                   
RPUSH: 94517.96 requests per second, p50=0.255 msec                   
LPOP: 94339.62 requests per second, p50=0.247 msec                    
RPOP: 94073.38 requests per second, p50=0.255 msec                   
SADD: 93109.87 requests per second, p50=0.263 msec                   
HSET: 90252.70 requests per second, p50=0.263 msec                   
SPOP: 92850.51 requests per second, p50=0.255 msec                   
ZADD: 93370.68 requests per second, p50=0.255 msec                   
ZPOPMIN: 94786.73 requests per second, p50=0.255 msec                   
LPUSH (needed to benchmark LRANGE): 93545.37 requests per second, p50=0.255 msec                   
LRANGE_100 (first 100 elements): 41373.60 requests per second, p50=0.591 msec                   
LRANGE_300 (first 300 elements): 16572.75 requests per second, p50=1.463 msec                   
LRANGE_500 (first 500 elements): 10578.65 requests per second, p50=2.279 msec                   
LRANGE_600 (first 600 elements): 9074.41 requests per second, p50=2.671 msec                  
MSET (10 keys): 99108.03 requests per second, p50=0.247 msec 

redis-benchmark  -q -d 100
PING_INLINE: 88028.16 requests per second, p50=0.263 msec                   
PING_MBULK: 89206.06 requests per second, p50=0.263 msec                   
SET: 88339.23 requests per second, p50=0.255 msec                   
GET: 88105.73 requests per second, p50=0.263 msec                   
INCR: 92421.44 requests per second, p50=0.255 msec                   
LPUSH: 91575.09 requests per second, p50=0.255 msec                   
RPUSH: 93023.25 requests per second, p50=0.255 msec                   
LPOP: 95328.88 requests per second, p50=0.255 msec                    
RPOP: 92250.92 requests per second, p50=0.255 msec                   
SADD: 93720.71 requests per second, p50=0.263 msec                   
HSET: 91240.88 requests per second, p50=0.263 msec                   
SPOP: 93196.65 requests per second, p50=0.255 msec                   
ZADD: 94607.38 requests per second, p50=0.255 msec                   
ZPOPMIN: 92764.38 requests per second, p50=0.263 msec                   
LPUSH (needed to benchmark LRANGE): 93283.58 requests per second, p50=0.255 msec                   
LRANGE_100 (first 100 elements): 38580.25 requests per second, p50=0.631 msec                   
LRANGE_300 (first 300 elements): 15365.70 requests per second, p50=1.511 msec                   
LRANGE_500 (first 500 elements): 9125.75 requests per second, p50=1.807 msec                  
LRANGE_600 (first 600 elements): 7726.78 requests per second, p50=2.159 msec                  
MSET (10 keys): 96618.36 requests per second, p50=0.271 msec 

redis-benchmark  -q -d 200
PING_INLINE: 92592.59 requests per second, p50=0.255 msec                   
PING_MBULK: 86132.64 requests per second, p50=0.263 msec                   
SET: 87108.02 requests per second, p50=0.271 msec                   
GET: 90909.09 requests per second, p50=0.255 msec                   
INCR: 87642.41 requests per second, p50=0.263 msec                   
LPUSH: 94161.95 requests per second, p50=0.255 msec                   
RPUSH: 92592.59 requests per second, p50=0.255 msec                   
LPOP: 96153.85 requests per second, p50=0.247 msec                    
RPOP: 95602.30 requests per second, p50=0.247 msec                    
SADD: 90909.09 requests per second, p50=0.263 msec                   
HSET: 92081.03 requests per second, p50=0.255 msec                   
SPOP: 83822.30 requests per second, p50=0.271 msec                   
ZADD: 85178.88 requests per second, p50=0.271 msec                   
ZPOPMIN: 95419.85 requests per second, p50=0.255 msec                    
LPUSH (needed to benchmark LRANGE): 95238.10 requests per second, p50=0.247 msec                   
LRANGE_100 (first 100 elements): 30075.19 requests per second, p50=0.759 msec                   
LRANGE_300 (first 300 elements): 11782.73 requests per second, p50=1.375 msec                   
LRANGE_500 (first 500 elements): 5903.19 requests per second, p50=1.799 msec                  
LRANGE_600 (first 600 elements): 4598.12 requests per second, p50=2.079 msec                  
MSET (10 keys): 93896.71 requests per second, p50=0.279 msec 

redis-benchmark  -q -d 1000
PING_INLINE: 94876.66 requests per second, p50=0.255 msec                    
PING_MBULK: 83542.19 requests per second, p50=0.271 msec                   
SET: 76804.91 requests per second, p50=0.295 msec                   
GET: 90497.73 requests per second, p50=0.263 msec                   
INCR: 83822.30 requests per second, p50=0.279 msec                   
LPUSH: 85763.29 requests per second, p50=0.279 msec                   
RPUSH: 75018.76 requests per second, p50=0.327 msec                   
LPOP: 87950.75 requests per second, p50=0.263 msec                   
RPOP: 81699.35 requests per second, p50=0.287 msec                   
SADD: 87873.46 requests per second, p50=0.271 msec                   
HSET: 90415.91 requests per second, p50=0.271 msec                   
SPOP: 91996.32 requests per second, p50=0.263 msec                   
ZADD: 95328.88 requests per second, p50=0.263 msec                    
ZPOPMIN: 93457.94 requests per second, p50=0.263 msec                   
LPUSH (needed to benchmark LRANGE): 92421.44 requests per second, p50=0.263 msec                   
LRANGE_100 (first 100 elements): 13449.90 requests per second, p50=0.807 msec                   
LRANGE_300 (first 300 elements): 4525.09 requests per second, p50=1.039 msec                  
LRANGE_500 (first 500 elements): 2574.14 requests per second, p50=1.183 msec                  
LRANGE_600 (first 600 elements): 2062.03 requests per second, p50=1.239 msec                  
MSET (10 keys): 72780.20 requests per second, p50=0.351 msec        


redis-benchmark  -q -d 5000
PING_INLINE: 92165.90 requests per second, p50=0.255 msec                   
PING_MBULK: 90334.23 requests per second, p50=0.263 msec                   
SET: 80000.00 requests per second, p50=0.295 msec                   
GET: 80000.00 requests per second, p50=0.287 msec                   
INCR: 86805.56 requests per second, p50=0.271 msec                   
LPUSH: 83542.19 requests per second, p50=0.295 msec                   
RPUSH: 84245.99 requests per second, p50=0.295 msec                   
LPOP: 83056.48 requests per second, p50=0.287 msec                   
RPOP: 91074.68 requests per second, p50=0.263 msec                   
SADD: 89847.26 requests per second, p50=0.271 msec                   
HSET: 83194.67 requests per second, p50=0.295 msec                   
SPOP: 91074.68 requests per second, p50=0.271 msec                   
ZADD: 90744.10 requests per second, p50=0.271 msec                   
ZPOPMIN: 88967.98 requests per second, p50=0.287 msec                   
LPUSH (needed to benchmark LRANGE): 84602.37 requests per second, p50=0.303 msec                   
LRANGE_100 (first 100 elements): 3506.93 requests per second, p50=2.903 msec                  
LRANGE_300 (first 300 elements): 1005.85 requests per second, p50=4.183 msec                  
LRANGE_500 (first 500 elements): 541.82 requests per second, p50=4.943 msec                   
LRANGE_600 (first 600 elements): 425.79 requests per second, p50=5.535 msec                   
MSET (10 keys): 33355.57 requests per second, p50=1.327 msec 
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