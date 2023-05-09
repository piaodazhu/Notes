# Redis quick start (redis v6)

## A. types
1. string: any value, less than 512MB.
2. hash: object with some field=>value.
3. list: simple strings list.
4. set: unordered set of strings. (implemented with hash, fast)
5. zset: ordered set of strings with `score`.

## B. key
```sh
127.0.0.1:6379> set testkey hello
OK
127.0.0.1:6379> get testkey
"hello"
127.0.0.1:6379> del testkey
(integer) 1
127.0.0.1:6379> del testkey
(integer) 0
127.0.0.1:6379> set testkey hello
OK
127.0.0.1:6379> dump testkey
"\x00\x05hello\t\x00\xb3\x80\x8e\xba1\xb2C\xbb"
127.0.0.1:6379> exists testkey
(integer) 1
127.0.0.1:6379> expire testkey 1
(integer) 1
127.0.0.1:6379> exists testkey
(integer) 0
127.0.0.1:6379> set testkey hello
OK
127.0.0.1:6379> keys te*
1) "testkey"
127.0.0.1:6379> ttl testkey
(integer) -1
127.0.0.1:6379> expire testkey 1000
(integer) 1
127.0.0.1:6379> ttl testkey
(integer) 998
127.0.0.1:6379> rename testkey testKey
OK
127.0.0.1:6379> ttl testkey
(integer) -2
127.0.0.1:6379> ttl testKey
(integer) 971
127.0.0.1:6379> type testKey
string
```

## C. string
```sh
127.0.0.1:6379> set testKey 'I am test key'
OK
127.0.0.1:6379> get testKey
"I am test key"
127.0.0.1:6379> getrange testKey 1 3
" am"
127.0.0.1:6379> getrange testKey 1 -1
" am test key"
127.0.0.1:6379> getset testKey "I'm test key"
"I am test key"
127.0.0.1:6379> get testKey
"I'm test key"
127.0.0.1:6379> set testKey2 testKey
OK
127.0.0.1:6379> get testKey
"I'm test key"
127.0.0.1:6379> mget testKey testKey2
1) "I'm test key"
2) "testKey"

127.0.0.1:6379> setex testKey3 100 12
OK
127.0.0.1:6379> get testKey3
"12"
127.0.0.1:6379> ttl testKey3
(integer) 85
127.0.0.1:6379> incr testKey3
(integer) 13
127.0.0.1:6379> 
```

## D. hash

```sh
127.0.0.1:6379> hmset xiaoming id 123 name xiaoming sex M age 19
OK
127.0.0.1:6379> hgetall xiaoming
1) "id"
2) "123"
3) "name"
4) "xiaoming"
5) "age"
6) "19"
7) "sex"
8) "M"
127.0.0.1:6379> hdel xiaoming sex
(integer) 1
127.0.0.1:6379> hexists xiaoming sex
(integer) 0
127.0.0.1:6379> hincrby xiaoming age 1
(integer) 20
127.0.0.1:6379> hkeys xiaoming
1) "id"
2) "name"
3) "age"
127.0.0.1:6379> hlen xiaoming
(integer) 3
127.0.0.1:6379> hmget xiaoming id
1) "123"
127.0.0.1:6379> hset xiaoming id 1234
(integer) 0
127.0.0.1:6379> hvals xiaoming
1) "1234"
2) "xiaoming"
3) "20"
```

## E. list

```sh
127.0.0.1:6379> lpush class xiaoming
(integer) 1
127.0.0.1:6379> lpush class xiaoli
(integer) 2
127.0.0.1:6379> lpush class xiaozhang
(integer) 3
127.0.0.1:6379> lrange class 0 10
1) "xiaozhang"
2) "xiaoli"
3) "xiaoming"
127.0.0.1:6379> rpush class zhangsan
(integer) 4
127.0.0.1:6379> lrange class 0 10
1) "xiaozhang"
2) "xiaoli"
3) "xiaoming"
4) "zhangsan"
127.0.0.1:6379> blpop class 1
1) "class"
2) "xiaozhang"
127.0.0.1:6379> lrange class 0 10
1) "xiaoli"
2) "xiaoming"
3) "zhangsan"
127.0.0.1:6379> brpop class 1
1) "class"
2) "zhangsan"
127.0.0.1:6379> brpoplpush class newclass 1
"xiaoming"
127.0.0.1:6379> lrange newclass 0 10
1) "xiaoming"
127.0.0.1:6379> llen class
(integer) 1
```

