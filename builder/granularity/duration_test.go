package granularity

import (
	"testing"
	"time"

	"github.com/raakasf/go-druid/builder/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewDuration(t *testing.T) {
	d := NewDuration()
	d.SetDuration(time.Minute)
	d.SetOrigin(time.Unix(1613779200, 0))

	expected := []byte(`{"type":"duration","duration":60000000000,"origin":"2021-02-19T16:00:00-08:00"}`)

	built, err := Load(expected)
	assert.Nil(t, err)

	testutil.Compare(t, expected, d, built)
}
