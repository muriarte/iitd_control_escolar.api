package entity

import "strconv"

//import "github.com/google/uuid"

//ID entity ID
type ID = int
//type ID = uuid.UUID

//NewID create a new entity ID
func NewID() ID {
	return 0
}
//func NewID() ID {
//	return ID(uuid.New())
//}

//StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := strconv.Atoi(s)
	return id, err
}

//func StringToID(s string) (ID, error) {
//	id, err := uuid.Parse(s)
//	return ID(id), err
//}

func IDToString(i ID) string {
	return strconv.Itoa(i)
}
