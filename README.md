# Blink

`blink` is a fast, cross-platform command-line tool written in Go to quickly open URLs, local files, folders, and search site shortcuts directly in your default web browser.

It works seamlessly across **macOS**, **Linux**, **Windows**, and **WSL** (Windows Subsystem for Linux).

---

## Features

- 🌐 **Smart Browser Opener**: Opens standard URLs in your default system browser.
- 📁 **File & Folder Support**: Detects and opens local files (like HTML, PDFs, images) and directories using proper `file://` URIs.
- 🔍 **Site Search Aliases**: Use shortcuts to trigger search queries on popular search engines and developer tools (e.g. `blink yt "minecraft"`).
- ⚙️ **Dry Run Mode**: Resolve inputs and view the final URL/URI without launching the browser.
- 🛠️ **Zero Dependencies**: Built entirely using the Go standard library.

---

## Installation

### Prerequisites

Ensure you have [Go](https://go.dev/doc/install) installed (Go 1.24+ recommended).

### Install via Make

Clone this repository and run:

```bash
make install
```

This builds and installs the `blink` binary into your `$GOPATH/bin` directory. Make sure `$GOPATH/bin` (or `~/go/bin`) is added to your shell's `$PATH` variable.

---

## Usage

```bash
blink [flags] <target> [query...]
```

### Examples

#### 1. Open remote URLs
```bash
blink google.com
blink https://github.com
```

#### 2. Open local files and directories
```bash
blink ./index.html
blink ~/Documents/report.pdf
blink ../my-project-folder
```

#### 3. Use site search shortcuts
```bash
# Search YouTube
blink yt "golang tutorial"

# Search GitHub
blink gh "sapirrior/blink"

# Search Stack Overflow
blink so "go url parser"
```

---

## Flags

- `--help`: Show usage instructions.
- `--version`: Print the version of `blink`.
- `--list-aliases`: Show a list of all configured search aliases.
- `--dry-run`: Output the resolved URL or local file path to the terminal instead of opening it.

```bash
$ blink --dry-run yt "open source"
https://youtube.com/results?search_query=open+source
```

---

## Built-In Search Aliases

| Alias | Description | Base URL |
|---|---|---|
| `g`, `google` | Google Search | `https://google.com` |
| `yt`, `youtube` | YouTube | `https://youtube.com` |
| `gh`, `github` | GitHub | `https://github.com` |
| `gist` | GitHub Gists | `https://gist.github.com` |
| `reddit`, `r` | Reddit | `https://reddit.com` |
| `wiki`, `w`, `wikipedia` | Wikipedia | `https://en.wikipedia.org` |
| `dict`, `wiktionary` | Wiktionary | `https://en.wiktionary.org` |
| `npm` | npm Package Registry | `https://npmjs.com` |
| `pkg`, `godoc` | Go Packages | `https://pkg.go.dev` |
| `mdn` | Mozilla Developer Network | `https://developer.mozilla.org` |
| `cargo`, `crates` | Rust Crates Registry | `https://crates.io` |
| `docker`, `hub` | Docker Hub Registry | `https://hub.docker.com` |
| `aw`, `archwiki` | Arch Linux Wiki | `https://wiki.archlinux.org` |
| `tw`, `twitter`, `x` | X/Twitter | `https://x.com` |
| `maps` | Google Maps | `https://maps.google.com` |
| `translate` | Google Translate | `https://translate.google.com` |
| `ddg`, `duckduckgo` | DuckDuckGo | `https://duckduckgo.com` |
| `so`, `stackoverflow` | Stack Overflow | `https://stackoverflow.com` |
| `pypi` | Python PyPI Registry | `https://pypi.org` |
| `mvn` | Maven Central Repository | `https://search.maven.org` |

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
