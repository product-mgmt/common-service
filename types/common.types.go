package types

type Paginate struct {
	SearchColumn string
	SearchTerm   string
	SortColumn   string
	SortOrder    string
	Limit        int
	Offset       int
}

type StandardResponse struct {
	Message    string `json:"message" db:"message"`
	InsertedID int    `json:"insertedID" db:"insertedID"`
}
