// Package mobiledoc holds structures and helper functions for working with
// mobiledoc data.
package mobiledoc

const version = "0.3.1"

type (
	Mobiledoc struct {
		Version  string          `json:"version"`
		Atoms    []interface{}   `json:"atoms"`
		Cards    [][]interface{} `json:"cards"`
		Markups  []interface{}   `json:"markups"`
		Sections [][]interface{} `json:"sections"`
	}

	Card struct {
		CardName *string `json:"cardName,omitempty"`
		HTML     *string `json:"html,omitempty"`
		Markdown *string `json:"markdown,omitempty"`
	}
)

// FromMarkdown wraps the given Markdown in a Mobiledoc struct
func FromMarkdown(md string) *Mobiledoc {
	return &Mobiledoc{
		Version: version,
		Atoms:   []interface{}{},
		Cards: [][]interface{}{
			[]interface{}{
				"markdown",
				Card{
					CardName: str("markdown"),
					Markdown: str(md),
				},
			},
		},
		Markups: []interface{}{},
		Sections: [][]interface{}{
			[]interface{}{10, 0},
		},
	}
}

func str(s string) *string {
	return &s
}
