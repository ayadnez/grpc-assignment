package types

type User struct {
	ID      int32
	FName   string
	City    string
	Phone   string
	Height  float32
	Married bool
}

var users = map[int32]User{
	1: {ID: 1, FName: "zahid", City: "SGR", Phone: "1234567890", Height: 5.9, Married: true},
	2: {ID: 2, FName: "John", City: "NY", Phone: "2345678901", Height: 5.9, Married: false},
	3: {ID: 3, FName: "ahmed", City: "Delhi", Phone: "234567345", Height: 5.7, Married: false},
	4: {ID: 4, FName: "George", City: "Srinagar", Phone: "234598777", Height: 5.9, Married: true},
	5: {ID: 5, FName: "Gautham", City: "Chennai", Phone: "2345678901", Height: 5.8, Married: false},
}
