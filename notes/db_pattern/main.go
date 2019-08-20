

type Query interface {
	GUID() string
	Run() (*sql.Rows, error)
}

type Store struct {
	Generator UserQueryGenerator	
}

func (u UserStore) Find(userID string) (*User, error) {
	// do we even need our own query object? how about sql.Query
	query := u.Generator.FindUser(userID)
	result, exists := u.cache.Find(query.GUID())
	if exists {
		return result.(*User), nil
	}

	rows, err := query.Run() 
	if err != nil {
		return err
	}

	for _, row := range rows.Next() {
		row.Scan(&user)			
	}

	return rows.Err()
}



