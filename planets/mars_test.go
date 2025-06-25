package planets_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"

	"planetTime/planets"
)

func TestMarsDateTotalSols(t *testing.T) {
	t.Run("OneRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1611-01-28T04:04:33Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 669)
	})
	t.Run("ThreeRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1615-03-25T18:35:08Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 2146)
	})
	t.Run("TenRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1628-06-02T16:20:28Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 6835)
	})
	t.Run("HundredRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1797-09-12T22:32:45Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 67009)
	})
	t.Run("TwoHundredRotations", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1985-05-14T21:08:18Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 133719)
	})
	t.Run("Now", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2025-04-14T08:48:29Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, marsTime.TotalSols, 147908)
	})
}

func TestMarsDateParams(t *testing.T) {
	t.Run("OneSols", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "1609-03-12T19:20:10Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
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
		marsTime := planets.NewMarsTime(&earthTime)
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
		marsTime := planets.NewMarsTime(&earthTime)
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
		marsTime := planets.NewMarsTime(&earthTime)
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
		marsTime := planets.NewMarsTime(&earthTime)
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
		marsTime := planets.NewMarsTime(&earthTime)
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R=%0M=%0S%'T%0V|%0L|%0F", "201=02=03T04|05|06"},
			{"rot %R m%M sol %S started %V vinquas %L layers %F fragments ago", "rot 201 m2 sol 3 started 4 vinquas 5 layers 6 fragments ago"},
			{"%R %NM %S%'th", "201 Dhanus 3th"},
			{"%R=W%0W=%WS", "201=W05=3"},
			{"rotation %R week %W %NS", "rotation 201 week 5 Martis"},

			{"%R=%0M=%0D%'T%0V|%0L|%0F", "201=02=03T04|05|06"},
			{"rot %R m%M sol %D started %V vinquas %L layers %F fragments ago", "rot 201 m2 sol 3 started 4 vinquas 5 layers 6 fragments ago"},
			{"%R %NM %D%'th", "201 Dhanus 3th"},
			{"%R=W%0W=%WD", "201=W05=3"},
			{"rotation %R week %W %ND", "rotation 201 week 5 Martis"},

			{"%R=%0M=%0D%Vu%0V11|%0L|%0F", "201=02=03A04|05|06"},
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
		marsTime := planets.NewMarsTime(&earthTime)
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R=%0M=%0S%'T%0V|%0L|%0F", "207=21=14T16|14|10"},
			{"rot %R m%M sol %S started %V vinquas %L layers %F fragments ago", "rot 207 m21 sol 14 started 16 vinquas 14 layers 10 fragments ago"},
			{"%R %NM %S%'th", "207 Libra 14th"},
			{"%R=W%0W=%WS", "207=W82=7"},
			{"rotation %R week %W %NS", "rotation 207 week 82 Saturni"},

			{"%R=%0M=%0D%'T%0V|%0L|%0F", "207=21=14T16|14|10"},
			{"rot %R m%M sol %D started %V vinquas %L layers %F fragments ago", "rot 207 m21 sol 14 started 16 vinquas 14 layers 10 fragments ago"},
			{"%R %NM %D%'th", "207 Libra 14th"},
			{"%R=W%0W=%WD", "207=W82=7"},
			{"rotation %R week %W %ND", "rotation 207 week 82 Saturni"},

			{"%R=%0M=%0D%Vu%0V11|%0L|%0F", "207=21=14P04|14|10"},
			{"%R=%0M=%0D %Vl.m. %V11|%0L|%0F", "207=21=14 p.m. 4|14|10"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				result := marsTime.Format(tc.layout)
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("Errors", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2025-04-15T06:36:43Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		tests := []string{
			// only digit will never be used
			"%0 %1 %2 %3 %4 %5 %6 %7 %8 %9",

			// break compatability later; might need update
			"the %0A mistake",
		}
		for _, tc := range tests {
			t.Run(tc, func(t *testing.T) {
				result := marsTime.Format(tc)
				fmt.Println(result)
				assert.Equal(t, strings.HasPrefix(result, "error"), true)
			})
		}
	})
}

func TestParseMarsTime(t *testing.T) {
	t.Run("BasicTest", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R=%0M=%0S%'T%0V|%0L|%0F", "201=02=03T04|05|06"},
			{"rot %R m%M sol %S started %V vinquas %L layers %F fragments ago", "rot 201 m2 sol 3 started 4 vinquas 5 layers 6 fragments ago"},
			{"%R %NM %oS", "201 Dhanus 3rd"},
			{"%R %NM %S%'th", "201 Dhanus 3th"},
			{"%R=W%0W=%WS", "201=W05=3"},
			{"rotation %R week %W %NS", "rotation 201 week 5 Martis"},
			{"rotation%%%R week %W %NS", "rotation%201 week 5 Martis"},

			{"%R=%0M=%0D%'T%0V|%0L|%0F", "201=02=03T04|05|06"},
			{"rot %R m%M sol %D started %V vinquas %L layers %F fragments ago", "rot 201 m2 sol 3 started 4 vinquas 5 layers 6 fragments ago"},
			{"%R %NM %oD", "201 Dhanus 3rd"},
			{"%R %NM %D%'th", "201 Dhanus 3th"},
			{"%R=W%0W=%WD", "201=W05=3"},
			{"rotation %R week %W %ND", "rotation 201 week 5 Martis"},
			{"rotation%%%R week %W %ND", "rotation%201 week 5 Martis"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				result := marsTime.Format(tc.layout)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("NanoFragments", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R=%0M=%0S%'T%0V|%0L|%0F.%f", "201=02=03T04|05|06.5"},
			{"%R=%0M=%0S%'T%0V|%0L|%0F.%f0", "201=02=03T04|05|06.500000000"},
			{"%R=%0M=%0S%'T%0V|%0L|%0F.%f", "201=02=03T04|05|06.000000005"},
			{"%R=%0M=%0S%'T%0V|%0L|%0F.%f0", "201=02=03T04|05|06.000000005"},
			{"%R=%0M=%0S%'T%0V|%0L|%0F.%f", "201=02=03T04|05|06.712563515"},
			{"%R=%0M=%0S%'T%0V|%0L|%0F.%f0", "201=02=03T04|05|06.712563515"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				result := marsTime.Format(tc.layout)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("SolSaturni", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R=%0M=%0S%'T%0V|%0L|%0F", "207=21=14T16|14|10"},
			{"rot %R m%M sol %S started %V vinquas %L layers %F fragments ago", "rot 207 m21 sol 14 started 16 vinquas 14 layers 10 fragments ago"},
			{"%R %NM %S%'th", "207 Libra 14th"},
			{"%R %NM %oS", "207 Libra 14th"},
			{"%R=W%0W=%WS", "207=W82=7"},
			{"rotation %R week %W %NS", "rotation 207 week 82 Saturni"},

			{"%R %nM %_S", "207 Lib 14"},
			{"%R %nM %_S", "207 Lib  7"},
			{"rotation %R week %W %nS", "rotation 207 week 82 Sat"},

			{"%R=%0M=%0D%'T%0V|%0L|%0F", "207=21=14T16|14|10"},
			{"rot %R m%M sol %D started %V vinquas %L layers %F fragments ago", "rot 207 m21 sol 14 started 16 vinquas 14 layers 10 fragments ago"},
			{"%R %NM %D%'th", "207 Libra 14th"},
			{"%R %NM %oD", "207 Libra 14th"},
			{"%R=W%0W=%WD", "207=W82=7"},
			{"rotation %R week %W %ND", "rotation 207 week 82 Saturni"},

			{"%R %nM %_D", "207 Lib 14"},
			{"%R %nM %_D", "207 Lib  7"},
			{"rotation %R week %W %nD", "rotation 207 week 82 Sat"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				result := marsTime.Format(tc.layout)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("SplitVinquaInfo", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			{"%R=%0M=%0S%Vu%0V11|%0L|%0F", "207=21=14A04|14|10"},
			{"%R=%0M=%0S %Vl.m. %V11|%0L|%0F", "207=21=14 a.m. 4|14|10"},
			{"%R=%0M=%0S %Vl.m. %V12|%0L|%0F", "207=21=14 a.m. 4|14|10"},

			{"%R=%0M=%0D%Vu%0V11|%0L|%0F", "207=21=14P04|14|10"},
			{"%R=%0M=%0D %Vl.m. %V11|%0L|%0F", "207=21=14 p.m. 4|14|10"},
			{"%R=%0M=%0D %Vl.m. %V12|%0L|%0F", "207=21=14 p.m. 4|14|10"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				result := marsTime.Format(tc.layout)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})

	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			// only digit will never be used
			{"%0 %1 %2 %3 %4 %5 %6 %7 %8 %9", "0 1 2 3 4 5 6 7 8 9"},

			// break compatability later; might need update
			{"the %0A mistake", "the AMPLE mistake"},

			// Wrong data
			{"%R %nM %D", "225 MONTH 14"},
			{"%R %NM %D", "225 MONTH 14"},
			{"%R %NM %oD", "225 Kumbha DAY"},
			{"rotation %R week %W %nD", "rotation 207 week 82 SOL"},
			{"rotation %R week %W %ND", "rotation 207 week 82 WEEKSOL"},
			{"rotation %R", "rotation 207"},
			{"%R %M %D", "207 Libra 14"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				_, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				assert.NotEqual(t, err, nil)
				fmt.Println(err)
			})
		}
	})
	t.Run("LayoutPrintErrors", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			{"!rotation %R week %W %NS", "rotation 207 week 82 Saturni"},
			{"%%rotation %R week %W %NS", "rotation 207 week 82 Saturni"},
			{"rotation %R week %W %NS %%", "rotation 207 week 82 Saturni"},
			{"rotation %R week %W %NS %%", "rotation 207 week 82 Saturni !"},
			{"rotation %R week %W %NS", "rotation 207 week 82 Saturni !"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				_, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				assert.NotEqual(t, err, nil)
				fmt.Println(err)
			})
		}
	})
	t.Run("LayoutTokenErrors", func(t *testing.T) {
		tests := []struct {
			layout   string
			expected string
		}{
			{"%E rotation %R week %W %NS", "rotation 207 week 82 Saturni"},
			{"%ERR rotation %R week %W %NS", "rotation 207 week 82 Saturni"},

			// there support vinquas 00 thru 11 thus without AM/PM specifier leaving vinquas 12 thru 23 impossible
			{"rotation %R week %W %NS %V12", "rotation 207 week 82 Saturni 1"},
			{"rotation %R week %W %NS %V12", "rotation 207 week 82 Saturni 11"},
			{"rotation %R week %W %NS %V12", "rotation 207 week 82 Saturni 12"},
			{"rotation %R week %W %NS %V12", "rotation 207 week 82 Saturni 13"},
			{"rotation %R week %W %NS %0V12", "rotation 207 week 82 Saturni 01"},
			{"rotation %R week %W %NS %0V12", "rotation 207 week 82 Saturni 11"},
			{"rotation %R week %W %NS %0V12", "rotation 207 week 82 Saturni 12"},
			{"rotation %R week %W %NS %0V12", "rotation 207 week 82 Saturni 13"},
			{"rotation %R week %W %NS %_V12", "rotation 207 week 82 Saturni  1"},
			{"rotation %R week %W %NS %_V12", "rotation 207 week 82 Saturni 11"},
			{"rotation %R week %W %NS %_V12", "rotation 207 week 82 Saturni 12"},
			{"rotation %R week %W %NS %_V12", "rotation 207 week 82 Saturni 13"},

			// invalid '%Vl' specifier
			{"%R=%0M=%0D %Vl.m. %V11|%0L|%0F", "207=21=14 m.m. 4|14|10"},
		}
		for _, tc := range tests {
			t.Run(tc.layout, func(t *testing.T) {
				_, err := planets.MarsTime{}.Parse(tc.layout, tc.expected)
				assert.NotEqual(t, err, nil)
				fmt.Println(err)
			})
		}
	})
}

