package ahaslides

import su "github.com/grokify/gogoogle/slidesutil/v1"

var DefaultFormatting = RoadmapFormatting{
	Line: LineFormatting{
		ColorHex:  "#f29d39", // "#6f6f6f", // "#ff8800"
		DashStyle: "DASH",
	},
	//ColumnHeadingForegroundColorHex: "#1072bd",
	Column: ColumnFormatting{
		Heading: su.CreateShapeTextBoxRequestInfo{
			FontBold:           true,
			FontItalic:         false,
			FontSize:           12.0,
			FontSizeUnit:       "PT",
			ForegroundColorHex: "#6f6f6f", // "#1072bd",
		},
	},
	Row: RowFormatting{
		Heading: su.CreateShapeTextBoxRequestInfo{
			FontBold:           true,
			FontItalic:         false,
			FontSize:           15.0,
			FontSizeUnit:       "PT",
			ForegroundColorHex: "#6f6f6f",
		},
		//HeadingForegroundColorHex: "#6f6f6f", // "#ff8800",
		RowOddBackgroundColorHex: "#ffffff", // "#ededed" // "#f6f6f6",
	},
	Disclaimer: TextboxFormatting{
		DefaultBackgroundColorHex: "#ff8800",
		DefaultForegroundColorHex: "#ffffff",
	},
	Textbox: TextboxFormatting{
		DefaultBackgroundColorHex: "#2e73a9", // "#4688f1",
		DefaultForegroundColorHex: "#ffffff",
		DeadBackgroundColorHex:    "#aaaaaa",
		DeadForegroundColorHex:    "#ffffff",
		DoneBackgroundColorHex:    "#00ac00",
		DoneForegroundColorHex:    "#ffffff",
		ProblemBackgroundColorHex: "#ea8c34", // "#ff8800",
		ProblemForegroundColorHex: "#ffffff",
	},
}

type RoadmapFormatting struct {
	//ColumnHeadingForegroundColorHex string            `json:"columnHeadingFontColorHex"`
	Column     ColumnFormatting  `json:"column"`
	Row        RowFormatting     `json:"row"`
	Line       LineFormatting    `json:"line"`
	Textbox    TextboxFormatting `json:"textbox"`
	Disclaimer TextboxFormatting `json:"disclaimer"`
}

type ColumnFormatting struct {
	Heading su.CreateShapeTextBoxRequestInfo `json:"heading"`
}

type RowFormatting struct {
	Heading                   su.CreateShapeTextBoxRequestInfo `json:"heading"`
	RowOddBackgroundColorHex  string                           `json:"rowOddBgColorHex"`
	RowEvenBackgroundColorHex string                           `json:"rowEvenBgColorHex"`
	//HeadingForegroundColorHex string `json:"headingFontColorHex"`
}

type LineFormatting struct {
	ColorHex  string `json:"defaultColor"`
	DashStyle string `json:"dashStyle"` // e.g. `DASH`
}

type TextboxFormatting struct {
	DefaultForegroundColorHex string `json:"defaultForegroundColorHex"`
	DefaultBackgroundColorHex string `json:"defaultBackgroundColorHex"`
	DeadForegroundColorHex    string `json:"deadForegroundColorHex"`
	DeadBackgroundColorHex    string `json:"deadBackgroundColorHex"`
	DoneForegroundColorHex    string `json:"doneForegroundColorHex"`
	DoneBackgroundColorHex    string `json:"doneBackgroundColorHex"`
	ProblemForegroundColorHex string `json:"problemForegroundColorHex"`
	ProblemBackgroundColorHex string `json:"problemBackgroundColorHex"`
}
