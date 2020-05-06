package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//广度优先算法实现迷宫的最短路径
//深度优先和广度优先的实现对出队的点以及新入队点的顺序不同

//任何一个关联图上的点的实现逻辑，都是将该点随机放在一个点，而不是固定在起始
//从随机的一个点来找通用的规则和跳出条件【这就是通用的逻辑：构造逻辑单元】
//然后再想特殊的情况来进行匹配测试

//以及检查结果（最短路径）从结果开始寻找最短路径【广度优先】

//读取迷宫的配置文件，并形成二维数组
func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		//
	}

	//配置文件格式化
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col) //先读取第一行来获取基础的配置信息

	maze := make([][]int, row) //二维数组的基本构造
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j]) //形成的构造关系
		}
	}

	return maze
}

//表达对应位置即可【不使用坐标】
type point struct {
	i int
	j int
}

//定义一个方向的结构【保持顺序：上左下右】：和坐标有区别，对数据需要进行更贴合的处理
var dirs = [4]point{
	{-1, 0}, //上
	{0, -1}, //左
	{1, 0},  //下
	{0, 1},  //右
}

//根据单位方向演进下一节点
func (cur point) add(dir point) point {
	return point{cur.i + dir.i, cur.j + dir.j}
}

//判定某点在地图上的位置【是否越界和获取该值】
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

//更加配置化
func walk(maze [][]int, start, end point) [][]int {
	//构造一个结构记录每步形成的结构，值就是底层的步数，后面通过对该结构的反向查找来获得最终的路径

	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//初始化结构完成
	//广度优先就是从该点开始逐个遍历可能性，进行入队操作【队列实现】
	Q := []point{start} //将开始点进行入队操作

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:] //出队

		//广度优先的退出：显然因为当前点如果已经到达目的地，显然就是最快到达的，因为后面还没有到达的必然不是最优
		//因为记录的是step的长度，判定的是next带来的step的数值，所以steps结构标记了该点的最先到达的step就是局部的最短
		//如果改点最后是经过的点，那么该step值最小显然也是应该被利用的标记值
		if cur == end {
			break
		}

		//进行广度优先
		for _, dir := range dirs {
			//需要结合方向来获得下一个点：但是显然是需要进行的计算不是简单的+
			//next := cur + dir

			//对下一点的演进
			next := cur.add(dir)

			//接着判断下一点的特性：
			val, ok := next.at(maze)
			//说明该点越界或者是墙
			if !ok || val == 1 {
				continue
			}

			//再根据steps这个结构来过滤已经走过的点【说明已经入队过，因为入队的点已经遍历过所有的周围的点，所以不进入就是不回退】
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			//说明此时的next就是可走的下一点【也就是广度扩展了一个新的点】
			//除了新增队列点，还需要更新steps,以及下一步的点数
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}

	}

	return steps
}

func main() {
	filePath := GetCurrentPath() + "/test.txt"
	maze := readMaze(filePath) //将配置信息格式化

	/*for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}

		fmt.Println()
	}*/

	//开始执行
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)

	//对步骤内各步进行查看
	for _, row := range steps {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}

		fmt.Println()
	}
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
