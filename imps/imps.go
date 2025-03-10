package imps

func Imps(point int) (imps int) {
	return GiveMeImps(point)
}

func GiveMeImps(points int) (imps int) {
	if points < 0 {
		points = -points
	}
	if points < 20 {
		return 0
	}
	if points < 50 {
		return 1
	}
	if points < 90 {
		return 2
	}
	if points < 130 {
		return 3
	}
	if points < 170 {
		return 4
	}
	if points < 220 {
		return 5
	}
	if points < 270 {
		return 6
	}
	if points < 320 {
		return 7
	}
	if points < 370 {
		return 8
	}
	if points < 430 {
		return 9
	}
	if points < 500 {
		return 10
	}
	if points < 600 {
		return 11
	}
	if points < 750 {
		return 12
	}
	if points < 900 {
		return 13
	}
	if points < 1100 {
		return 14
	}
	if points < 1300 {
		return 15
	}
	if points < 1500 {
		return 16
	}
	if points < 1750 {
		return 17
	}
	if points < 2000 {
		return 18
	}
	if points < 2250 {
		return 19
	}
	if points < 2500 {
		return 20
	}
	if points < 3000 {
		return 21
	}
	if points < 3500 {
		return 22
	}
	if points < 4000 {
		return 23
	}
	return 24 //max as per The Laws of Duplicate Bridge 2017
}
