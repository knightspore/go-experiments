package main

import "fmt"

type Loc struct {
	Row  int     `json:"row"`
	Col  int     `json:"col"`
	Path *string `json:"path"`
}

func NewLoc(Row int, Col int, Path *string) *Loc {
	path := fmt.Sprintf("%s:%d:%d", *Path, Row+1, Col+1)
	return &Loc{
		Row:  Row + 1,
		Col:  Col + 1,
		Path: &path,
	}
}
