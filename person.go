package person

type user struct {

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Address	  string `json:"address"`
	JobInfo   map[string]string  `json:"jobinfo"`
	username   int    `json:"username"`
	password  string  `json:"password"`
}
