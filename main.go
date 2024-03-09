package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {
	err := SendSMS("+11111111", "Merhaba, bu bir test mesajıdır")
	if err != nil {
		log.Fatalf("SMS gönderme hatası: %v", err)
	}
}

func SendSMS(phoneNumber, message string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("key12341243", "secretkey431123123", "")),
	)
	if err != nil {
		return fmt.Errorf("AWS yapılandırması yüklenirken hata: %w", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	input := &sns.PublishInput{
		Message:     &message,
		PhoneNumber: &phoneNumber,
	}

	_, err = snsClient.Publish(context.Background(), input)
	if err != nil {
		return fmt.Errorf("SMS gönderilirken hata: %w", err)
	}

	fmt.Println("SMS başarıyla gönderildi!")
	return nil
}
