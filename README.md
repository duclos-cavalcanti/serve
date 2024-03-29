<h1 align="center">Serve</h1>
<p align="center">
  Menu CLI tool
</p>

A very simple tool that receives a given number of options, displays them in a TUI menu and enables the selection of one of them. The
user can choose between the options through vim keys (`jk`) and finally select an option through `Enter`. The selected option will then
be written to `stdout`.

## Usage
Serve can read a given string through the `-o` or `--options` flag and split it into options, which by default is separated through `\n` characters. Different delimiters can be chosen through
the `-d` option. Additionally, serve recognizes when it is used in command chaining or piping and then proceeds to read the options from `stdin`. When
reading from `stdin` however, serve can currently only delimit options through newlines.

Options:
```sh
usage: serve [options]
    -m, --mode      mode in which serve will execute, currently only one
    -o, --options   options to select from
    -p, --prompt    menu prompt
    -d, --delimiter string or char used to separate options, default is '\n'
    -h, --help      print usage
```

Example:
```sh
serve --prompt PROMPT --options "foo\nbar\nbaz"
```

<div align="center">

![usage](./.github/assets/demo.gif?)

</div>

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

## License
This project is released under the MIT license 3.0. See [LICENSE](LICENSE).

## Contributions
Please follow the instructions in the contributions guide at [CONTRIBUTING.md](CONTRIBUTING.md).

## Thanks
- [dmenu](http://tools.suckless.org/dmenu/)
- [tcell](https://github.com/gdamore/tcell)
- [go-project-example](https://github.com/albertwidi/go-project-example)
- [vhs](https://github.com/charmbracelet/vhs)

## Donations
I have a ko-fi and a buy-me-a-coffee account, so if you found this repo useful and would like to show your appreciation, feel free to do so!

<p align="center">
<a href="https://ko-fi.com/duclos">
<img src="https://img.shields.io/badge/donation-ko--fi-red.svg">
</a>

<a href="https://www.buymeacoffee.com/danielduclos">
<img src="https://img.shields.io/badge/donation-buy--me--coffee-green.svg">
</a>

</p>

---
<p align="center">
<a href="https://github.com/duclos-cavalcanti/templates/LICENSE">
  <img src="https://img.shields.io/badge/license-MIT-blue.svg" />
</a>
<a>
  <img src="https://img.shields.io/github/languages/code-size/duclos-cavalcanti/serve.svg" />
</a>
<a>
  <img src="https://img.shields.io/github/commit-activity/m/duclos-cavalcanti/serve.svg" />
</a>

<!-- --- -->
<!-- <a href="https://ko-fi.com/duclos" target="_blank"><img src="https://ko-fi.com/img/githubbutton_sm.svg" alt="Support me on Ko-fi" style="height: 30px !important;" ></a> -->
<!-- <a href="https://www.buymeacoffee.com/danielduclos" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/lato-green.png" alt="Buy Me A Coffee" style="height: 30px !important;" ></a> -->
<!-- <a href="https://www.paypal.com/donate/?hosted_button_id=NHPMH2UR93APC"> <img alt="Support via PayPal" style="height: 30px !important;" src="https://cdn.rawgit.com/twolfson/paypal-github-button/1.0.0/dist/button.svg"/> </a> -->
