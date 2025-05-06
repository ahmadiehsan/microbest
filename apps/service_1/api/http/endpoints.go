package http

import (
	"context"
	"encoding/json"
	"net/http"
	"service_1/internal/pb/service2pb"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2PingHttp(c *gin.Context) {
	log.Info().Msg("call Service 2 ping API")
	url := "http://" + s.Configs.Service2HttpAddress + "/api/ping/"
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2EventHttp(c *gin.Context) {
	log.Info().Msg("call Service 2 event API")
	url := "http://" + s.Configs.Service2HttpAddress + "/api/event/"
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func (s *Server) service2EchoGrpc(c *gin.Context) {
	log.Info().Msg("call Service 2 echo RPC")

	conn, err := grpc.NewClient(s.Configs.Service2GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error: " + err.Error()})
		return
	}
	defer conn.Close()

	client := service2pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Echo(ctx, &service2pb.EchoRequest{Message: "hello from Service 1"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
}
