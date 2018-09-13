# Kafka Go(lang) quick start

Docker based example for Kafka using Go client

- `git clone https://github.com/abhirockzz/kafka-go-docker-quickstart`
- `cd kafka-go-quickstart`
- edit `docker-compose.yml` and provide values for `KAFKA_BROKER` and `KAFKA_TOPIC` variables as per your setup
- start - `docker-compose up`
- check logs

		consumer_1  | waiting for event...
		consumer_1  | Message 2018-09-13 16:20:08.506907019 +0000 UTC m=+8854.499461595
		producer_1  | Delivered message to offset 7432 in partition test[0]@7432
		consumer_1  | waiting for event...
		consumer_1  | EOF at test[1]@7404(Broker: No more messages)
		consumer_1  | waiting for event...
		consumer_1  | OffsetsCommitted (<nil>, [test[1]@7404])
		producer_1  | Delivered message to offset 7403 in partition test[1]@7403
		consumer_1  | waiting for event...
		producer_1  | Delivered message to offset 7433 in partition test[0]@7433
		consumer_1  | Message 2018-09-13 16:20:13.514229263 +0000 UTC m=+8859.506783855
		producer_1  | Delivered message to offset 7434 in partition test[0]@7434
		consumer_1  | waiting for event...
		producer_1  | Delivered message to offset 7404 in partition test[1]@7404
		consumer_1  | Message 2018-09-13 16:20:13.517535428 +0000 UTC m=+8859.510090009
		consumer_1  | waiting for event...
		consumer_1  | Message 2018-09-13 16:20:13.520525347 +0000 UTC m=+8859.513079916
		producer_1  | Delivered message to offset 7405 in partition test[1]@7405
		consumer_1  | waiting for event...
		consumer_1  | EOF at test[1]@7407(Broker: No more messages)
		consumer_1  | waiting for event...
		consumer_1  | OffsetsCommitted (<nil>, [test[1]@7407])
		consumer_1  | waiting for event...
		consumer_1  | Message 2018-09-13 16:20:18.524126429 +0000 UTC m=+8864.516680997
		producer_1  | Delivered message to offset 7406 in partition test[1]@7406
		consumer_1  | waiting for event...
		consumer_1  | Message 2018-09-13 16:20:18.528697235 +0000 UTC m=+8864.521251825
		producer_1  | Delivered message to offset 7407 in partition test[1]@7407
		producer_1  | Delivered message to offset 7408 in partition test[1]@7408
		producer_1  | Delivered message to offset 7435 in partition test[0]@7435
		producer_1  | Delivered message to offset 7436 in partition test[0]@7436