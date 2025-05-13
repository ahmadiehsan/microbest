package apis

import (
	"encoding/json"
	"net/http"

	"github.com/agoda-com/opentelemetry-go/otelzap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"service_1/internal/pb/service2pb"
)

func (a *App) hello(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	otelzap.Ctx(ctx).Info("hello API")
	endpoints := []string{
		"/api",
		"/api/external-api-http",
		"/api/service-2-ping-http",
		"/api/service-2-event-http",
		"/api/service-2-echo-grpc",
	}
	ginCtx.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!", "end_points": endpoints})
}

func (a *App) externalAPIHTTP(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	otelzap.Ctx(ctx).Info("call external API")
	url := "https://httpbin.org/get"

	resp, err := otelhttp.Get(ctx, url)
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

func (a *App) service2PingHTTP(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	otelzap.Ctx(ctx).Info("call Service 2 ping API")
	url := "http://" + a.configs.Service2HttpAddress + "/api/ping/"

	resp, err := otelhttp.Get(ctx, url)
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

func (a *App) service2EventHTTP(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	otelzap.Ctx(ctx).Info("call Service 2 event API")
	url := "http://" + a.configs.Service2HttpAddress + "/api/event/"

	resp, err := otelhttp.Get(ctx, url)
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

func (a *App) service2EchoGrpc(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	otelzap.Ctx(ctx).Info("call Service 2 echo RPC")

	resp, err := a.service2RpcClient.Echo(
		ctx,
		&service2pb.EchoRequest{Message: "hello from Service 1"},
	)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "error: " + err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": resp.GetMessage()})
}
