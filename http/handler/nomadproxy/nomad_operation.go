package nomadproxy

import (
	"net/http"

	"github.com/portainer/agent"

	httperror "github.com/portainer/libhttp/error"
	"github.com/rs/zerolog/log"
)

func (handler *Handler) nomadOperation(rw http.ResponseWriter, request *http.Request) *httperror.HandlerError {
	log.Debug().
		Str("nomad_request_url", request.URL.String()).
		Msg("nomadOperation")

	request.Header.Set(agent.HTTPNomadTokenHeaderName, handler.nomadConfig.NomadToken)
	http.StripPrefix("/nomad", handler.nomadProxy).ServeHTTP(rw, request)

	return nil
}
