package service

import (
	"github.com/liqieli/raiden/internal/repository"
	"github.com/liqieli/raiden/pkg/helper/sid"
	"github.com/liqieli/raiden/pkg/jwt"
	"github.com/liqieli/raiden/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
