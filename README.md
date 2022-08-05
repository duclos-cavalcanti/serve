# Serve
<a href="https://github.com/duclos-cavalcanti/serve/LICENSE">
  <img src="https://img.shields.io/badge/license-MIT-blue.svg" />
</a>

## Introduction
A very simple tool that receives a given number of options, displays them in a TUI menu and enables the selection of one of them. The
user can choose between the options through vim keys (`jk`) and finally select an option through `Enter`. The selected option will then
be written to stdout.

## Usage
```sh
serve --prompt PROMPT --options "foo\nbar\nbaz"
```
![usage](./.assets/usage.gif)

## Dependencies
It simply uses Go's built in libraries and the amazing [tcell](https://github.com/gdamore/tcell) library. So to be able to compile the project, it is only needed to have `Go` installed on your system
### Arch-based
```sh
sudo pacman -S go
```

### Debian-based
```sh
sudo apt install go
```

## Installation
### Steps
1. `make build`: pulls necessary dependencies and builds the binary
2. `make install`: installs it onto your system such that it is visible in $PATH (not complete)

## Thanks
- [dmenu](http://tools.suckless.org/dmenu/)
- [tcell](https://github.com/gdamore/tcell)
- [go-project-example](https://github.com/albertwidi/go-project-example)

## License
Serve is released under the MIT license. See [LICENSE](LICENSE).

## Contributions
Please follow the instructions in the contributions guide at [CONTRIBUTING.md](CONTRIBUTING.md).

### ToDo's
- [ ] make better fig
- [ ] implement new modes, maybe categorized options
- [ ] implement `install` target
- [ ] implement `format` target

---
<a href="https://www.buymeacoffee.com/danielduclos" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/lato-green.png" alt="Buy Me A Coffee" style="height: 40px !important;" ></a>
<a href="https://www.paypal.com/donate/?hosted_button_id=NHPMH2UR93APC"> <img alt="Support via PayPal" style="height: 40px !important;" src="https://cdn.rawgit.com/twolfson/paypal-github-button/1.0.0/dist/button.svg"/> </a>
