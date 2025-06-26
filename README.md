# MarsTime

**MarsTime** is a Go library for converting Earth time to Mars time using a custom calendar system. It includes rotation-based tracking, formatting, parsing, and detailed decomposition of Mars time components such as months, sols, vinquas, layers, and fragments.

## Features

- Convert Earth time (`time.Time`) to Mars time.
- Retrieve total sols since the Mars epoch.
- Decompose Mars time into:
  - Rotation
  - Month
  - Sol
  - Vinqua
  - Layer
  - Fragment
- Format Mars time using a flexible token-based layout syntax.
- Parse Mars time from formatted strings.

## Format Tokens

| Token   | Description                                                   |
|---------|---------------------------------------------------------------|
| `%R`    | Rotation (year)                                               |
| `%M`    | Month (month)                                                 |
| `%NM`   | Full Month name                                               |
| `%nM`   | Abbreviated Month name                                        |
| `%0M`   | Zero Padded Month                                             |
| `%W`    | Week number                                                   |
| `%S`    | Sol (day)                                                     |
| `%oS`   | Ordinal Sol (e.g. 14th)                                       |
| `%_S`   | Space Padded Sol                                              |
| `%0S`   | Zero Padded Sol                                               |
| `%WS`   | Weekday number (1–7)                                          |
| `%NS`   | Full Sol (weekday) name                                       |
| `%nS`   | Abbreviated Sol name                                          |
| `%V`    | Vinqua (hour)                                                 |
| `%0V`   | Zero Padded Vinqua                                            |
| `%_V`   | Underscore Padded Vinqua                                      |
| `%Vl`   | `a` if Vinqua % 12 == 0 else `p`                              |
| `%Vu`   | `A` if Vinqua % 12 == 0 else `P`                              |
| `%V11`  | Vinqua % 12                                                   |
| `%0V11` | Zero Padded Vinqua % 12                                       |
| `%_V11` | Underscore Padded Vinqua % 12                                 |
| `%V12`  | Vinqua % 12 unless it is 0 in that case 12                    |
| `%0V12` | Zero Padded Vinqua % 12 unless it is 0 in that case 12        |
| `%_V12` | Underscore Padded Vinqua % 12 unless it is 0 in that case 12  |
| `%L`    | Layer (minute)                                                |
| `%0L`   | Zero Padded Layer                                             |
| `%F`    | Fragment (second)                                             |
| `%0F`   | Zero Padded Fragment                                          |
| `%f`    | NanoFragment                                                  |
| `%f0`   | Zero Padded NanoFragment                                      |
| `%%`    | Literal `%` character                                         |
| `%'`    | Used to split tokens from text                                |


## Example

```go
package main

import (
	"fmt"
	"planetTime/planets"
	"time"
)

func main() {
	earthTime := time.Now().UTC()
	marsTime := planets.NewMarsTime(&earthTime)
	fmt.Println(marsTime.Format("rot %R %NM %oD started %V vinquas %L layers %F fragments ago"))
	// example output: "rot 221 Mina 26th started 14 vinquas 35 layers 0 fragments ago"

	// specific time
	customTime, _ := planets.MarsTime{}.Parse("%R=%0M=%0S%'T%0V|%0L|%0F", "221=08=26T14|35|00")
	fmt.Println(customTime.Format("rot %R %NM %oD started %V vinquas %L layers %F fragments ago"))
	// output: "rot 221 Mina 26th started 14 vinquas 35 layers 0 fragments ago"

	fmt.Println(customTime.Format("%NS, %R %NM %oD @ %0V|%0L|%0F"))
	// output: "Jovis, 221 Mina 26th @ 14|35|00"

	customTime, _ := planets.MarsTime{}.Parse("%R=%0M=%0S%'T%0V|%0L|%0F", "221=08=03T14|35|00")
	fmt.Println(customTime.Format("%NS, %R %NM %oD @ %0V|%0L|%0F"))
	// output: "Martis, 221 Mina 3rd @ 14|35|00"
}
```

`planetTime.planets.MarsTime` also has methods `ParseExample`, `FormatExample` that use examples instead of layouts.
Those methods use 203=04=05T00|01|02 as basic date.
See file `./main.go` for example.
