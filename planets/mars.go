package planets

import (
	"fmt"
	"strings"
	"time"

	"planetTime/format"
)

type MarsTime struct {
	TotalSols            int
	DurationOfCurrentSol time.Duration
}

const (
	Fragment = time.Nanosecond * 1_027_491_251
	Layer    = Fragment * 60
	Vinqua   = Layer * 60
	Sol      = Vinqua * 24
)

// Functions returning the slices, instead of global variables, to emulate immutability

// Deprecated: use MarsTime{}.LongWeekSolNames() instead
func MarsLongWeekSolNames() []string {
	return MarsTime{}.LongWeekSolNames()
}

func (t MarsTime) LongWeekSolNames() []string {
	return []string{
		"Solis",
		"Lunae",
		"Martis",
		"Mercurii",
		"Jovis",
		"Veneris",
		"Saturni",
	}
}

// Deprecated: use MarsTime{}.ShortWeekSolNames() instead
func MarsShortWeekSolNames() []string {
	return MarsTime{}.ShortWeekSolNames()
}

func (t MarsTime) ShortWeekSolNames() []string {
	return []string{
		"Sol", // Solis
		"Lun", // Lunae
		"Mar", // Martis
		"Mer", // Mercurii
		"Jov", // Jovis
		"Ven", // Veneris
		"Sat", // Saturni
	}
}

// Deprecated: use MarsTime{}.LongMonthNames() instead
func MarsLongMonthNames() []string {
	return MarsTime{}.LongMonthNames()
}

func (t MarsTime) LongMonthNames() []string {
	return []string{
		"Sagittarius",
		"Dhanus",
		"Capricornus",
		"Makara",
		"Aquarius",
		"Kumbha",
		"Pisces",
		"Mina",
		"Aries",
		"Mesha",
		"Taurus",
		"Vrishabha",
		"Gemini",
		"Mithuna",
		"Cancer",
		"Karka",
		"Leo",
		"Simha",
		"Virgo",
		"Kanya",
		"Libra",
		"Tula",
		"Scorpio",
		"Vrishika",
	}
}

// Deprecated: use MarsTime{}.ShortMonthNames() instead
func MarsShortMonthNames() []string {
	return MarsTime{}.ShortMonthNames()
}

func (t MarsTime) ShortMonthNames() []string {
	return []string{
		"Sag", // Sagittarius
		"Dha", // Dhanus
		"Cap", // Capricornus
		"Mak", // Makara
		"Aqu", // Aquarius
		"Kum", // Kumbha
		"Pis", // Pisces
		"Min", // Mina
		"Ari", // Aries
		"Mes", // Mesha
		"Tau", // Taurus
		"Vrb", // Vrishabha (chosen to differentiate from Vrishika)
		"Gem", // Gemini
		"Mit", // Mithuna
		"Can", // Cancer
		"Kar", // Karka
		"Leo", // Leo
		"Sim", // Simha
		"Vir", // Virgo
		"Kan", // Kanya
		"Lib", // Libra
		"Tul", // Tula
		"Sco", // Scorpio
		"Vrk", // Vrishika (chosen to differentiate from Vrishabha)
	}
}

// Deprecated: use MarsTime{}.RotationHasLeapVrishika28th() instead
func RotationHasLeapVrishika28th(rotation int) (res bool) {
	return MarsTime{}.RotationHasLeapVrishika28th(rotation)
}

func (t MarsTime) RotationHasLeapVrishika28th(rotation int) (res bool) {
	if rotation%500 == 0 {
		return true
	}
	if rotation%100 == 0 {
		return false
	}
	if rotation%10 == 0 {
		return true
	}
	if rotation%2 == 0 {
		return false
	}
	return true
}

// Deprecated: use MarsTime{}.SolsInRotation(rotation int) instead
func SolsInRotation(rotation int) (sols int) {
	return MarsTime{}.SolsInRotation(rotation)
}

func (t MarsTime) SolsInRotation(rotation int) (sols int) {
	sols = 668
	if t.RotationHasLeapVrishika28th(rotation) {
		sols += 1
	}
	return
}

// Deprecated: use(t MarsTime) SolsInMonth(month int) instead
func SolsInMonth(month int) (sols int) {
	return MarsTime{}.SolsInMonth(month)
}

func (t MarsTime) SolsInMonth(month int) (sols int) {
	if month%6 == 0 {
		return 27
	}
	return 28
}

