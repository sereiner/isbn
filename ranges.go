package isbn

type PrefixeItem struct {
	Prefix string
	Agency string
	Rules  []*Rule
}

type Rule struct {
	Range  string
	Length string
}

type GroupItem struct {
	Prefix string
	Agency string
	Rules  []*Rule
}

var prefixes = []PrefixeItem{
	PrefixeItem{
		Prefix: "978",
		Agency: "International ISBN Agency",
		Rules: []*Rule{
			&Rule{
				Range:  "0000000-5999999",
				Length: "1",
			},
			&Rule{
				Range:  "6000000-6499999",
				Length: "3",
			},
			&Rule{
				Range:  "6500000-6599999",
				Length: "2",
			},
			&Rule{
				Range:  "6600000-6999999",
				Length: "0",
			},
			&Rule{
				Range:  "7000000-7999999",
				Length: "1",
			},
			&Rule{
				Range:  "8000000-9499999",
				Length: "2",
			},
			&Rule{
				Range:  "9500000-9899999",
				Length: "3",
			},
			&Rule{
				Range:  "9900000-9989999",
				Length: "4",
			},
			&Rule{
				Range:  "9990000-9999999",
				Length: "5",
			},
		},
	},
	PrefixeItem{
		Prefix: "979",
		Agency: "International ISBN Agency",
		Rules: []*Rule{
			&Rule{
				Range:  "0000000-0999999",
				Length: "0",
			},
			&Rule{
				Range:  "1000000-1299999",
				Length: "2",
			},
			&Rule{
				Range:  "1300000-9999999",
				Length: "0",
			},
		},
	},
}
