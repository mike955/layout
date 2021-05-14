package data

import (
	"github.com/mike955/zrpc-layout/internal/dao"
	"github.com/sirupsen/logrus"
)

type LayoutData struct {
	logger *logrus.Entry
	dao    *dao.LayoutDao
}

func NewLayoutData(logger *logrus.Entry) *LayoutData {
	return &LayoutData{
		logger: logger,
		dao:    dao.NewLayoutDao(),
	}
}

func (s *LayoutData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

func (s *LayoutData) Test() (res uint64, err error) {
	res = 200
	return
}
