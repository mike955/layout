package service

import (
	"context"
	"errors"

	"github.com/mike955/zrpc-layout/internal/data"
	"github.com/sirupsen/logrus"

	pb "github.com/mike955/zrpc-layout/api/layout"
)

type LayoutService struct {
	pb.UnimplementedLayoutServiceServer
	logger *logrus.Entry
	data   *data.LayoutData
}

func NewLayoutService(logger *logrus.Entry) *LayoutService {
	return &LayoutService{
		logger: logger,
		data:   data.NewLayoutData(logger),
	}
}

func (s *LayoutService) Test(ctx context.Context, request *pb.TestRequest) (response *pb.TestResponse, err error) {
	response = new(pb.TestResponse)
	if logger := ctx.Value("logger"); logger != nil {
		s.logger = logger.(*logrus.Entry)
		s.data.SetLogger(logger.(*logrus.Entry))
	} else {
		ctx = context.WithValue(ctx, "logger", s.logger)
	}
	res, err := s.data.Test()
	if err != nil {
		s.logger.Errorf("app:layout|service:layout|func:test|request:%+v|error:%s", request, err.Error())
		err = errors.New("test error")
		return
	}
	response.Data = res
	return
}
