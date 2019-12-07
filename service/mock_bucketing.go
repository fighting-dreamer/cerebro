package service

import (
	"github.com/stretchr/testify/mock"
	"kilvish.io/cerebro/entity"
)

type MockBucketingService struct {
	mock.Mock
}

func (mbs *MockBucketingService) Assign(aar *entity.AutoAssignRequest) string {
	args := mbs.Called(aar)
	return args.Get(0).(string)
}
