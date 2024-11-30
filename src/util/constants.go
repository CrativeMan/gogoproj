package util

const (
	C     = 0
	CPP   = 1
	GO    = 2
	RUST  = 3
	VLANG = 4
)

var TypeIntToStr = map[int]string{
	C:     "c-cpp",
	CPP:   "c-cpp",
	GO:    "go",
	RUST:  "rust",
	VLANG: "vlang",
}

var TypeExtension = map[int]string{
	C:     ".c",
	CPP:   ".cpp",
	GO:    ".go",
	RUST:  ".rs",
	VLANG: ".v",
}

var TemplatesMain = map[int]string{
	C:     "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/mainC",
	CPP:   "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/mainCpp",
	GO:    "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/mainGo",
	RUST:  "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/mainRust",
	VLANG: "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/mainV",
}

var TemplatesMakefile = map[int]string{
	C:     "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/makefileC",
	CPP:   "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/makefileCpp",
	GO:    "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/makefileGo",
	RUST:  "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/makefileRust",
	VLANG: "https://raw.githubusercontent.com/CrativeMan/gogoproj/refs/heads/master/static/makefileV",
}

var Flake []string = []string{`nix`, `flake`, `init`, `--template`, `https://flakehub.com/f/the-nix-way/dev-templates/*#c-cpp`}
