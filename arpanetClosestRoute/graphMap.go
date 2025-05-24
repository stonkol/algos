// graph.go
//
// AMES(15, NASA), ARPA(Advanced Research Projects Agency), CMU, Harvard, Hawaii, Illinois, UCL(London, first intl' ARPANET node), MIT(6), NBS(National Bureau of Standards), NSA, NYU, Pentagon, Scott, SDCA(Seismic Data Analysis Center), Stanford, Texas, UCLA(first node), Utah, Xerox

package main

type Edge struct {
	To       string
	Distance int
}

var graph = map[string][]Edge{
	"Ames":     {{"Hawaii", 3}, {"Stanford", 7}, {"Utah", 12}, {"Ucla", 15}, {"Xerox", 8}},
	"Arpa":     {{"Pentagon", 5}, {"Sdac", 2}, {"Texas", 14}},
	"Cmu":      {{"Lincoln", 6}, {"Scott", 7}, {"Sdac", 15}},
	"Harvard":  {{"Lincoln", 9}, {"Mit", 14}, {"Nyu", 1}, {"Scott", 12}},
	"Hawaii":   {},
	"Illinois": {{"Mit", 10}, {"Utah", 3}},
	"Lincoln":  {{"Mit", 7}},
	"Mit":      {},
	"Nbs":      {{"Nsa", 3}, {"Nyu", 4}, {"Pentagon", 4}},
	"Nsa":      {{"Nyu", 4}, {"Sdac", 3}},
	"Nyu":      {{"SCAC", 10}},
	"Pentagon": {{"Texas", 14}},
	"Scott":    {{"Stanford", 16}, {"Ucla", 11}, {"Xerox", 13}},
	"Sdac":     {{"Texas", 8}, {"Ucl", 7}},
	"Stanford": {{"Texas", 22}, {"Ucla", 11}, {"Xerox", 9}},
	"Texas":    {{"Ucla", 20}},
	"Ucl":      {},
	"Ucla":     {{"Xerox", 15}},
	"Utah":     {{"Xerox", 7}},
	"Xerox":    {},
}