func NewMarsTime(t *time.Time) (res MarsTime) {
	ref := time.Date(1609, time.March, 11, 18, 40, 34, 0, time.UTC)

	duration := t.Sub(ref)

	// t.Sub is not able to go more than 2562047h47m16.854775807s
	largeDuration := 100000 * Sol
	for duration > largeDuration {
		ref = ref.Add(largeDuration)
		res.TotalSols += int(largeDuration / Sol)
		duration = t.Sub(ref)
	}

	res = MarsTime{
		TotalSols:            res.TotalSols + int(duration/Sol),
		DurationOfCurrentSol: duration % Sol,
	}

	return res
}

func (t MarsTime) Params() (rotation int, month int, sol int, vinqua int, layer int, fragment int, rem int) {
	month = 1
	sol = t.TotalSols
	for sol >= t.SolsInRotation(rotation) {
		sol -= t.SolsInRotation(rotation)
		rotation += 1
	}
	for sol >= t.SolsInMonth(month) {
		sol -= t.SolsInMonth(month)
		month += 1
	}
	sol += 1

	vinqua = int(t.DurationOfCurrentSol / Vinqua)
	layer = int(t.DurationOfCurrentSol % Vinqua / Layer)
	fragment = int(t.DurationOfCurrentSol % Layer / Fragment)
	rem = int(t.DurationOfCurrentSol%Fragment) * int(time.Second) / int(Fragment)
	return
}

func (t MarsTime) Format(layout string) (res string) {
	rotation, month, sol, vinqua, layer, fragment, rem := t.Params()
	week := (4 * (month - 1)) + (sol-1)/7 + 1
	weekSol := sol % 7
	if weekSol == 0 {
		weekSol = 7
	}

	longSol := t.LongWeekSolNames()[weekSol-1]
	shortSol := t.ShortWeekSolNames()[weekSol-1]
	longMonth := t.LongMonthNames()[month-1]
	shortMonth := t.ShortMonthNames()[month-1]

	vinquaLowercase := "a"
	vinquaUppercase := "A"
	if vinqua >= 12 {
		vinquaLowercase = "p"
		vinquaUppercase = "P"
	}
	vinqua11 := vinqua % 12
	vinqua12 := vinqua11
	if vinqua11 == 0 {
		vinqua12 = 12
	}

	replacements := map[string]string{
		"%R":    format.Iota(rotation),
		"%M":    format.Iota(month),
		"%0M":   format.Pad2(month),
		"%_M":   format.PadSpace(month),
		"%nM":   shortMonth,
		"%NM":   longMonth,
		"%S":    format.Iota(sol),
		"%oS":   format.Ordinal(sol),
		"%0S":   format.Pad2(sol),
		"%_S":   format.PadSpace(sol),
		"%D":    format.Iota(sol),
		"%oD":   format.Ordinal(sol),
		"%0D":   format.Pad2(sol),
		"%_D":   format.PadSpace(sol),
		"%w":    format.Iota(week),
		"%0W":   format.Pad2(week),
		"%_W":   format.PadSpace(week),
		"%W":    format.Iota(week),
		"%WS":   format.Iota(weekSol),
		"%NS":   longSol,
		"%nS":   shortSol,
		"%WD":   format.Iota(weekSol),
		"%ND":   longSol,
		"%nD":   shortSol,
		"%V":    format.Iota(vinqua),
		"%0V":   format.Pad2(vinqua),
		"%_V":   format.PadSpace(vinqua),
		"%V11":  format.Iota(vinqua11),
		"%0V11": format.Pad2(vinqua11),
		"%_V11": format.PadSpace(vinqua11),
		"%V12":  format.Iota(vinqua12),
		"%0V12": format.Pad2(vinqua12),
		"%_V12": format.PadSpace(vinqua12),
		"%Vl":   vinquaLowercase,
		"%Vu":   vinquaUppercase,
		"%L":    format.Iota(layer),
		"%0L":   format.Pad2(layer),
		"%_L":   format.PadSpace(layer),
		"%F":    format.Iota(fragment),
		"%0F":   format.Pad2(fragment),
		"%_F":   format.PadSpace(fragment),
		"%f":    format.RemoveZeroesFromDecimalPortionOfNumber(format.Pad9(rem)),
		"%f0":   format.Pad9(rem),
		"%%":    "%",
	}

	var builder strings.Builder
	for i := 0; i < len(layout); {
		if layout[i] == '%' {
			matched := false
			for length := 5; length > 1; length-- {
				if i+length <= len(layout) {
					token := layout[i : i+length]
					if token == "%'" {
						matched = true
						i += 2
						break
					}
					if val, ok := replacements[token]; ok {
						builder.WriteString(val)
						i += length
						matched = true
						break
					}
				}
			}
			if !matched {
				end := i + 1
				for end < len(layout) && end < i+4 {
					c := layout[end]
					if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_' {
						end++
					} else {
						break
					}
				}
				fragment := layout[i:end]
				return (`` +
					`error: fragment "` + fragment + `" not recognized: ` +
					`use "%%" for literal "%" and ` +
					`use "%'" to avoid conflict with possible future update ` +
					`for example use "%V%'E" when you want vinqua followed by "E" so that "%VE" can be used in future`)
			}
		} else {
			builder.WriteByte(layout[i])
			i++
		}
	}
	return builder.String()
}

