---
title: "Hugo Craft"
date: 2018-05-26T23:12:08-06:00
tags: [thoughts]
slug: hugocraft
written: ["2018-05-26","2018-05","2018"]
---

This is a ready-to-go [Hugo](https://gohugo.io/)-based blog publishing system. The styling has a lot of inspiration from [the blog of Chris Siebenmann](https://utcc.utoronto.ca/~cks/space/blog/BlogGenesis) which I think excels in simplicity and minimalism. It uses [mage](https://github.com/magefile/mage) for an OS-agnostic build system to make it easy to chain commands. It has a simple search and comments ready-to-go.

## Features

- Integrated search ([try it](https://hugocraft.schollz.com/search/?s=cat))
- Integrated comments [at the bottom of each post](https://hugocraft.schollz.com/my-first-post/)
- Automatic image captioning using markdown
- Very fast (base page ~3.9kB)
- Easy to get started and run

## Get started

You need to have Go installed. Then install Mage:

```
$ go get github.com/magefile/mage
```

The install Hugo by just downloading the [latest release](https://github.com/gohugoio/hugo/releases/latest).

## How to use

To get started, you can make a new post:

```
$ mage new
```

You'll be prompted for a name and then the post will be created. You can open the post in your favorite editor to edit it. Or, if you like using `vim` you can just do 

```
$ mage write
```

While you are editing you can watch the changes in the blog by just doing

```
$ mage
```

And when you are ready to publish just do

```
$ mage publish
```

which will build and minify the results into the `tmp/` directory.

# License

MIT