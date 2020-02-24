package users

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/jackc/pgx"
	"github.com/bugscatcher/messages/application"
	"github.com/bugscatcher/messages/services"
)

type Handler struct {
	db            *pgx.ConnPool
	kafkaProducer sarama.SyncProducer
}

func (h *Handler) SendMessage(context.Context, *services.SendMessageRequest) (*services.Response, error) {
	panic("implement me")
}

func New(app *application.App) *Handler {
	return &Handler{
		db:            app.DB,
		kafkaProducer: app.KafkaProducer,
	}
}