// Deprrecated: use MarsTime{}.Parse(layout string, input string) instead
func ParseMarsTime(layout string, input string) (mt MarsTime, err error) {
	return mt.Parse(layout, input)
}

func (t MarsTime) Parse(layout string, input string) (mt MarsTime, err error) {
	var (
		rotation int
		month    int
		sol      int
		vinqua   int
		layer    int
		fragment int
		rem      int

		week    int
		weekSol int
	)

	vinquaRequiresAMPM := false
	vinquaFullfillsAMPM := false
	i, j := 0, 0
	for i < len(layout) {
		if layout[i] == '%' {
			if i+2 <= len(layout) && layout[i:i+2] == "%'" {
				i += 2
				continue
			}
			var token string
			found := false
			maxTokenLen := 5
			if i+maxTokenLen > len(layout) {
				maxTokenLen = len(layout) - i
			}
			for l := maxTokenLen; l >= 2; l-- {
				token = layout[i : i+l]
				if validToken(token) {
					i += l
					found = true
					break
				}
			}
			if !found {
				fragEnd := i + 1
				for fragEnd < len(layout) && fragEnd < i+maxTokenLen {
					fragEnd++
				}
				return MarsTime{}, fmt.Errorf(`error: fragment "%s" not recognized: use "%%%%" for literal "%%" and use "%%'" to avoid conflict with possible future update`, layout[i:fragEnd])
			}

			var value int
			var consumed int
			var parseErr error

			switch token {
			case
				"%R",
				"%M", "%0M", "%_M",
				"%S", "%0S", "%_S",
				"%D", "%0D", "%_D",
				"%V", "%0V", "%_V",
				"%L", "%0L", "%_L",
				"%V11", "%0V11", "%_V11",
				"%V12", "%0V12", "%_V12",
				"%F", "%0F", "%_F", "%f0",
				"%w", "%0W", "%_W", "%W",
				"%WS", "%wS", "%WD", "%wD":
				value, consumed, parseErr = format.ParseNumeric(input[j:])
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
				if token == "%V12" || token == "%_V12" || token == "%0V12" {
					vinquaRequiresAMPM = true
					for value >= 12 {
						value -= 12
					}
				}

			case "%f":
				value, consumed, parseErr = format.ParseDecimal(input[j:], 9)
			case "%Vl", "%Vu":
				value, consumed, parseErr = 0, 1, nil
				vinquaFullfillsAMPM = true
				if input[j] == 'a' || input[j] == 'A' {
					value = 0
				} else if input[j] == 'p' || input[j] == 'P' {
					value = 12
				} else {
					return MarsTime{}, fmt.Errorf("expected element of {'a', 'A', 'p', 'P'} for token %q at position %d in input", token, j)
				}
			case "%oS", "%oD":
				value, consumed, parseErr = format.ParseNumeric(input[j:])
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
				consumed += 2
			case "%%":
				if j >= len(input) || input[j] != '%' {
					return MarsTime{}, fmt.Errorf("expected literal '%%' at position %d in input", j)
				}
				consumed = 1
			case "%NM":
				value, consumed, parseErr = ParseMarsMonthName(input[j:], true)
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
			case "%nM":
				value, consumed, parseErr = ParseMarsMonthName(input[j:], false)
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
			case "%NS", "%ND":
				value, consumed, parseErr = ParseMarsWeekSolName(input[j:], true)
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
			case "%nS", "%nD":
				value, consumed, parseErr = ParseMarsWeekSolName(input[j:], false)
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
			}
			j += consumed

			switch token {
			case "%R":
				rotation = value
			case "%M", "%0M", "%_M", "%NM", "%nM":
				month = value
			case "%S", "%0S", "%_S", "%oS", "%D", "%0D", "%_D", "%oD":
				sol = value
			case "%V", "%0V", "%_V", "%V11", "%0V11", "%_V11", "%V12", "%0V12", "%_V12", "%Vl", "%Vu":
				vinqua += value
			case "%L", "%0L", "%_L":
				layer = value
			case "%F", "%0F", "%_F":
				fragment = value
			case "%f", "%f0":
				rem = value*int(Fragment)/int(time.Second) + 1
			case "%w", "%WS", "%nS", "%NS", "%WD", "%nD", "%ND":
				weekSol = value
			case "%W", "%0W", "%_W":
				week = value
			}
		} else {
			if j >= len(input) {
				return MarsTime{}, fmt.Errorf("input not fully consumed, remaining: %q", input[j:])
			}
			if layout[i] != input[j] {
				return MarsTime{}, fmt.Errorf("literal mismatch at layout[%d]=%q vs input[%d]=%q", i, layout[i], j, input[j])
			}
			i++
			j++
		}
	}
	if j != len(input) {
		return MarsTime{}, fmt.Errorf("input not fully consumed, remaining: %q", input[j:])
	}

	if month == 0 && sol == 0 && week > 0 && weekSol > 0 {
		month = (week-1)/4 + 1
		weekIndex := (week - 1) % 4
		sol = 7*weekIndex + weekSol
	}

	if rotation == 0 || month == 0 || sol == 0 {
		return MarsTime{}, fmt.Errorf("insufficient data to reconstruct MarsTime (month: %d, sol: %d)", month, sol)
	}

	if vinquaRequiresAMPM && !vinquaFullfillsAMPM {
		return MarsTime{}, fmt.Errorf("vinqua requires AM/PM specification but not provided; use token among {`%%V`, `%%0V`, `%%_V`} for vinqua parsing 00 thru 23")
	}

	totalSols := 0
	for r := 0; r < rotation; r++ {
		totalSols += t.SolsInRotation(r)
	}
	for m := 1; m < month; m++ {
		totalSols += t.SolsInMonth(m)
	}
	totalSols += (sol - 1)

	duration := time.Duration(vinqua)*Vinqua +
		time.Duration(layer)*Layer +
		time.Duration(fragment)*Fragment +
		time.Duration(rem)

	return MarsTime{
		TotalSols:            totalSols,
		DurationOfCurrentSol: duration,
	}, nil
}

