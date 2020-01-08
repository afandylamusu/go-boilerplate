package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/afandylamusu/moonlay.mcservice/customer"
	"github.com/afandylamusu/moonlay.mcservice/dbconn"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func startKafkaConsumers(c *kafka.Consumer, db *dbconn.DbConnection) {

	// c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)
	c.SubscribeTopics([]string{
		"newCustomerTopic", "updateCustomerTopic", "removeCustomerTopic",
		"newKurirTopic", "updateKurirTopic", "removeKurirTopic",
		"newVehicleTopic", "updateVehicleTopic", "removeVehicleTopic",
	}, nil)

	log.Println("Run Kafka Consumers")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// startGrpcServer to starting GRPC Server
func startGrpcServer(port string, db *dbconn.DbConnection) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Second * 10,
			Timeout:           time.Second * 20,
		}),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Second,
				PermitWithoutStream: true,
			}),
		grpc.MaxConcurrentStreams(5),
	)

	customer.RegisterCustomerQueryServiceServer(s, &customer.GrpcHandler{Port: port, Db: db})
	log.Println("Run GRPC Server at port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return err
}

func main() {

	// Connection for KAFKA
	conn := &dbconn.DbConnection{}
	conn.Open()

	// Connection for GRPC
	conn2 := &dbconn.DbConnection{}
	conn2.Open()

	defer conn.Close()
	defer conn2.Close()

	// KAFKA Consumers
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:9092",
		"group.id":          "g-customer-opr",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	go startKafkaConsumers(c, conn)

	// new connection for grpc
	startGrpcServer(viper.GetString("server.grpc-port"), conn2)

	<-make(chan int)
}
