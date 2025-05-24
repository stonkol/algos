// graph.go
package main

type Edge struct {
	To     string
	Weight int
}

// AMES(15), ARPA, CMU, Harvard, Hawaii, Illinois, London, MIT(6), NBS, NSA, NYU, Pentagon, Stanford, Texas, UCLA, Utah, Xerox
var graph = map[string][]Edge{
	"AMES":     {{"ARPA", 15}},
	"ARPA":     {{"CMU", 10}, {"Harvard", 20}},
	"CMU":      {{"Harvard", 5}, {"NSA", 12}},
	"Harvard":  {{"NSA", 7}, {"NYU", 15}},
	"Hawaii":   {{"Illinois", 25}},
	"Illinois": {{"London", 30}},
	"London":   {{"MIT", 6}},
	"MIT":      {{"NBS", 8}},
	"NBS":      {{"NSA", 3}},
	"NSA":      {{"NYU", 4}},
	"NYU":      {{"Pentagon", 10}},
	"Pentagon": {{"Stanford", 18}},
	"Stanford": {{"Texas", 20}, {"UCLA", 15}},
	"Texas":    {{"UCLA", 10}},
	"UCLA":     {{"Utah", 12}, {"Xerox", 15}},
	"Utah":     {{"Xerox", 7}},
	"Xerox":    {},
}