func validToken(token string) bool {
	valid := map[string]bool{
		"%R":    true,
		"%M":    true,
		"%0M":   true,
		"%_M":   true,
		"%nM":   true,
		"%NM":   true,
		"%S":    true,
		"%oS":   true,
		"%0S":   true,
		"%_S":   true,
		"%D":    true,
		"%oD":   true,
		"%0D":   true,
		"%_D":   true,
		"%w":    true,
		"%0W":   true,
		"%_W":   true,
		"%W":    true,
		"%WS":   true,
		"%NS":   true,
		"%nS":   true,
		"%WD":   true,
		"%ND":   true,
		"%nD":   true,
		"%V":    true,
		"%0V":   true,
		"%_V":   true,
		"%V11":  true,
		"%0V11": true,
		"%_V11": true,
		"%V12":  true,
		"%0V12": true,
		"%_V12": true,
		"%Vl":   true,
		"%Vu":   true,
		"%L":    true,
		"%0L":   true,
		"%_L":   true,
		"%F":    true,
		"%0F":   true,
		"%_F":   true,
		"%f":    true,
		"%f0":   true,
		"%%":    true,
		"%'":    true,
	}
	return valid[token]
}

func ParseMarsMonthName(s string, long bool) (n int, nameLen int, err error) {
	return MarsTime{}.ParseMonthName(s, long)
}

func (t MarsTime) ParseMonthName(s string, long bool) (n int, nameLen int, err error) {
	var names []string
	if long {
		names = t.LongMonthNames()
	} else {
		names = t.ShortMonthNames()
	}
	for i, name := range names {
		if strings.HasPrefix(s, name) {
			return i + 1, len(name), nil
		}
	}
	err = fmt.Errorf("no matching month name found in %q", s)
	return
}

func ParseMarsWeekSolName(s string, long bool) (n int, nameLen int, err error) {
	return MarsTime{}.ParseWeekSolName(s, long)
}

func (t MarsTime) ParseWeekSolName(s string, long bool) (n int, nameLen int, err error) {
	var names []string
	if long {
		names = t.LongWeekSolNames()
	} else {
		names = t.ShortWeekSolNames()
	}
	for i, name := range names {
		if strings.HasPrefix(s, name) {
			return i + 1, len(name), nil
		}
	}
	err = fmt.Errorf("no matching weekSol name found in %q", s)
	return
}

func (mt MarsTime) Time() (result time.Time) {
	result = time.Date(1609, time.March, 11, 18, 40, 34, 0, time.UTC)

	// t.Sub is not able to go more than 2562047h47m16.854775807s
	quof := 100000
	for mt.TotalSols > quof {
		fmt.Println(Sol * time.Duration(quof))
		result = result.Add(Sol * time.Duration(quof))
		mt.TotalSols -= quof
	}
	result = result.Add(Sol * time.Duration(mt.TotalSols))
	result = result.Add(mt.DurationOfCurrentSol)

	return
}
