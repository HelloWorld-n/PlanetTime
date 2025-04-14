package planets

import (
	"fmt"
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

var longWeekSolNames = []string{
	"Solis",
	"Lunae",
	"Martis",
	"Mercurii",
	"Jovis",
	"Veneris",
	"Saturni",
}

var shortWeekSolNames = []string{
	"Sol", // for "Sol Solis"
	"Lun", // for "Sol Lunae"
	"Mar", // for "Sol Martis"
	"Mer", // for "Sol Mercurii"
	"Jov", // for "Sol Jovis"
	"Ven", // for "Sol Veneris"
	"Sat", // for "Sol Saturni"
}

var longMonthNames = []string{
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

var shortMonthNames = []string{
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
	longSol := longWeekSolNames[weekSol-1]
	shortSol := shortWeekSolNames[weekSol-1]
	longMonth := longMonthNames[month-1]
	shortMonth := shortMonthNames[month-1]

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
				return (`` +
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
