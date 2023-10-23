package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolaev-roman/simbir-go/config"
	"gorm.io/gorm"
)

type Server struct {
	gin *gin.Engine
	cfg *config.Config
	db  *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{gin: gin.Default(), cfg: cfg, db: db}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.gin); err != nil {
		return err
	}

	s.gin.Run(s.cfg.Server.Port)
	return nil
}
