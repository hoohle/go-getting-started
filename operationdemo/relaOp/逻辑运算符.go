package relationship

func Logic(x, y bool) string {
	if x && y {
		return "x与b为真 x或y为假	一假为假"
	} else if x || y {
		return "x或y为真 x与y为假	一真全真"
	} else {
		return "false"
	}
}
