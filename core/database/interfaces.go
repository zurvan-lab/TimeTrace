package database

type IDataBase interface {
	SetsMap() Sets

	AddSet([]string) string
	AddSubSet([]string) string
	PushElement([]string) string
	DropSet([]string) string
	DropSubSet([]string) string
	CleanSet([]string) string
	CleanSets([]string) string
	CleanSubSet([]string) string
}
