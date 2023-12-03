package database

import (
	"fmt"
	"time"

	"github.com/zurvan-lab/TimeTrace/config"
)

type Database struct {
	Sets   Sets
	Config *config.Config
}

func Init(path string) IDataBase {
	return &Database{
		Sets:   make(Sets, 1024),
		Config: config.LoadFromFile(path),
	}
}

func (db *Database) SetsMap() Sets {
	return db.Sets
}

// ! TQL Commands.
func (db *Database) AddSet(args []string) string {
	if len(args) != 1 {
		return "INVALID"
	}

	db.Sets[args[0]] = make(Set) // args[0] is set name. see: TQL docs.

	return "DONE"
}

func (db *Database) AddSubSet(args []string) string {
	if len(args) != 2 {
		return "INVALID"
	}

	s, ok := db.Sets[args[0]] // set name args[0]
	if !ok {
		return "SNF"
	}

	s[args[1]] = make(SubSet, 0) // subset name args[1]

	return "DONE"
}

func (db *Database) PushElement(args []string) string {
	if len(args) != 4 {
		return "INVALID"
	}

	setName := args[0]
	subSetName := args[1]
	elementValue := args[2]
	timeStr := args[3]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	t, err := time.Parse(time.UnixDate, timeStr)
	if err != nil {
		return "INVALID"
	}

	e := NewElement(elementValue, t)
	db.Sets[setName][subSetName] = append(db.Sets[setName][subSetName], e)

	return "DONE"
}

func (db *Database) DropSet(args []string) string {
	if len(args) != 1 {
		return "INVALID"
	}

	setName := args[0]
	_, ok := db.Sets[setName]

	if !ok {
		return "SNF"
	}

	delete(db.Sets, setName)

	return "DONE"
}

func (db *Database) DropSubSet(args []string) string {
	if len(args) != 2 {
		return "INVALID"
	}

	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	delete(db.Sets[setName], subSetName)

	return "DONE"
}

func (db *Database) CleanSets(_ []string) string {
	db.Sets = make(Sets)

	return "DONE"
}

func (db *Database) CleanSet(args []string) string {
	if len(args) != 1 {
		return "INVALID"
	}

	setName := args[0]

	_, ok := db.Sets[setName]
	if !ok {
		return "SNF"
	}

	db.Sets[setName] = make(Set)

	return "DONE"
}

func (db *Database) CleanSubSet(args []string) string {
	if len(args) != 2 {
		return "INVALID"
	}

	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	db.Sets[setName][subSetName] = make(SubSet, 0)

	return "DONE"
}

func (db *Database) CountSets(args []string) string {
	i := 0
	for range db.Sets {
		i++
	}

	return fmt.Sprint(i)
}

func (db *Database) CountSubSets(args []string) string {
	if len(args) != 1 {
		return "INVALID"
	}

	set, ok := db.Sets[args[0]]
	if !ok {
		return "SNF"
	}

	i := 0
	for range set {
		i++
	}

	return fmt.Sprint(i)
}

func (db *Database) CountElements(args []string) string {
	if len(args) != 2 {
		return "INVALID"
	}

	subSet, ok := db.Sets[args[0]][args[1]]
	if !ok {
		return "SSNF"
	}

	i := 0
	for range subSet {
		i++
	}

	return fmt.Sprint(i)
}
