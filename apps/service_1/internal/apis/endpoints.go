package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"service_1/internal/pb/service2pb"
)

func (s *Server) hello(ginCtx *gin.Context) {
	log.Info().Msg("hello API")
	endpoints := []string{
		"/api",
		"/api/external-api-http",
		"/api/service-2-ping-http",
		"/api/service-2-event-http",
		"/api/service-2-echo-grpc",
	}
	ginCtx.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!", "end_points": endpoints})
}

func (s *Server) externalAPIHTTP(ginCtx *gin.Context) {
	log.Info().Msg("call external API")
	url := "https://httpbin.org/get"

	resp, err := otelhttp.Get(ginCtx.Request.Context(), url)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close() // nolint:errcheck // Popular projects don't check this error

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2PingHTTP(ginCtx *gin.Context) {
	log.Info().Msg("call Service 2 ping API")
	url := "http://" + s.configs.Service2HttpAddress + "/api/ping/"

	resp, err := otelhttp.Get(ginCtx.Request.Context(), url)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close() // nolint:errcheck // Popular projects don't check this error

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2EventHTTP(ginCtx *gin.Context) {
	log.Info().Msg("call Service 2 event API")
	url := "http://" + s.configs.Service2HttpAddress + "/api/event/"

	resp, err := otelhttp.Get(ginCtx.Request.Context(), url)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close() // nolint:errcheck // Popular projects don't check this error

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2EchoGrpc(ginCtx *gin.Context) {
	log.Info().Msg("call Service 2 echo RPC")

	resp, err := s.service2RpcClient.Echo(
		ginCtx.Request.Context(),
		&service2pb.EchoRequest{Message: "hello from Service 1"},
	)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "error: " + err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
}
