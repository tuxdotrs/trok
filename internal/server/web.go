package server

import (
	"crypto/tls"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/rs/zerolog/log"
	"github.com/tuxdotrs/trok/internal/web"
	"golang.org/x/crypto/acme/autocert"
)

type TrokWeb struct {
	app  *fiber.App
	addr string
}

func NewTrokWeb(addr string) *TrokWeb {
	return &TrokWeb{
		app:  fiber.New(),
		addr: addr,
	}
}

func (t *TrokWeb) Start() {
	t.app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(web.EmbedDirStatic),
		PathPrefix: "dist",
		Browse:     true,
	}))

	t.app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(web.EmbedDirStatic),
		PathPrefix: "dist/assets",
		Browse:     true,
	}))

	cfg := t.GetTLSCert()

	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil {
		log.Panic().Msgf("unable to start trok webserver: %v", err)
	}

	t.app.Listener(ln)
}

func (t *TrokWeb) GetTLSCert() *tls.Config {
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("trok.cloud"),
		Cache:      autocert.DirCache("./certs"),
	}

	return &tls.Config{
		GetCertificate: m.GetCertificate,
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}
}

func (t *TrokWeb) Stop() {
	t.app.Shutdown()
}
