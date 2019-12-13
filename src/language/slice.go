package main

import "fmt"

func updateSlice(sliceData []int) {
	//
}

func main() {

	sliceData := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("s1:[2:6] :", sliceData[2:6])

	fmt.Println("s2:[:6] :", sliceData[:6])

	fmt.Println("s3:[2:] :", sliceData[2:])

	fmt.Println("s4:[:] :", sliceData[:])

	//slice ext
	sExt := sliceData[2:6]
	fmt.Println("sExt :", sExt) //[2 3 4 5]
	for i := 0; i < len(sExt); i++ {
		fmt.Println(sExt[i])
	}
	//fmt.Println(sExt[6])//这里会报错
	sExt2 := sExt[3:5]            //重点关注，显然长度中不存在4，5下标，但是因为slice是arr的view，所以依然能指向底层数组的下标
	fmt.Println("sExt2 :", sExt2) // [5 6]

	//opt【注意这里的机制】
	sExt = append(sExt, 100)
	fmt.Println("sExt :", sExt)           //[2 3 4 5 100]
	fmt.Println("sliceData :", sliceData) //当append时没有超过cap时，会覆盖sExt中len下标最后指向的arr的元素的值[0 1 2 3 4 5 100 7]

	sExt = append(sExt, 101)
	fmt.Println("sExt :", sExt)           //[2 3 4 5 100 101]
	fmt.Println("sliceData :", sliceData) //[0 1 2 3 4 5 100 101]

	sExt = append(sExt, 102)              //此时已经超出slice原定的cap，此时sExt只会view一个新的arr，查看旧的arr发现不会进行扩展
	fmt.Println("sExt :", sExt)           //[2 3 4 5 100 101 102]
	fmt.Println("sliceData :", sliceData) //[0 1 2 3 4 5 100 101]
}
