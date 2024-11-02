package ec2

import "fmt"

type VPC struct {
    ID string
}

type Geo struct {
    REGION string
    AVAILABILITY_ZONE string
    VPC VPC
}

type EFS struct {
    ID string
    NAME string
    GEO Geo
}

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
