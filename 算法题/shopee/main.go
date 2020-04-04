package main

import "fmt"

//
//
//e string
//Password string
//LastWrongTime int64
//WrongCount int
//State int
//}
//
//// Product
//type Product struct{
//	ID int
//	Name string
//	Remain int64
//	CategoryId int
//	Price int
//}
//// login/logout/buy
//
//
//// Category
//type Category struct {
//	ID int
//	Name string
//}
//
//// Order
//type Order struct {
//	ID int
//	UserId int
//	ProductId int
//	Count int
//	time time
//}
//
//
//// field_name / field_type
//
//// input: UserId
//// output: CategoryId
///*
//User.GetUserByUserId(UserId) -> User user -> user.Id, user.Name, error
//
//Product.GetProductListByCategoryId(CategoryId) -> Product[] {product, product, ....} -> product.Name, product.Price
//
//Product.GetProductListByCategoryIdList(int[]) -> Product[]
//*/
//var productPrice = map[int][]int{}
//func getMaxOutputCategory(UserId int) int{
//	userID,_,err := User.GetUserByUserId(UserId)
//	if err !=nil{
//		....
//	}
//	productOutput := map[int]int{}
//	orders := Order.GetOrdersByUserId(userID)
//	maxOutput := 0
//	,errresult := 0
//	for _,order := range
//	if err !=nil{
//		....
//	}orders{
//		if _,ok := productOutput[order.ProductId];!ok{
//		productOutput[order.ProductId] = 0
//	}
//		price := 0
//		categoryId := 0
//		if _,ok := productPrice[order.ProductId];!ok{
//		price,categoryId,err := Product.GetProductDetail(order.ProductId)
//		if err !=nil{
//		....
//	}
//		productPrice[order.ProductId] = []int{price,categoryId}productPrice[order.ProductId] = []int{price,categoryId}
//	}else{
//		price = productPrice[order.ProductId][0]
//		categoryId = productPrice[order.ProductId][1]
//	}
//
//		productOutput[order.ProductId] += order.Count*price
//		if maxOutput < productOutput[order.ProductId]{
//		maxOutput= productOutput[order.ProductId]
//		result = categoryId
//	}
//	}
//
//	return result
//}

var data = Node{
	Name:  "data",
	Value: "",
	NodeList: []Node{
		{"a1", "", []Node{
			{"a2", "Product AA", []Node{
				{"a3", "Product A", nil},
			}},
		}},
		{"b1", "Special Sell", nil},
		{"c1", "", []Node{
			{"c2", "", []Node{
				{"c2a", "flower1", nil},
				{"c2b", "flower2", nil},
			}},
		}},
		{"d1", "", []Node{
			{"d2a", "", []Node{
				{"d2a1", "a", nil},
				{"d2a2", "Product B", nil},
			}},
			{"d2b", "", []Node{
				{"d3b", "", []Node{
					{"d4b", "Product C", nil},
				}},
				{"d3c", "product D", nil},
			}},
		}},
	},
}

type Node struct {
	Name     string
	Value    string
	NodeList []Node
}

//data{Name: "data", Value: nil, NodeList: [a1, b1, c1, d1]}
//a1 {Name: "a1", Value: nil, NodeList: [a2]}
//a3 {Name: "a3", Value: nil, NodeList: [ProducA]}
//ProductA {Name: nil, Value: "Product A", NodeList: nil, []}

// intput: data, name
// output string[]

// data, "Product A"
// output: ["a1", "a2", "a3"]

func main() {
	fmt.Println(getTranverserPath(data, "Product B"))
}

func getTranverserPath(data Node, name string) []string {
	var nodeList = []Node{data}
	var handled = map[string]bool{}
	for len(nodeList) > 0 {
		fmt.Println(len(nodeList))
		last := nodeList[len(nodeList)-1]
		fmt.Println(last.Value)
		if last.Value == name {
			break
		}
		if handled[last.Name] {
			nodeList = nodeList[:len(nodeList)-1]
			continue
		}
		if len(last.NodeList) > 0 {
			//handled[last.Name] = true
			head := last.NodeList[0]
			nodeList = append(nodeList, head)
			last.NodeList = last.NodeList[1:]
			continue
		}

		if len(last.NodeList) == 0 {
			handled[last.Name] = true
		}
	}

	res := []string{}
	for _, node := range nodeList {
		res = append(res, node.Name)
	}
	return res

}
