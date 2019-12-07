package service

import (
	"fmt"

	"kilvish.io/cerebro/config"
	"kilvish.io/cerebro/entity"
)

type BucketingService struct {
	Kp *KafkaPublisher
}

func (bs *BucketingService) InferBucket(aar *entity.AutoAssignRequest) string {
	return config.HomeLoan
}

func (bs *BucketingService) AddRequestToBucket(bucket string, aar *entity.AutoAssignRequest) string {

	return "request-Added"
}

func (bs *BucketingService) Assign(aar *entity.AutoAssignRequest) string {
	bucket := bs.InferBucket(aar)
	fmt.Println(bucket)
	status := bs.AddRequestToBucket(bucket, aar)
	fmt.Println(status)
	return bucket
}
