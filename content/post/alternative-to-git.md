---
title: "Alternative to git"
date: 2018-06-10T10:01:55-07:00
draft: true
tags: [thoughts]
slug: alternative-to-git
written: ["2018-06-10","2018-06","2018"]
---

Github, Gitlab, Bitbucket all control git. They all do similar things - issues and wiki, etc. *fossil* is an alternative to *git* that actually has built-in issues, wiki, and easily self-hosts. This is a quickstart to get started with hosting and cloning *fossil* repos.

First, download fossil. There are already binaries available for *fossil*, you can find them [here](https://www.fossil-scm.org/xfer/uv/download.html).

## New repo

Make sure to save each fossil repo with the extension `.fossil` - this is very important if you want to do hosting.

Now you can make a new fossil with the `init` subcommand.

```
> fossil init -A schollz 3.fossil
```

This will make a new repository with the default user `schollz`. You can then change the default password for `schollz`. Here you have to enter the password on the command-line, but I like to add it to a file to prevent it from showing up in the bash history.

```
> vim pass # enter your password in this file
> fossil user password schollz `cat pass` -R 3.fossil
> rm pass # this ensures your password doesn't enter the bash history
```

# Hosting

Say now you have several fossils in the same folder. Make sure that each has the `.fossil` suffix - this is very important.

```
> ls fossils
1.fossil 2.fossil 3.fossil
```

Then you can start a server that uses HTTPS with

```
> fossil server . --https --port 8079 --repolist
```

If you are using Caddy as a reverse-proxy (which easily adds HTTPS), your Caddyfile will look like:

```
fossil.schollz.com {
        proxy / 127.0.0.1:8079 {
                transparent
        }
        gzip
        log logs/fossil.schollz.com.log
}
```


## Cloning

Cloning is easy, just make sure to include your user name. If you are using `fossil server .` like above, then you need to add `/repo` to tell it you want to clone `repo.fossil`.

```
fossil clone https://schollz@fossil.schollz.com/3 3.fossil
```

You will be prompted for the password, which is the same password that you set up top. Now you can open the repo

```
> fossil open 3.fossil
```

Things are then pretty similar to *git*. 

```
> touch README.md
> fossil add README.md
> fossil commit README.md
> fossil push
```

## General

I like to have the `README.md` in the main repo be the default documentation. To do this, make a `README.md` file and then goto Admin -> Configuration and look fro the **Index Page** and change it to `/doc/tip/README.md`. Make sure to then press "Apply Changes" at the top.
