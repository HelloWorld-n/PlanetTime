package planets

import (
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestMarsDateTotalSols(t *testing.T) {
	t.Run("OneRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1611-01-28T04:04:33Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 669)
	})
	t.Run("ThreeRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1615-03-25T18:35:08Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 2146)
	})
	t.Run("TenRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1628-06-02T16:20:28Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 6835)
	})
	t.Run("HundredRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1797-09-12T22:32:45Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 67009)
	})
	t.Run("TwoHundredRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1985-05-14T21:08:18Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 133719)
	})
	t.Run("Now", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2025-04-14T08:48:29Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 147908)
	})
}

func TestMarsDateParams(t *testing.T) {
	t.Run("OneSols", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1609-03-12T19:20:10Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		rotation, month, sol, vinqua, layer, fragment, _ := marsTime.Params()
		assert.Equal(t, rotation, 0)
		assert.Equal(t, month, 1)
		assert.Equal(t, sol, 2)
		assert.Equal(t, vinqua, 0)
		assert.Equal(t, layer, 0)
		assert.Equal(t, fragment, 0)
	})
	t.Run("OneRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1611-01-28T04:04:33Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		rotation, month, sol, vinqua, layer, fragment, _ := marsTime.Params()
		assert.Equal(t, rotation, 1)
		assert.Equal(t, month, 1)
		assert.Equal(t, sol, 1)
		assert.Equal(t, vinqua, 0)
		assert.Equal(t, layer, 0)
		assert.Equal(t, fragment, 0)
	})
	t.Run("OneRotationsAndOneMonth", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1611-02-25T22:33:00Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		rotation, month, sol, vinqua, layer, fragment, _ := marsTime.Params()
		assert.Equal(t, rotation, 1)
		assert.Equal(t, month, 2)
		assert.Equal(t, sol, 1)
		assert.Equal(t, vinqua, 0)
		assert.Equal(t, layer, 0)
		assert.Equal(t, fragment, 0)
	})
	t.Run("NewSol", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2025-04-13T22:53:57Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		rotation, month, sol, vinqua, layer, fragment, _ := marsTime.Params()
		assert.Equal(t, rotation, 221)
		assert.Equal(t, month, 6)
		assert.Equal(t, sol, 10)
		assert.Equal(t, vinqua, 0)
		assert.Equal(t, layer, 0)
		assert.Equal(t, fragment, 0)
	})
	t.Run("Now", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2025-04-14T10:29:56Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		rotation, month, sol, vinqua, layer, fragment, _ := marsTime.Params()
		assert.Equal(t, rotation, 221)
		assert.Equal(t, month, 6)
		assert.Equal(t, sol, 10)
		assert.Equal(t, vinqua, 11)
		assert.Equal(t, layer, 17)
		assert.Equal(t, fragment, 22)
	})
}
