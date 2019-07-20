package main

var trips = [][]string{
	[]string{"重庆", "新疆"},
	[]string{"广州", "南京"},
	[]string{"上海", "广州"},
	[]string{"西藏", "重庆"},
	[]string{"北京", "上海"},
	[]string{"南京", "西藏"},
}

type Trip struct {
	Index int
	From string
	To string
}


func main() {
	//本题时间为20分钟
	//trips是小明出差的行程， 请找出小明这次行程，并注释时间复杂度与空间复杂度
	//结果 [北京"，"上海", "广州", "南京", "西藏", "重庆"，"新疆"]
	newTrips := []Trip{}
	for k,v := range trips{
		oneTrip := Trip{
			Index:k,
			From:v[0],
			To:v[1],
		}
		newTrips = append(newTrips,oneTrip)
	}



	//head := make(map[string]int,len(trips))
	//tail := make(map[string]int,len(trips))
	//for k,v := range trips{
	//	head[v[0]]=k
	//	tail[v[1]]=k
	//}
	//fmt.Println("head:",head)
	//fmt.Println("tail:",tail)
	//res := []string{}
	//var headindex int
	////resindex := []int{}
	//for k,_ := range head{
	//	if _,ok := tail[k]; !ok{
	//		headindex = head[k]
	//		break
	//	}
	//}
	//
	//for i:=0;i<len(trips);i++{
	//	if i==0{
	//		res = append(res,trips[headindex][0])
	//		headindex = head[trips[headindex][1]]
	//		fmt.Println(headindex)
	//		continue
	//	}
	//	res = append(res,trips[headindex][0])
	//	headindex = head[trips[headindex][1]]
	//	fmt.Println(headindex)
	//	if i == len(trips)-1{
	//		res = append(res,trips[headindex][1])
	//	}
	//}
	//
	//fmt.Println(res)
}