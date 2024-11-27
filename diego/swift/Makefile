obu:
	@go run obu/main.go

.PHONY: obu

rece:
	@go run obu_receive/main.go

netw:
	@docker network create kafka-network

zook:
	@docker run -d --network=kafka-network --name=zookeeper \
    -e ZOOKEEPER_CLIENT_PORT=2181 \
    -e ZOOKEEPER_TICK_TIME=2000 \
    confluentinc/cp-zookeeper:latest

kafka:
	@docker run -d --network=kafka-network --name=kafka \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
    -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
    -e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1 \
    -e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1 \
    confluentinc/cp-kafka:latest

topic:
	@docker run --network=kafka-network --rm confluentinc/cp-kafka:latest \
    kafka-topics --create --topic test_topic --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:2181

# need to fix is
pull:
	@docker pull confluentinc/cp-zookeeper:latest
	@docker pull confluentinc/cp-kafka:latest
