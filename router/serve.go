package router

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/util"
)

func ListenAndServe(addr string) {
	http.ListenAndServe(addr, router.mux)
}

func ListenAndServeAsLambda() {
	lambda.Start(func(proxyReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		reader := strings.NewReader(proxyReq.Body)
		req, err := http.NewRequest(proxyReq.HTTPMethod, proxyReq.Path, reader)
		if err != nil {
			logger.Error("failed to handle Lambda's APIGatewayProxyRequest due to error (%s)", err.Error())
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
		}

		responseRecorder := httptest.NewRecorder()
		router.mux.ServeHTTP(responseRecorder, req)

		singleValHeaders, multiValHeaders := util.GetHeaderMaps(responseRecorder)

		return events.APIGatewayProxyResponse{
			StatusCode:        responseRecorder.Code,
			Headers:           singleValHeaders,
			MultiValueHeaders: multiValHeaders,
			Body:              responseRecorder.Body.String(),
		}, nil
	})
}
