package gameofthrones

import (
	"bytes"
	_ "embed"

	"github.com/grokify/gocharts/v2/data/table"
)

//go:embed characters.csv
var charactersDataBytes []byte

func Characters() []Character {
	var chars []Character
	tbl, err := table.ParseReadSeeker(nil, bytes.NewReader(charactersDataBytes))
	if err != nil {
		panic(err)
	}
	for _, row := range tbl.Rows {
		if len(row) != 4 {
			continue
		}
		chars = append(chars, NewCharacterSimple(NewCharacterSimpleOpts{
			ActorName:       row[0],
			GivenName:       row[1],
			FamilyName:      row[2],
			NickName:        row[3],
			AddOrganization: true,
		}))
	}
	return chars
}
