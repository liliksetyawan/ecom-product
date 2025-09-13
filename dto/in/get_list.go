package in

type Pagination struct {
	Limit  int
	Offset int
	Search string
}

type GetListDTO struct {
	Limit  int
	Offset int
	Search string
}
