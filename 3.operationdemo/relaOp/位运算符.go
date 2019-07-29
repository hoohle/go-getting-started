package relationship

/*
位运算符的演示
													相同为0
				按位与	  按位或	按位异或	～按位取反		<<左移		>>右移
p	q		p & q		p | q			p ^ q
0	0			0				0					0
0	1			0				1					1
1	1			1				1					0
1	0			0				1					1
*/

func Bit2(x, y uint) uint {
	//位运算符为对整数在内存中的二进制位进行操作
	/*假定A =60 B=13
	  A = 			0011 	1100
	  B =			0000	1101
	  A&B =		0
	*/
	var c uint = 0
	c = x & y
	return c
}
