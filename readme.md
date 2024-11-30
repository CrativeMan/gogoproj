> [!WARN] VLANG SUPPORT
> Vlang support is currently not working as indended

# GOGO

A simple project creator intended for personal use. Built in [golang](https://go.dev) and using the [huh](https://github.com/charmbracelet/huh) library.
Gogo is built for the nix(flake) workflow and builid process.
It uses the [nix dev templates](https://github.com/the-nix-way/dev-templates) for handeling the flake.nix and flake.lock files.

## Supported languages

Gogo supports at the moment:
- C
- C++
- Golang
- Rust (partially)
- Vlang

When creating a Rust project, gogo only initializes direnv and the flake files. The project setup has to be done manually for example with cargo.

## How to install
### Manual install

```bash
git clone https://github.com/CrativeMan/gogoproj gogo
cd gogo
make && make run
```

### Nix/Flake install
Add 
```nix
gogo = {
	url = "github:CrativeMan/gogo";
};
```
to your flake inputs, and then in your system packages add this.
```nix
environment.systemPackages = [
	pkgs.inputs.gogo.packages.x86_64-linux.default
}
```
After that, update your flake and rebuild your system configuration and it should be installed.
