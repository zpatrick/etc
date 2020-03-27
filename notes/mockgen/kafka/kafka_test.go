package kafka

import (
	"testing"

	"github.com/etc/notes/mockgen/kafka/kafkamock"
	"github.com/golang/mock/gomock"
)

func TestExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := kafkamock.NewMockClient(ctrl)
	client.EXPECT().
		WriteMessage(gomock.Any(), gomock.Any()).
		Return(nil)
}
