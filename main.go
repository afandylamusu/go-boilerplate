package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/afandylamusu/stnkku.mdm/customer"
	"github.com/afandylamusu/stnkku.mdm/dbconn"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/jinzhu/gorm"
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

func startKafkaConsumers(c *kafka.Consumer) {

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
func startGrpcServer(port string, db *gorm.DB) error {
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

	customer.RegisterCustomerQueryServiceServer(s, &customer.ServiceHandler{Port: port, Db: db})
	log.Println("Run GRPC Server at port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return err
}

func main() {
	isLocal := viper.GetString("env") == "local"
	var dbHost, dbPort, dbUser, dbPass, dbName string

	if isLocal {
		dbHost = viper.GetString(`database-local.host`)
		dbPort = viper.GetString(`database-local.port`)
		dbUser = viper.GetString(`database-local.user`)
		dbPass = viper.GetString(`database-local.pass`)
		dbName = viper.GetString(`database-local.name`)
	} else {
		dbHost = viper.GetString(`database.host`)
		dbPort = viper.GetString(`database.port`)
		dbUser = viper.GetString(`database.user`)
		dbPass = viper.GetString(`database.pass`)
		dbName = viper.GetString(`database.name`)
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		log.Fatalf("failed to established db connection: %v", err)
	}

	defer db.Close()

	dbconn.Migrate(db)

	// lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// KAFKA Consumers

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:9092,0.0.0.0:9093,0.0.0.0:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	defer c.Close()

	if err != nil {
		panic(err)
	}

	go startKafkaConsumers(c)

	go startGrpcServer(viper.GetString("server.grpc-port"), db)

	<-make(chan int)
}
