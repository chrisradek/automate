package v1_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/server"
	v1 "github.com/chef/automate/components/automate-cli/pkg/verifyserver/server/api/v1"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/services/statusservice"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/utils/fiberutils"
	"github.com/chef/automate/lib/logger"
	"github.com/gofiber/fiber"

	"github.com/stretchr/testify/assert"
)

func SetupMockStatusService(httpStatus int) *statusservice.MockStatusService {
	if httpStatus == 200 {
		return &statusservice.MockStatusService{
			GetServicesFunc: func() (*[]models.ServiceDetails, error) {
				return &[]models.ServiceDetails{
					{
						ServiceName: "deployment-service",
						Version:     "chef/deployment-service/0.1.0/20230502070345",
						Status:      constants.OK,
					},
				}, nil
			},
		}
	} else {
		return &statusservice.MockStatusService{
			GetServicesFunc: func() (*[]models.ServiceDetails, error) {
				return nil, fiber.NewError(fiber.StatusInternalServerError, "Some error occurred")
			},
		}
	}
}

func SetupDefaultHandlers(ss statusservice.IStatusService) (*fiber.App, error) {
	log, err := logger.NewLogger("text", "debug")
	if err != nil {
		return nil, err
	}
	fconf := &fiber.Settings{
		ServerHeader: server.SERVICE,
		ErrorHandler: fiberutils.CustomErrorHandler,
	}
	app := fiber.New(fconf)
	handler := v1.NewHandler(log).
		AddStatusService(ss)
	vs := &server.VerifyServer{
		Port:    server.DEFAULT_PORT,
		Log:     log,
		App:     app,
		Handler: handler,
	}
	vs.Setup(false)
	return vs.App, nil
}

func TestStatusAPI(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		expectedBody string
	}{
		{
			description:  "200:success status route",
			expectedCode: 200,
			expectedBody: "{\"status\":\"SUCCESS\",\"result\":{\"status\":\"OK\",\"services\":[{\"service_name\":\"deployment-service\",\"status\":\"OK\",\"version\":\"chef/deployment-service/0.1.0/20230502070345\"}]}}",
		},
		{
			description:  "500:InternalServerError status route",
			expectedCode: 500,
			expectedBody: "{\"status\":\"FAILED\",\"result\":null,\"error\":{\"code\":500,\"message\":\"Some error occurred\"}}",
		},
	}
	statusEndpoint := "/status"

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			app, err := SetupDefaultHandlers(SetupMockStatusService(test.expectedCode))
			assert.NoError(t, err)
			req := httptest.NewRequest("GET", statusEndpoint, nil)
			req.Header.Add("Content-Type", "application/json")
			res, err := app.Test(req, -1)
			assert.NoError(t, err)
			body, err := ioutil.ReadAll(res.Body)
			assert.NoError(t, err, test.description)
			assert.Contains(t, string(body), test.expectedBody)
			assert.Equal(t, res.StatusCode, test.expectedCode)
		})
	}
}
