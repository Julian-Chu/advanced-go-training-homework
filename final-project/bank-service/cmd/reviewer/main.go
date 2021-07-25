package main

import (
	accounts "bankservice/api/accountsservice"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
	"os"
	"os/signal"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"0.0.0.0:9092"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("applications", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0

	accountsClient := NewAccountsServiceClient()

ConsumerLoop:
	for {
		fmt.Println("consuming...")
		select {
		case msg := <-partitionConsumer.Messages():
			username := string(msg.Value)
			log.Printf("Consumed message offset %d, username: %s\n", msg.Offset, username)
			consumed++
			accountReply, err := accountsClient.CreateAccount(context.Background(), &accounts.CreateAccountRequest{Username: username})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("account:%s created for username:%s\n", accountReply.AccountID, username)
			}

		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

func NewAccountsServiceClient() accounts.AccountClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9001"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return accounts.NewAccountClient(conn)
}
