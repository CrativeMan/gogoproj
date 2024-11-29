package util

var TypeIntToStr = map[int]string{
	C:    "c-cpp",
	CPP:  "c-cpp",
	GO:   "go",
	RUST: "rust",
}

const (
	C    = 0
	CPP  = 1
	GO   = 2
	RUST = 3
)

var Flake []string = []string{`nix`, `flake`, `init`, `--template`, `https://flakehub.com/f/the-nix-way/dev-templates/*#c-cpp`}
