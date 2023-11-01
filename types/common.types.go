package types

type Paginate struct {
	SearchColumn string
	SearchTerm   string
	SortColumn   string
	SortOrder    string
	Limit        int
	Offset       int
}
