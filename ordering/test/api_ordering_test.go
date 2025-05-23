/*
Ordering

Testing OrderingAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ordering

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/leaseweb/leaseweb-go-sdk/ordering"
)

func Test_ordering_OrderingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrderingAPIService GetDedicatedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dedicatedServerId string

		resp, httpRes, err := apiClient.OrderingAPI.GetDedicatedServer(context.Background(), dedicatedServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderingAPIService GetDedicatedServerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrderingAPI.GetDedicatedServerList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderingAPIService OrderDedicatedServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dedicatedServerId string

		resp, httpRes, err := apiClient.OrderingAPI.OrderDedicatedServer(context.Background(), dedicatedServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
