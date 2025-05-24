// graph.go
//
// AMES(15, NASA), ARPA(Advanced Research Projects Agency), CMU, Harvard, Hawaii, Illinois, UCL(London, first intl' ARPANET node), MIT(6), NBS(National Bureau of Standards), NSA, NYU, Pentagon, Scott, SDCA(Seismic Data Analysis Center), Stanford, Texas, UCLA(first node), Utah, Xerox

package main

type Edge struct {
	To     string
	Weight int
}

var graph = map[string][]Edge{
	"AMES":     {{"Hawaii", 3}, {"Stanford", 7}, {"UCLA", 15}, {"Utah", 12}, {"Xerox", 8}},
	"ARPA":     {{"SDAC", 2}, {"Pentagon", 5}, {"Texas", 14}},
	"CMU":      {{"SDCA", 15}, {"Scott", 7}, {"Lincoln", 6}},
	"Harvard":  {{"Scott", 12}, {"MIT", 14}, {"NYU", 1}, {"Lincoln", 9}},
	"Hawaii":   {},
	"Illinois": {{"MIT", 10}, {"Utah", 3}},
	"Lincoln":  {{"MIT", 7}},
	"MIT":      {},
	"NBS":      {{"NYU", 4}, {"NSA", 3}, {"Pentagon", 4}},
	"NSA":      {{"NYU", 4}, {"SDAC", 3}},
	"NYU":      {{"SCAC", 10}},
	"Pentagon": {{"Texas", 14}},
	"Scott":    {{"Stanford", 16}, {"UCLA", 11}, {"Xerox", 13}},
	"SDAC":     {{"UCL", 7}, {"Texas", 8}},
	"Stanford": {{"Texas", 22}, {"UCLA", 11}, {"Xerox", 9}},
	"Texas":    {{"UCLA", 20}},
	"UCL":      {},
	"UCLA":     {{"Xerox", 15}},
	"Utah":     {{"Xerox", 7}},
	"Xerox":    {},
}