func TestMarsTimeExamples(t *testing.T) {
	t.Run("BasicTest", func(t *testing.T) {
		tests := []struct {
			example  string
			expected string
		}{
			{"203=04=05T00|01|02", "201=02=03T04|05|06"},
			{"rot 203 m4 sol 5 started 0 vinquas 1 layers 2 fragments ago", "rot 201 m2 sol 3 started 4 vinquas 5 layers 6 fragments ago"},
			{"203 Makara 5th", "201 Dhanus 3rd"},
			{"203=W13=5", "201=W05=3"},
			{"rotation 203 week !13 Jovis", "rotation 201 week 5 Martis"},
			{"rotation%203 week !13 Jovis", "rotation%201 week 5 Martis"},
		}
		for _, tc := range tests {
			t.Run(tc.example, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.ParseExample(tc.example, tc.expected)
				result := marsTime.FormatExample(tc.example)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("NanoFragments", func(t *testing.T) {
		tests := []struct {
			example  string
			expected string
		}{
			{"203=04=05T00|01|02.3", "201=02=03T04|05|06.5"},
			{"203=04=05T00|01|02.300000000", "201=02=03T04|05|06.500000000"},
			{"203=04=05T00|01|02.3", "201=02=03T04|05|06.000000005"},
			{"203=04=05T00|01|02.300000000", "201=02=03T04|05|06.000000005"},
			{"203=04=05T00|01|02.3", "201=02=03T04|05|06.712563515"},
			{"203=04=05T00|01|02.300000000", "201=02=03T04|05|06.712563515"},
		}
		for _, tc := range tests {
			t.Run(tc.example, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.ParseExample(tc.example, tc.expected)
				result := marsTime.FormatExample(tc.example)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("SolSaturni", func(t *testing.T) {
		tests := []struct {
			example  string
			expected string
		}{
			{"203=04=05T00|01|02", "207=21=14T16|14|10"},
			{"rot 203 m04 sol 05 started 00 vinquas 01 layers 02 fragments ago", "rot 207 m21 sol 14 started 16 vinquas 14 layers 10 fragments ago"},
			{"203 Makara 05th", "207 Libra 14th"},
			{"203=W13=5", "207=W82=7"},
			{"rotation 203 week 13 Jovis", "rotation 207 week 82 Saturni"},

			{"203 Mak 5", "207 Lib 14"},
			{"203 Mak _5", "207 Lib  7"},
			{"rotation 203 week 13 Jov", "rotation 207 week 82 Sat"},
		}
		for _, tc := range tests {
			t.Run(tc.example, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.ParseExample(tc.example, tc.expected)
				result := marsTime.FormatExample(tc.example)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
	t.Run("SplitVinquaInfo", func(t *testing.T) {
		tests := []struct {
			example  string
			expected string
		}{
			{"203=04=05A00|01|02", "207=21=14A04|14|10"},
			{"203=04=05 a.m. 0|01|02", "207=21=14 a.m. 4|14|10"},
			{"203=04=05 a.m. 0|01|02", "207=21=14 a.m. 4|14|10"},

			{"203=04=05 a.m. 0|01|02", "207=21=14 a.m. 0|14|10"},
			{"203=04=05 a.m. 00|01|02", "207=21=14 a.m. 00|14|10"},
			{"203=04=05 a.m. 12|01|02", "207=21=14 a.m. 12|14|10"},
			{"203=04=05 a.m. 12|01|02", "207=21=14 a.m. 4|14|10"},
		}
		for _, tc := range tests {
			t.Run(tc.example, func(t *testing.T) {
				marsTime, err := planets.MarsTime{}.ParseExample(tc.example, tc.expected)
				result := marsTime.FormatExample(tc.example)
				if err != nil {
					fmt.Println(err)
				}
				assert.Equal(t, result, tc.expected)
			})
		}
	})
}

func TestMarsDateTime(t *testing.T) {
	t.Run("Now", func(t *testing.T) {
		earthTime, err := time.Parse(time.RFC3339, "2025-04-14T08:48:29Z")
		assert.Equal(t, err, nil)
		marsTime := planets.NewMarsTime(&earthTime)
		assert.Equal(t, earthTime, marsTime.Time())
	})
}
