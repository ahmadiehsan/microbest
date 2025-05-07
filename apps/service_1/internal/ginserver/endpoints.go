package ginserver

import (
	"encoding/json"
	"net/http"
	"service_1/internal/pb/service2pb"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func (s *Server) hello(c *gin.Context) {
	log.Info().Msg("hello API")
	endpoints := []string{
		"/api",
		"/api/external-api-http",
		"/api/service-2-ping-http",
		"/api/service-2-event-http",
		"/api/service-2-echo-grpc",
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!", "end_points": endpoints})
}

func (s *Server) externalApiHttp(c *gin.Context) {
	log.Info().Msg("call external API")
	url := "https://httpbin.org/get"

	resp, err := otelhttp.Get(c.Request.Context(), url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close() // nolint:errcheck

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2PingHttp(c *gin.Context) {
	log.Info().Msg("call Service 2 ping API")
	url := "http://" + s.configs.Service2HttpAddress + "/api/ping/"

	resp, err := otelhttp.Get(c.Request.Context(), url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close() // nolint:errcheck

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2EventHttp(c *gin.Context) {
	log.Info().Msg("call Service 2 event API")
	url := "http://" + s.configs.Service2HttpAddress + "/api/event/"

	resp, err := otelhttp.Get(c.Request.Context(), url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close() // nolint:errcheck

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2EchoGrpc(c *gin.Context) {
	log.Info().Msg("call Service 2 echo RPC")

	resp, err := s.Service2RpcClient.Echo(c.Request.Context(), &service2pb.EchoRequest{Message: "hello from Service 1"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
}
