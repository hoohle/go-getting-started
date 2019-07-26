package relationship

import "fmt"

/*
=直接赋值
+=相加后再赋值
-=相减后再赋值
*=相乘后再赋值
/=相处后再赋值
%=求余后再赋值
<<=左移后赋值
>>=右移后赋值
&=按位与后赋值
^=按位异或后赋值
|=按位或后赋值
*/

func Ass(x, z int) int {
	var y int = 1
	switch z {
	case 1:
		fmt.Println("=")
		y = x
		return y
	case 2:
		fmt.Println("+=")
		y += x
		return y
	case 3:
		fmt.Println("-=")
		y -= x
		return y
	default:
		fmt.Println("")
		return x
	}
}
