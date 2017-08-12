package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KyleBanks/wilks"
)

const (
	genderMale   = "m"
	genderFemale = "f"
)

var (
	lbs    bool
	bw     float64
	total  float64
	gender string
)

func init() {
	flag.BoolVar(&lbs, "lbs", false, "[Optional] Specifies pounds as the unit of weight.")
	flag.Float64Var(&bw, "bw", 0, "The bodyweight of the lifter.")
	flag.Float64Var(&total, "total", 0, "The total weight lifted.")
	flag.StringVar(&gender, "gender", "", "The gender of the lifter, either 'm' for male or 'f' for female.")
	flag.Parse()
}

func main() {
	if !validate(bw, total, gender) {
		flag.Usage()
		os.Exit(1)
	}

	if lbs {
		bw = wilks.PoundsToKilos(bw)
		total = wilks.PoundsToKilos(total)
	}

	fmt.Printf("%.02f\n", wilks.Score(gender == genderFemale, bw, total))
}

func validate(bw, total float64, gender string) bool {
	if bw <= 0 || total <= 0 || (gender != genderMale && gender != genderFemale) {
		return false
	}

	return true
}
