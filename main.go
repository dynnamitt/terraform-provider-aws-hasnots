package main

import (
  "fmt"
  "log"
  "context"
	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main(){

  cfg, err := config.LoadDefaultConfig(context.TODO())
  if err != nil {
    log.Fatalf("failed to load configuration, %v", err)
  }

  client := iam.NewFromConfig(cfg)

  fmt.Println("Calling aws---")
  resp, err := client.ListRoles(context.TODO(), &iam.ListRolesInput{})
  if err != nil {
    log.Fatal(err)
  }


    fmt.Println("Roles:")
    for _, role := range resp.Roles {
        fmt.Println(*role.Arn)
    }

}
