Wally fetches random featured images from [deprecated](https://changelog.unsplash.com/deprecations/2021/11/25/source-deprecation.html) Unsplash Source API, and set your desktop wallpaper.

Wally currently works only on macOS. PRs are welcome.

## Installation 

### Using Homebrew

```
brew tap ltpquang/tap
brew install wally
```

### Using pre-built binaries

Download proper binary from [releases section](https://github.com/ltpquang/wally/releases) and put into your paths.

### Build from source

Alternative, you can clone this repo and install the binary yourself, `go` is required

```
git clone https://github.com/ltpquang/wally
cd wally
go install
```

## Usage

From your Terminal, simply

```
wally
```

you can also provide additional search query to filter only desired topics

```
wally cats
```
