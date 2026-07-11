package ormcp

import "github.com/tinywasm/model"

var QueryArgsModel = model.Definition{
	Name: "query_args",
	Fields: model.Fields{
		{
			Name: "SQL",
			Type: model.Text(),
			NotNull: true,
			Permitted: model.Permitted{
				Letters:   true,
				Numbers:   true,
				Spaces:    true,
				BreakLine: true,
				Tab:       true,
				Extra:     []rune{'*', '=', '>', '<', '"', '\'', '(', ')', ',', '.', ';', '_', '-', '+', '/', '%', '?', '!', '@', '#', '$', '^', '&', '|', '~', '`', '[', ']', '{', '}', ':'},
			},
		},
	},
}

var ExecArgsModel = model.Definition{
	Name: "exec_args",
	Fields: model.Fields{
		{
			Name: "SQL",
			Type: model.Text(),
			NotNull: true,
			Permitted: model.Permitted{
				Letters:   true,
				Numbers:   true,
				Spaces:    true,
				BreakLine: true,
				Tab:       true,
				Extra:     []rune{'*', '=', '>', '<', '"', '\'', '(', ')', ',', '.', ';', '_', '-', '+', '/', '%', '?', '!', '@', '#', '$', '^', '&', '|', '~', '`', '[', ']', '{', '}', ':'},
			},
		},
	},
}
