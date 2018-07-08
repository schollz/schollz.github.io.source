# hugocraft

This is a ready-to-go [Hugo](https://gohugo.io/)-based blog publishing system. The styling has a lot of inspiration from [the blog of Chris Siebenmann](https://utcc.utoronto.ca/~cks/space/blog/BlogGenesis) which I think excels in simplicity and minimalism. It uses [mage](https://github.com/magefile/mage) for an OS-agnostic build system to make it easy to chain commands. It has a simple search and comments ready-to-go.

## Features

- Integrated search ([try it](https://hugocraft.schollz.com/search/?s=blog))
- Integrated comments [at the bottom of each post](https://hugocraft.schollz.com/hugocraft/#commentssection)
- Automatic image captioning using markdown
- Very fast (base page ~3.9kB)
- Easy to get started and run

## Get started

You need to have Go installed. Then install dependencies:

```
$ go get github.com/magefile/mage
$ go get github.com/spf13/viper
```

The install Hugo by just downloading the [latest release](https://github.com/gohugoio/hugo/releases/latest).

Now you should start a new repo and add hugocraft.

```
$ mkdir my-blog && cd my-blog
$ git init 
$ git remote add upstream https://github.com/schollz/hugocraft.git
$ git fetch upstream
$ git merge upstream/master
```

Now you are ready to use!

## How to use

### Writing 

To get started, you can make a new post:

```
$ mage new
```

You'll be prompted for a name and then the post will be created. You can open the post in your favorite editor to edit it. 

Or, if you like using `vim` you can just do 

```
$ mage write
```

When you want to push your latest changes just do

```
$ mage push
```

### IPFS publishing

First install `ipfs` and make sure you have `wget`. Run your *ipfs* instance (`ipfs daemon`) Then you can easily publish to IPFS using

```
$ mage ipfs
```

which will return your hash. Now you can view your site at `https://ipfs.ip/ipfs/<your hash>`.

### Github publishing

And when you are ready to publish you need to set your repo name. Goto the `config.toml` and just change `githubPublish` to your Github repo, e.g. `user/user.github.io`. Then you can publish by just doing

```
$ mage publish
```

which will build and minify the results into the `tmp/` directory.

### Get latest *hugocraft*

After you've forked this repo, you can update with 

```
$ mage update
```

which will do the following commands to merge the upstream *hugocraft* fork:

```
git remote add upstream https://github.com/schollz/hugocraft.git
git fetch upstream
git merge upstream/master
```

# License

MIT
