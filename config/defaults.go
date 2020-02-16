package config

import "runtime"

var DefaultConfig = Config{
	DebugMode: false,
	ColourScheme: ColourScheme{
		Cursor:       strToColourNoErr("#969a97"),
		Foreground:   strToColourNoErr("#feffff"),
		Background:   strToColourNoErr("#2d2b33"),
		Black:        strToColourNoErr("#383640"),
		Red:          strToColourNoErr("#a09ebb"),
		Green:        strToColourNoErr("#92b6b1"),
		Yellow:       strToColourNoErr("#edaf97"),
		Blue:         strToColourNoErr("#8789c0"),
		Magenta:      strToColourNoErr("#800080"),
		Cyan:         strToColourNoErr("#6b9080"),
		LightGrey:    strToColourNoErr("#585660"),
		DarkGrey:     strToColourNoErr("#484650"),
		LightRed:     strToColourNoErr("#c0bedb"),
		LightGreen:   strToColourNoErr("#b2d6d1"),
		LightYellow:  strToColourNoErr("#fdbfa7"),
		LightBlue:    strToColourNoErr("#a7a9e0"),
		LightMagenta: strToColourNoErr("#a000a0"),
		LightCyan:    strToColourNoErr("#8bb0a0"),
		White:        strToColourNoErr("#969a97"),
		Selection:    strToColourNoErr("#a6aaa7"),
	},
	KeyMapping:            KeyMappingConfig(map[string]string{}),
	SearchURL:             "https://www.google.com/search?q=$QUERY",
	MaxLines:              1000,
	CopyAndPasteWithMouse: true,
}

func init() {
	DefaultConfig.KeyMapping[string(ActionCopy)] = addMod("c")
	DefaultConfig.KeyMapping[string(ActionPaste)] = addMod("v")
	DefaultConfig.KeyMapping[string(ActionSearch)] = addMod("g")
	DefaultConfig.KeyMapping[string(ActionToggleDebug)] = addMod("d")
	DefaultConfig.KeyMapping[string(ActionToggleSlomo)] = addMod(";")
	DefaultConfig.KeyMapping[string(ActionReportBug)] = addMod("r")
}

func addMod(keys string) string {
	standardMod := "ctrl + shift + "

	if runtime.GOOS == "darwin" {
		standardMod = "super + "
	}

	return standardMod + keys
}
