package chrono

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromStdTime(t *testing.T) {
	tm, err := time.Parse(time.RFC3339, "2012-03-04T05:06:07.012345+09:00")
	require.NoError(t, err)

	got := FromStdTime(tm).Unwrap()
	assert.Equal(t, int32(2012), got.Year())
	assert.Equal(t, uint32(3), got.Month())
	assert.Equal(t, uint32(4), got.Day())
	assert.Equal(t, uint32(5), got.Hour())
	assert.Equal(t, uint32(6), got.Minute())
	assert.Equal(t, uint32(7), got.Second())
	assert.Equal(t, uint32(12345000), got.Nanosecond())
}
