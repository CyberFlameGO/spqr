package client

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/jackc/pgproto3/v2"
)

type Client interface {
	ID() string

	ReplyErr(errmsg string) error
	ReplyNotice(msg string) error
	DefaultReply() error

	Init(cfg *tls.Config, reqssl string) error

	Auth() error

	PasswordCT() string
	PasswordMD5() string

	StartupMessage() *pgproto3.StartupMessage

	Usr() string
	DB() string

	Send(msg pgproto3.BackendMessage) error
	Receive() (pgproto3.FrontendMessage, error)
	ProcQuery(query *pgproto3.Query) (byte, error)

	Shutdown() error
	Reset() error

	Reply(msg string) error
}

type InteractRunner interface {
	ProcClient(ctx context.Context, conn net.Conn) error
}
