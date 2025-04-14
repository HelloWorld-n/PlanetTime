package planets

import (
	"fmt"
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
	fmt.Println(sol, SolsInRotation(rotation))
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
	fmt.Println(rotation, month, sol, vinqua, layer, fragment, rem)
	return
}
