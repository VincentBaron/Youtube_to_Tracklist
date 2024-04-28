// Definition of the structures and SQL interaction functions
package models

// Models return one of every model the database must create..
func Models() []interface{} {
	return []interface{}{
		&Playlist{},
		&Track{},
	}
}
