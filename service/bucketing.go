package service

import (
	"encoding/json"
	"fmt"

	"kilvish.io/cerebro/config"
	"kilvish.io/cerebro/entity"
)

type BucketingService struct {
	Kp *KafkaPublisher
}

func (bs *BucketingService) InferBucket(aar *entity.AutoAssignRequest) string {
	if aar.Type == config.HomeLoan {
		return config.HomeLoan
	}
	if aar.Type == config.CarLoan {
		return config.HomeLoan
	}
	if aar.Type == config.EducationLoan {
		return config.HomeLoan
	}
	return "default"
}

func (bs *BucketingService) AddRequestToBucket(bucket string, aar *entity.AutoAssignRequest) string {
	request, _ := json.Marshal(aar)
	topic := "kilvish-" + bucket
	bs.Kp.Publish(topic, string(request))
	return "request-Added"
}

func (bs *BucketingService) Assign(aar *entity.AutoAssignRequest) string {
	bucket := bs.InferBucket(aar)
	fmt.Println(bucket)
	status := bs.AddRequestToBucket(bucket, aar)
	fmt.Println(status)
	return bucket
}
