package main

import "fmt"

type gender struct{
	male,female bool
}

type salary struct {
	gender
	basic ,hra ,pf int
}

type emp struct {
	id int
	name ,address string
	worklocation string
	salary // embeded struct
	
}
 
// var e1 emp = emp {
// 	id: 1,
// 	name: "John Doe",
// 	address: "123 Elm St",
// 	Worklocation: "New York",
// }

// passing struct as value :
func withpointer (point *emp) emp{

	 return emp{
		id : point.id * 2,
		name: "Kumar",
		}
}

func main(){

	// p := emp{1, "abhi", "whitefield", "bangalore"}
	// p := &emp{1, "abhi", "whitefield", "bangalore"}
	p := emp{
		id : 1,
		name : "abhi",
		address : "Whitefield",
        worklocation : "Bangalore",
		salary : salary{
			gender : gender{
				male: true,
				female: false,
			},
			basic : 20000,
			hra : 5000,
			pf : 2000,
		},
	}

	// p:= emp{1,"abhi","Whitefield","bangalore",salary{gender{true,false},20000,5000,2000}}
	fmt.Println(p.id)
	fmt.Println(p.name)
	fmt.Println(p.address)
	fmt.Println(p.worklocation)
	fmt.Printf("%#v\n",p)
	
	withpointer(&p)
}
	

