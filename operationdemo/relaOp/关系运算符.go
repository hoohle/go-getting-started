package relationship

/*
==判断两个值是否相等
！=检查两个值是否不想等
>
<
>=
<=
以上全返回布尔函数boolen
*/
//---------------------demo-----------------
func Isequal(x, y int) string {

	if x == y {
		return "相等"
	} else {
		if x < y {
			return "第二个值大"
		} else if x > y {
			return "地一个值大"
		} else {
			return "不等"
		}
	}

}
