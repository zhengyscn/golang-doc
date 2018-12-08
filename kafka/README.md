# 简易日志搜集


## jdk
```
$ java -version
java version "1.8.0_191"
Java(TM) SE Runtime Environment (build 1.8.0_191-b12)
Java HotSpot(TM) 64-Bit Server VM (build 25.191-b12, mixed mode)
```

## zookeeper
```
zookeeper-3.4.13.tar.gz

$ tar -zxvf zookeeper-3.4.13.tar.gz
$ cd zookeeper-3.4.13
$ cp conf/zoo_sample.cfg conf/zoo.cfg
$ ./bin/zkServer.sh start  # port: 2181
```

## kafka
```
kafka_2.11-2.1.0.tgz

$ cd kafka_2.11-2.1.0

producer
$ ./bin/kafka-server-start.sh config/server.properties

consumer
$ ./kafka-console-consumer.sh --topic important --bootstrap-server localhost:9092 --from-beginning
$ ./kafka-console-consumer.sh --topic important --bootstrap-server localhost:9092
```