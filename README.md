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

| Token   | Description                     |
|---------|---------------------------------|
| `%R`    | Rotation (year)                 |
| `%M`    | Month (month)                   |
| `%S`    | Sol (day)                       |
| `%V`    | Vinqua (hour)                   |
| `%L`    | Layer (minute)                  |
| `%F`    | Fragment (second)               |
| `%NM`   | Full month name                 |
| `%nM`   | Abbreviated month name          |
| `%oS`   | Ordinal Sol (e.g. 14th)         |
| `%_S`   | Space Padded Sol                |
| `%0M`   | Zero Padded Month               |
| `%0S`   | Zero Padded Sol                 |
| `%0V`   | Zero Padded Vinqua              |
| `%0L`   | Zero Padded Layer               |
| `%0F`   | Zero Padded Fragment            |
| `%W`    | Week number                     |
| `%WS`   | Weekday number (1–7)            |
| `%NS`   | Full Sol (weekday) name         |
| `%nS`   | Abbreviated Sol name            |
| `%%`    | Literal `%` character           |
| `%'`    | Used to split tokens from text  |


## Example

```go
earthTime, _ := time.Parse(time.RFC3339, "2025-04-14T10:29:56Z")
marsTime := planets.NewMarsTime(&earthTime)
rotation, month, sol, vinqua, layer, fragment, _ := marsTime.Params()
fmt.Printf("Rotation: %d, Month: %d, Sol: %d\n", rotation, month, sol)
fmt.Println("Formatted:", marsTime.Format("%R-%0M-%0D"))
```
