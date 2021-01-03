# wally
This is a fork of [Go Si Mac](https://github.com/1995parham/gosimac), initially created by [1995parham](https://github.com/1995parham). It extends _Go Si Mac_ with images from [reddit](https://www.reddit.com/r/wallpaper/hot/) and is capable of setting a random background image for you.

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/quickstar/wally/lint.yaml?label=lint&logo=github&style=flat-square&branch=main)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/quickstar/wally?logo=github&style=flat-square)
![GitHub Release Date](https://img.shields.io/github/release-date/quickstar/wally?logo=github&style=flat-square)

[![AUR package](https://repology.org/badge/version-for-repo/aur/gosimac.svg?style=flat-square)](https://repology.org/project/gosimac/versions)

## Introduction

*wally* downloads Bing's daily wallpapers, Unsplash's and Reddit's hot random images for you to have a beautiful wallpaper on your desktop whenever you want.

## Installation
### brew
```
brew install 1995parham/tap/gosimac
```

## Usage

```bash
wally rev-4cbe101-dirty
Fetch the wallpaper from Bing, Reddit, Unsplash...

Usage:
  wally [command]

Available Commands:
  bing        fetches images from https://bing.com
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  unsplash    fetches images from https://unsplash.org
  reddit      fetches images from https://www.reddit.com/r/wallpaper/hot

Flags:
  -h, --help          help for wally
  -n, --number int    The number of photos to return (default 10)
  -p, --path string   A path to store the photos (default "$HOME/Pictures/wally")
  -v, --version       version for wally
```

As an example, the following command downloads 10 images from unsplash while using _deer_ as a search query.

```sh
wally u -q deer -n 10
```

By default, *wally* stores images in `$HOME/Pictures/wally`.

## Contribution

The `Init` method is called on initiation and returns number of available images to download.
Then for each image `Fetch` is called and the result is stored at the user specific location.
By implementing this interface you can create new sources for *wally*.

For adding new source you only need to create a new sub-command in `cmd` package
and then calling your new source with provided `path`. Also for saving images
you can use the following helper function:

```go
func (u *Unsplash) Store(name string, content io.ReadCloser) {
        path := path.Join(
                u.Path,
                fmt.Sprintf("%s-%s.jpg", u.Prefix, name),
        )

        if _, err := os.Stat(path); err == nil {
                pterm.Warning.Printf("%s is already exists\n", path)

                return
        }

        file, err := os.Create(path)
        if err != nil {
                pterm.Error.Printf("os.Create: %v\n", err)

                return
        }

        bytes, err := io.Copy(file, content)
        if err != nil {
                pterm.Error.Printf("io.Copy (%d bytes): %v\n", bytes, err)
        }

        if err := file.Close(); err != nil {
                pterm.Error.Printf("(*os.File).Close: %v", err)
        }

        if err := content.Close(); err != nil {
                pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
        }
}
```
