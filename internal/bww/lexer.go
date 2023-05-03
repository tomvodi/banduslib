package bww

import "github.com/alecthomas/participle/v2/lexer"

var BwwLexer = lexer.MustStateful(lexer.Rules{
	"Root": {
		{
			Name:    "BagpipeReader",
			Pattern: `Bagpipe Reader`,
			Action:  lexer.Push("BagpipeReader"),
		},
		{
			Name:    "TEMPO_DEF",
			Pattern: `TuneTempo`,
			Action:  lexer.Push("TuneTempo"),
		},
		{
			Name:    "PARAM_START",
			Pattern: `\(`,
			Action:  lexer.Push("ParamList"),
		},
		{
			Name:    "PARAM_DEF",
			Pattern: `MIDINoteMappings|FrequencyMappings|InstrumentMappings|GracenoteDurations|FontSizes|TuneFormat`,
		},
		{
			Name:    "PARAM_SEP",
			Pattern: `,`,
		},
		{
			Name:    "STAFF_START",
			Pattern: `&`,
			Action:  lexer.Push("Staff"),
		},
		{
			Name:    "STRING",
			Pattern: `"[^"]*"`,
		},
		{
			Name:    "WHITESPACE",
			Pattern: `\s+`,
		},
	},
	"BagpipeReader": {
		{
			Name:    "VERSION_SEP",
			Pattern: `:`,
		},
		{
			Name:    "VersionNumber",
			Pattern: `\d+\.\d+`,
			Action:  lexer.Pop(),
		},
	},
	"TuneTempo": {
		{
			Name:    "PARAM_SEP",
			Pattern: `,`,
		},
		{
			Name:    "TEMPO_VALUE",
			Pattern: `\d+`,
			Action:  lexer.Pop(),
		},
	},
	"ParamList": {
		{
			Name:    "PARAM_END",
			Pattern: `\)`,
			Action:  lexer.Pop(),
		},
		{
			Name:    "PARAM",
			Pattern: `[^,\)]+`,
		},
		{
			Name:    "PARAM_SEP",
			Pattern: `,`,
		},
	},
	"Staff": {
		{
			Name:    "STAFF_END",
			Pattern: `''!I|!t|!I`,
			Action:  lexer.Pop(),
		},
		{
			Name:    "PART_START",
			Pattern: `I!''|I!`,
		},
		{
			Name:    "SHARP",
			Pattern: `sharplg|sharpla|sharpb|sharpc|sharpd|sharpe|sharpf|sharphg|sharpha`,
		},
		{
			Name:    "NATURAL",
			Pattern: `naturallg|naturalla|naturalb|naturalc|naturald|naturale|naturalf|naturalhg|naturalha`,
		},
		{
			Name:    "FLAT",
			Pattern: `flatlg|flatla|flatb|flatc|flatd|flate|flatf|flathg|flatha`,
		},
		{
			Name:    "HALF_NOTE",
			Pattern: `LG_2|LA_2|B_2|C_2|D_2|E_2|F_2|HG_2|HA_2`,
		},
		{
			Name:    "QUARTER_NOTE",
			Pattern: `LG_4|LA_4|B_4|C_4|D_4|E_4|F_4|HG_4|HA_4`,
		},
		{
			Name:    "EIGHTH_NOTE",
			Pattern: `LG_8|LA_8|B_8|C_8|D_8|E_8|F_8|HG_8|HA_8`,
		},
		{
			Name:    "SIXTEENTH_NOTE",
			Pattern: `LG_16|LA_16|B_16|C_16|D_16|E_16|F_16|HG_16|HA_16`,
		},
		{
			Name:    "THIRTYSECOND_NOTE",
			Pattern: `LG_32|LA_32|B_32|C_32|D_32|E_32|F_32|HG_32|HA_32`,
		},
		{
			Name:    "WHOLE_NOTE",
			Pattern: `LG_1|LA_1|B_1|C_1|D_1|E_1|F_1|HG_1|HA_1`,
		},
		{
			Name:    "TIME_SIG",
			Pattern: `2_2|3_2|2_4|3_4|4_4|5_4|5_4|6_4|7_4|C_|C|2_8|3_8|4_8|5_8|6_8|7_8|8_8|9_8|10_8|11_8|12_8|15_8|18_8|21_8|2_16|3_16|4_16|5_16|6_16|7_16|8_16|9_16|10_16|11_16|12_16`,
		},
		{
			Name:    "REST",
			Pattern: `REST_16|REST_1|REST_2|REST_4|REST_8|REST_32`,
		},
		{
			Name:    "SINGLE_DOT",
			Pattern: `'lg|'la|'b|'c|'d|'e|'f|'hg|'ha`,
		},
		{
			Name:    "DOUBLE_DOT",
			Pattern: `''lg|''la|''b|''c|''d|''e|''f|''hg|''ha`,
		},
		{
			Name:    "FERMATA",
			Pattern: `fermatlg|fermatla|fermatb|fermatc|fermatd|fermate|fermatf|fermathg|fermatha`,
		},
		{
			Name:    "DOUBLING",
			Pattern: `dblg|dbla|dbb|dbc|dbd|dbe|dbf|dbhg|dbha`,
		},
		{
			Name:    "HALF_DOUBLING",
			Pattern: `hdblg|hdbla|hdbb|hdbc|hdbd|hdbe|hdbf`,
		},
		{
			Name:    "THUMB_DOUBLING",
			Pattern: `tdblg|tdbla|tdbb|tdbc|tdbd|tdbe|tdbf`,
		},
		{
			Name:    "STRIKE",
			Pattern: `strlg|strla|strb|strc|strd|stre|strf|strhg`,
		},
		{
			Name:    "G_STRIKE",
			Pattern: `gstla|gstb|gstc|gstd|lgstd|gste|gstf`,
		},
		{
			Name:    "THUMB_STRIKE",
			Pattern: `tstla|tstb|tstc|tstd|ltstd|tste|tstf|tsthg`,
		},
		{
			Name:    "HALF_STRIKE",
			Pattern: `hstla|hstb|hstc|hstd|lhstd|hste|hstf|hsthg`,
		},
		{
			Name:    "G_GRIP",
			Pattern: `ggrpla|ggrpb|ggrpc|ggrpdb|ggrpd|ggrpe|ggrpf`,
		},
		{
			Name:    "THUMB_GRIP",
			Pattern: `tgrpla|tgrpb|tgrpc|tgrpdb|tgrpd|tgrpe|tgrpf|tgrphg`,
		},
		{
			Name:    "HALF_GRIP",
			Pattern: `hgrpla|hgrpb|hgrpc|hgrpdb|hgrpd|hgrpe|hgrpf|hgrphg|hgrpha`,
		},
		{
			Name:    "GRIP",
			Pattern: `grpb|grp|hgrp|grpb`,
		},
		{
			Name:    "TAORLUATH",
			Pattern: `tarb|tar|htar`,
		},
		{
			Name:    "BUBLY",
			Pattern: `bubly|hbubly`,
		},
		{
			Name:    "BIRL",
			Pattern: `brl|abr|gbr|tbr`,
		},
		{
			Name:    "THROWD",
			Pattern: `thrd`,
		},
		{
			Name:    "HEAVY_THROWD",
			Pattern: `hvthrd`,
		},
		{
			Name:    "HALF_THROWD",
			Pattern: `hthrd`,
		},
		{
			Name:    "HEAVY_HALF_THROWD",
			Pattern: `hhvthrd`,
		},
		{
			Name:    "PELE",
			Pattern: `pella|pelb|pelc|peld|lpeld|pele|pelf`,
		},
		{
			Name:    "THUMB_PELE",
			Pattern: `tpella|tpelb|tpelc|tpeld|ltpeld|tpele|tpelf|tpelhg`,
		},
		{
			Name:    "HALF_PELE",
			Pattern: `hpella|hpelb|hpelc|hpeld|lhpeld|hpele|hpelf|hpelhg`,
		},
		{
			Name:    "DOUBLE_STRIKE",
			Pattern: `st2la|st2b|st2c|st2d|lst2d|st2e|st2f|st2hg|st2ha`,
		},
		{
			Name:    "G_DOUBLE_STRIKE",
			Pattern: `gst2la|gst2b|gst2c|gst2d|lgst2d|gst2e|gst2f`,
		},
		{
			Name:    "THUMB_DOUBLE_STRIKE",
			Pattern: `tst2la|tst2b|tst2c|tst2d|ltst2d|tst2e|tst2f|tst2hg`,
		},
		{
			Name:    "HALF_DOUBLE_STRIKE",
			Pattern: `hst2la|hst2b|hst2c|hst2d|lhst2d|hst2e|hst2f|hst2hg|hst2ha`,
		},
		{
			Name:    "TRIPLE_STRIKE",
			Pattern: `st3la|st3b|st3c|st3d|lst3d|st3e|st3f|st3hg|st3ha`,
		},
		{
			Name:    "G_TRIPLE_STRIKE",
			Pattern: `gst3la|gst3b|gst3c|gst3d|lgst3d|gst3e|gst3f`,
		},
		{
			Name:    "THUMB_TRIPLE_STRIKE",
			Pattern: `tst3la|tst3b|tst3c|tst3d|ltst3d|tst3e|tst3f|tst3hg`,
		},
		{
			Name:    "HALF_TRIPLE_STRIKE",
			Pattern: `hst3la|hst3b|hst3c|hst3d|lhst3d|hst3e|hst3f|hst3hg|hst3ha`,
		},
		{
			Name:    "D_DOUBLE_GRACE",
			Pattern: `dlg|dla|db|dc`,
		},
		{
			Name:    "E_DOUBLE_GRACE",
			Pattern: `elg|ela|eb|ec|ed`,
		},
		{
			Name:    "F_DOUBLE_GRACE",
			Pattern: `flg|fla|fb|fc|fd|fe`,
		},
		{
			Name:    "G_DOUBLE_GRACE",
			Pattern: `glg|gla|gb|gc|gd|ge|gf`,
		},
		{
			Name:    "THUMB_DOUBLE_GRACE",
			Pattern: `tlg|tla|tb|tc|td|te|tf|thg`,
		},
		{
			Name:    "SINGLE_GRACE",
			Pattern: `ag|bg|cg|dg|eg|fg|gg|tg`,
		},
		{
			Name:    "TIE_START",
			Pattern: `\^ts`,
		},
		{
			Name:    "TIE_END",
			Pattern: `\^te`,
		},
		{
			Name:    "TIE_OLD",
			Pattern: `\^tlg|\^tla|\^tb|\^tc|\^td|\^te|\^tf|\^thg|\^tha`,
		},
		{
			Name:    "IRREGULAR_GROUP_START",
			Pattern: `\^2s|\^3s|\^43s|\^46s|\^53s|\^54s|\^64s|\^74s|\^76s`,
		},
		{
			Name:    "IRREGULAR_GROUP_END",
			Pattern: `\^2e|\^3e|\^43e|\^46e|\^53e|\^54e|\^64e|\^74e|\^76e`,
		},
		{
			Name:    "TRIPLETS",
			Pattern: `\^3lg|\^3la|\^3b|\^3c|\^3d|\^3e|\^3f|\^3hg|\^3ha`,
		},
		{
			Name:    "TIMELINE_START",
			Pattern: `'224|'22|'23|'24|'intro|'25|'26|'27|'28|'1|'2`,
		},
		{
			Name:    "TIMELINE_END",
			Pattern: `_'`,
		},
		{
			Name:    "WHITESPACE",
			Pattern: `\s+`,
		},
	},
})