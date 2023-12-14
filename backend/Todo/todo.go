package todo

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{
		ID:        "1",
		Item:      "Clean room",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Read book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Record video",
		Completed: true,
	},
}
