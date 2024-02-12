package mockspanner_test

import (
	"context"
	"testing"

	mockspanner "github.com/arhea/go-mock-spanner"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	config := mockspanner.ClientConfig{
		T: t,
	}

	mock, err := mockspanner.NewClient(ctx, config)

	if err != nil {
		t.Fatalf("creating the client: %v", err)
		return
	}

	// close the mock
	defer mock.Close(ctx)

	assert.NotNil(t, mock.Client(), "client should not be nil")
}
