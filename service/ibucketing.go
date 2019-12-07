package service

import "kilvish.io/cerebro/entity"

type IBucketingService interface {
	Assign(aar *entity.AutoAssignRequest) string
}
