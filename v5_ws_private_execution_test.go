package bybit

import (
	"encoding/json"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5WebsocketPrivate_Execution(t *testing.T) {
	respBody := V5WebsocketPrivateExecutionResponse{
		Topic:        "Execution",
		ID:           "75d86e42f18b23b9ad2c1f10eaffa8bb:18483ff242aca593:0:01",
		CreationTime: 1677226839837,
		Data: []V5WebsocketPrivateExecutionData{
			{
				Category:        "linear",
				Symbol:          "XRPUSDT",
				ExecFee:         "0.005061",
				ExecID:          "7e2ae69c-4edf-5800-a352-893d52b446aa",
				ExecPrice:       "0.3374",
				ExecQty:         "25",
				ExecType:        "Trade",
				ExecValue:       "8.435",
				IsMaker:         false,
				FeeRate:         "0.0006",
				TradeIv:         "",
				MarkIv:          "",
				BlockTradeID:    "",
				MarkPrice:       "0.3391",
				IndexPrice:      "",
				UnderlyingPrice: "",
				LeavesQty:       "0",
				OrderID:         "f6e324ff-99c2-4e89-9739-3086e47f9381",
				OrderLinkID:     "",
				OrderPrice:      "0.3207",
				OrderQty:        "25",
				OrderType:       "Market",
				StopOrderType:   "UNKNOWN",
				Side:            "Sell",
				ExecTime:        "1672364174443",
				IsLeverage:      "0",
				ClosedSize:      "",
				Seq:             4688002127,
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewWebsocketServer(
		testhelper.WithWebsocketHandlerOption(V5WebsocketPrivatePath, bytesBody),
	)
	defer teardown()

	wsClient := NewTestWebsocketClient().
		WithBaseURL(server.URL).
		WithAuth("test", "test")

	svc, err := wsClient.V5().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	{
		_, err := svc.SubscribeExecution(func(response V5WebsocketPrivateExecutionResponse) error {
			assert.Equal(t, respBody, response)
			return nil
		})
		require.NoError(t, err)
	}

	assert.NoError(t, svc.Run())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
