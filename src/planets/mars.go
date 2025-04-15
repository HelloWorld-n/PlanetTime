package planets

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

func MarsLongWeekSolNames() []string {
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

func MarsShortWeekSolNames() []string {
	return []string{
		"Sol", // for "Sol Solis"
		"Lun", // for "Sol Lunae"
		"Mar", // for "Sol Martis"
		"Mer", // for "Sol Mercurii"
		"Jov", // for "Sol Jovis"
		"Ven", // for "Sol Veneris"
		"Sat", // for "Sol Saturni"
	}
}

func MarsLongMonthNames() []string {
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

func MarsShortMonthNames() []string {
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

func RotationHasLeapVrishika28th(rotation int) (res bool) {
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

func SolsInRotation(rotation int) (sols int) {
	sols = 668
	if RotationHasLeapVrishika28th(rotation) {
		sols += 1
	}
	return
}

func SolsInMonth(month int) (sols int) {
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
	for sol >= SolsInRotation(rotation) {
		sol -= SolsInRotation(rotation)
		rotation += 1
	}
	for sol >= SolsInMonth(month) {
		sol -= SolsInMonth(month)
		month += 1
	}
	sol += 1

	vinqua = int(t.DurationOfCurrentSol / Vinqua)
	layer = int(t.DurationOfCurrentSol % Vinqua / Layer)
	fragment = int(t.DurationOfCurrentSol % Layer / Fragment)
	rem = int(t.DurationOfCurrentSol % Fragment)
	return
}

func (t MarsTime) Format(layout string) (res string) {
	rotation, month, sol, vinqua, layer, fragment, rem := t.Params()
	week := (4 * (month - 1)) + (sol-1)/7 + 1
	weekSol := sol % 7
	if weekSol == 0 {
		weekSol = 7
	}
	
	longSol := MarsLongWeekSolNames()[weekSol-1]
	shortSol := MarsShortWeekSolNames()[weekSol-1]
	longMonth := MarsLongMonthNames()[month-1]
	shortMonth := MarsShortMonthNames()[month-1]

	replacements := map[string]string{
		"%R":  itoa(rotation),
		"%M":  itoa(month),
		"%0M": pad2(month),
		"%_M": padSpace(month),
		"%nM": shortMonth,
		"%NM": longMonth,
		"%S":  itoa(sol),
		"%0S": pad2(sol),
		"%_S": padSpace(sol),
		"%D":  itoa(sol),
		"%0D": pad2(sol),
		"%_D": padSpace(sol),
		"%w":  itoa(week),
		"%0W": pad2(week),
		"%_W": padSpace(week),
		"%W":  itoa(week),
		"%WS": itoa(weekSol),
		"%ws": itoa(weekSol),
		"%NS": longSol,
		"%nS": shortSol,
		"%V":  itoa(vinqua),
		"%0V": pad2(vinqua),
		"%_V": padSpace(vinqua),
		"%L":  itoa(layer),
		"%0L": pad2(layer),
		"%_L": padSpace(layer),
		"%F":  itoa(fragment),
		"%0F": pad2(fragment),
		"%_F": padSpace(fragment),
		"%f":  itoa(rem),
		"%f0": pad9(rem),
		"%%":  "%",
	}

	var builder strings.Builder
	for i := 0; i < len(layout); {
		if layout[i] == '%' {
			matched := false
			for length := 4; length > 1; length-- {
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

func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}

func pad2(n int) string {
	return fmt.Sprintf("%02d", n)
}

func pad9(n int) string {
	return fmt.Sprintf("%09d", n)
}

func padSpace(n int) string {
	return fmt.Sprintf("%2d", n)
}

func ParseMarsTime(layout string, input string) (mt MarsTime, err error) {
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

	i, j := 0, 0
	for i < len(layout) {
		if layout[i] == '%' {
			if i+2 <= len(layout) && layout[i:i+2] == "%'" {
				i += 2
				continue
			}
			var token string
			found := false
			if i+3 <= len(layout) {
				token = layout[i : i+3]
				if validToken(token) {
					i += 3
					found = true
				}
			}
			if !found && i+2 <= len(layout) {
				token = layout[i : i+2]
				if validToken(token) {
					i += 2
					found = true
				}
			}
			if !found {
				fragEnd := i + 1
				for fragEnd < len(layout) && fragEnd < i+4 {
					fragEnd++
				}
				return MarsTime{}, fmt.Errorf(`error: fragment "%s" not recognized: use "%%%%" for literal "%%" and use "%%'" to avoid conflict with possible future update`, layout[i:fragEnd])
			}

			var value int
			var consumed int
			var parseErr error

			switch token {
			case
				"%R", "%M", "%0M", "%_M",
				"%S", "%0S", "%_S",
				"%D", "%0D", "%_D",
				"%V", "%0V", "%_V",
				"%L", "%0L", "%_L",
				"%F", "%0F", "%_F",
				"%f", "%f0",
				"%w", "%0W", "%_W", "%W", "%WS", "%ws":
				value, consumed, parseErr = ParseNumeric(input[j:])
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
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
			case "%NS":
				value, consumed, parseErr = ParseMarsWeekSolName(input[j:], true)
				if parseErr != nil {
					return MarsTime{}, fmt.Errorf("token %q: %v", token, parseErr)
				}
			case "%nS":
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
			case "%S", "%0S", "%_S", "%D", "%0D", "%_D":
				sol = value
			case "%V", "%0V", "%_V":
				vinqua = value
			case "%L", "%0L", "%_L":
				layer = value
			case "%F", "%0F", "%_F":
				fragment = value
			case "%f", "%f0":
				rem = value
			case "%w", "%WS", "%ws", "%nS", "%NS":
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

	totalSols := 0
	for r := 0; r < rotation; r++ {
		totalSols += SolsInRotation(r)
	}
	for m := 1; m < month; m++ {
		totalSols += SolsInMonth(m)
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
		"%R":  true,
		"%M":  true,
		"%0M": true,
		"%_M": true,
		"%nM": true,
		"%NM": true,
		"%S":  true,
		"%0S": true,
		"%_S": true,
		"%D":  true,
		"%0D": true,
		"%_D": true,
		"%w":  true,
		"%0W": true,
		"%_W": true,
		"%W":  true,
		"%WS": true,
		"%ws": true,
		"%NS": true,
		"%nS": true,
		"%V":  true,
		"%0V": true,
		"%_V": true,
		"%L":  true,
		"%0L": true,
		"%_L": true,
		"%F":  true,
		"%0F": true,
		"%_F": true,
		"%f":  true,
		"%f0": true,
		"%%":  true,
		"%'":  true,
	}
	return valid[token]
}

func ParseNumeric(s string) (n int, nRunes int, err error) {
	// allow leading spaces
	for nRunes < len(s) && s[nRunes] == ' ' {
		nRunes++
	}
	start := nRunes
	for nRunes < len(s) && s[nRunes] >= '0' && s[nRunes] <= '9' {
		nRunes++
	}
	if start == nRunes {
		err = fmt.Errorf("expected numeric value, got %q", s)
		return
	}
	numStr := s[start:nRunes]
	n, err = strconv.Atoi(numStr)
	return
}

func ParseMarsMonthName(s string, long bool) (n int, nameLen int, err error) {
	var names []string
	if long {
		names = MarsLongMonthNames()
	} else {
		names = MarsShortMonthNames()
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
	var names []string
	if long {
		names = MarsLongWeekSolNames()
	} else {
		names = MarsShortWeekSolNames()
	}
	for i, name := range names {
		if strings.HasPrefix(s, name) {
			return i + 1, len(name), nil
		}
	}
	err = fmt.Errorf("no matching weekSol name found in %q", s)
	return
}
