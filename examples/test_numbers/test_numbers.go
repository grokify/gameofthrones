package main

import (
	"fmt"

	"github.com/grokify/gameofthrones"

	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/sort/sortutil"
	"github.com/grokify/simplego/strconv/phonenumber"
)

func main() {
	a2g := phonenumber.NewAreaCodeToGeo()
	a2g.ReadData()
	fmtutil.PrintJSON(a2g)

	acs := a2g.AreaCodes()
	fmtutil.PrintJSON(acs)
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

	demoOrgs := gameofthrones.GetDemoOrganizations()
	fmtutil.PrintJSON(demoOrgs)
}
