package main

import (
	"fmt"

	"github.com/grokify/gophonenumbers"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/sort/sortutil"

	"github.com/grokify/gameofthrones"
)

func main() {
	a2g := gophonenumbers.NewAreaCodeToGeo()
	err := a2g.ReadData()
	logutil.FatalErr(err)
	fmtutil.MustPrintJSON(a2g)

	acs := a2g.AreaCodes()
	fmtutil.MustPrintJSON(acs)
	sortutil.Uint16s(acs)
	fmt.Println("Ints:   ", acs)

	orgs := gameofthrones.Organizations
	for i, org := range orgs {
		j := i * 10
		if j >= len(acs) {
			panic("A")
		}
		ac := acs[j]
		fmt.Printf("%v %v %v\n", i, org, ac)
	}

	demoOrgs, err := gameofthrones.GetDemoOrganizations()
	logutil.FatalErr(err)
	fmtutil.MustPrintJSON(demoOrgs)
}
