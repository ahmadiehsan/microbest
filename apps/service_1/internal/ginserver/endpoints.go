package ginserver

import (
	"context"
	"encoding/json"
	"net/http"
	"service_1/internal/helpers"
	"service_1/internal/pb/service2pb"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func hello(c *gin.Context) {
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

func externalApiHttp(c *gin.Context) {
	log.Info().Msg("call external API")
	url := "https://httpbin.org/get"
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := ctxhttp.Get(c.Request.Context(), client, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func service2PingHttp(c *gin.Context) {
	log.Info().Msg("call Service 2 ping API")
	configs := helpers.GetConfigs()
	url := "http://" + configs.Service2HttpAddress + "/api/ping/"
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := ctxhttp.Get(c.Request.Context(), client, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func service2EventHttp(c *gin.Context) {
	log.Info().Msg("call Service 2 event API")
	configs := helpers.GetConfigs()
	url := "http://" + configs.Service2HttpAddress + "/api/event/"
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := ctxhttp.Get(c.Request.Context(), client, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": resp.StatusCode, "content": result})
}

func service2EchoGrpc(c *gin.Context) {
	log.Info().Msg("call Service 2 echo RPC")
	configs := helpers.GetConfigs()

	conn, err := grpc.NewClient(configs.Service2GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