## F. set

```sh
127.0.0.1:6379> sadd member xiaoming
(integer) 1
127.0.0.1:6379> sadd member xiaoli
(integer) 1
127.0.0.1:6379> sadd member zhangsan
(integer) 1
127.0.0.1:6379> sadd member xiaoli
(integer) 0
127.0.0.1:6379> smembers member
1) "zhangsan"
2) "xiaoli"
3) "xiaoming"
127.0.0.1:6379> scard member
(integer) 3
127.0.0.1:6379> sadd member2 xiaoming xiaoli goudan
(integer) 3
127.0.0.1:6379> sadd member3 xiaoli goudan xiaowang
(integer) 3
127.0.0.1:6379> smembers member2
1) "goudan"
2) "xiaoli"
3) "xiaoming"
127.0.0.1:6379> smembers member3
1) "goudan"
2) "xiaowang"
3) "xiaoli"
127.0.0.1:6379> sdiff member member2
1) "zhangsan"
127.0.0.1:6379> sdiff member member3
1) "zhangsan"
2) "xiaoming"
127.0.0.1:6379> sinter member member2
1) "xiaoli"
2) "xiaoming"
127.0.0.1:6379> sinter member member2 member3
1) "xiaoli"
127.0.0.1:6379> sismember member zhangsan
(integer) 1
127.0.0.1:6379> sismember member2 zhangsan
(integer) 0
127.0.0.1:6379> spop member3
"goudan"
127.0.0.1:6379> srem member zhangsan
(integer) 1
127.0.0.1:6379> smembers member
1) "xiaoli"
2) "xiaoming"
127.0.0.1:6379> sunion member member2 member3
1) "goudan"
2) "xiaowang"
3) "xiaoli"
4) "xiaoming"
```

## G. zset

```sh
127.0.0.1:6379> zadd class 1 xiaoming
(error) WRONGTYPE Operation against a key holding the wrong kind of value
127.0.0.1:6379> zadd company 1 xiaoming
(integer) 1
127.0.0.1:6379> zadd company 2 wanger
(integer) 1
127.0.0.1:6379> zadd company 3 zhangsan
(integer) 1
127.0.0.1:6379> zrange company 0 10 withscores
1) "xiaoming"
2) "1"
3) "wanger"
4) "2"
5) "zhangsan"
6) "3"
127.0.0.1:6379> zcount company 2 3
(integer) 2
127.0.0.1:6379> zrem company xiaoming
(integer) 1
127.0.0.1:6379> zrange company 0 10 
1) "wanger"
2) "zhangsan"
127.0.0.1:6379> zadd company 3 lisi 5 wangwu 6 chenliu
(integer) 3
127.0.0.1:6379> zrangebyscore company 3 5
1) "lisi"
2) "zhangsan"
3) "wangwu"
127.0.0.1:6379> zscore company wangwu
"5"
127.0.0.1:6379> zrange company 0 10 withscores
 1) "wanger"
 2) "2"
 3) "lisi"
 4) "3"
 5) "zhangsan"
 6) "3"
 7) "wangwu"
 8) "5"
 9) "chenliu"
10) "6"
127.0.0.1:6379> zadd company 4 lisi
(integer) 0
127.0.0.1:6379> zrange company 0 10 withscores
 1) "wanger"
 2) "2"
 3) "zhangsan"
 4) "3"
 5) "lisi"
 6) "4"
 7) "wangwu"
 8) "5"
 9) "chenliu"
10) "6"
127.0.0.1:6379> zrange company 0 1
1) "wanger"
2) "zhangsan"
127.0.0.1:6379> zrange company 0 0
1) "wanger"
```