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

func TestMarsTimeFormat(t *testing.T) {
	t.Run("BasicTest", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1987-05-02T05:52:09Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R-%0M-%0D%'T%0V:%0L:%0F", "201-02-03T04:05:06"},
			{"%R=%0M=%0D%'T%0V|%0L|%0F", "201=02=03T04|05|06"},
			{"rot %R m%M sol %D started %V vinquas %L layers %F fragments ago", "rot 201 m2 sol 3 started 4 vinquas 5 layers 6 fragments ago"},
			{"%R %NM %D%'th", "201 Dhanus 3th"},
			{"%R=W%0W-%WS", "201=W05-3"},
			{"rotation %R week %W %NS", "rotation 201 week 5 Martis"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				result := marsTime.Format(tc.layout)
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("SolSaturni", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2000-02-19T21:03:32Z")
		assert.Equal(t, err, nil)
		marsTime := NewMarsTime(&earthTime)
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R-%0M-%0D%'T%0V:%0L:%0F", "207-21-14T16:14:10"},
			{"%R=%0M=%0D%'T%0V|%0L|%0F", "207=21=14T16|14|10"},
			{"rot %R m%M sol %D started %V vinquas %L layers %F fragments ago", "rot 207 m21 sol 14 started 16 vinquas 14 layers 10 fragments ago"},
			{"%R %NM %D%'th", "207 Libra 14th"},
			{"%R=W%0W-%WS", "207=W82-7"},
			{"rotation %R week %W %NS", "rotation 207 week 82 Saturni"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				result := marsTime.Format(tc.layout)
				assert.Equal(t, result, tc.expected)
			})
		}
	})
}
