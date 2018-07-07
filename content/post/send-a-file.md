---
title: "Send a File"
date: 2018-05-21T17:36:50-07:00
draft: true
tags: [thoughts]
slug: send-a-file
written: ["2018-05-21","2018-05","2018"]
---

Is it secure? This question really is a bunch of questions together. Security that is important is end-to-end encryption and guaranteed that the right recipient(s) receives the file. 

Is it fast? This is probably the most important question. Basically the speed will be rate-limited by the network, but a fast file transfer should utilize the network as much as possible.

Is it easy to use (e.g. for an average person and not a CS major)? I'm going to discuss some interesting ways to send files that require specific OS and programs, which would not be easy for the average person.

Which OS? There are some utilities that are really just specific to linux operating systems.

# Table of contents

## Command-line 

- [Tor](#tor)
- [Dat](#dat)
- [zget](#zget)

## Web-based

- [transfer.sh](#transfersh)
- [ShareDrop](#sharedrop)

https://github.com/webtorrent/instant.io
https://github.com/kern/filepizza
https://github.com/warner/magic-wormhole
https://github.com/zerotier/toss
https://github.com/datproject/dat
https://github.com/schollz/croc
https://github.com/nneonneo/ffsend
https://github.com/jedisct1/piknik

---

# [sharedrop](https://github.com/cowbell/sharedrop) {#sharedrop}

ShareDrop is an HTML5 clone of Apple's AirDrop which can provide easy P2P file transfer powered by WebRTC. It is [open source](https://github.com/cowbell/sharedrop). 

- End-to-end encryption? No.
- Guarantee recipient? Yes.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? Yes.
- Easy to install? Yes.
- Which OS? All.

### Usage 

Go to https://www.sharedrop.io/. It will provide you a link for a specific room where you can tell the recipient to go. Once both there you can exchange files.

---

# [zget](https://github.com/nils-werner/zget) {#zget}

A simple [open source](https://github.com/nils-werner/zget) Zeroconf-based, peer to peer file transfer utility, where you transfer files by just relating the filename. *This tool requires that you are on the same network,* so it is not useful in most situations.

- End-to-end encryption? No.
- Guarantee recipient? No, anyone with the filename could download.
- Anonymous? Yes.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? No.
- Easy to install? No, requires installing Python.
- Which OS? All.

### Installation

```
$ pip install zget
```

### Usage 

To send a file just use `zput`:

```
$ zput somefile
```

and then anyone else can get the file

```
$ zget somefile
```

---

# [transfer.sh](https://github.com/dutchcoders/transfer.sh) {#transfersh}

transfer.sh is a neat [open-source](https://github.com/dutchcoders/transfer.sh/) file transfer server. You can use a command-line tool or the web interface to upload/download files.

- End-to-end encryption? No.
- Guarantee recipient? No, anyone with the link could download.
- Anonymous? Yes.
- Is it fast? No, requires upload and subsequent download.
- Easy to use? Yes.
- Web accessible? Yes.
- Easy to install? Yes, its just a webpage.
- Which OS? All.

---

# [Dat](https://github.com/datproject/dat) {#dat}

Dat is an alternative to [IPFS](#ipfs). Currently only supports Mac and Linux, but there is lots of growth.

- End-to-end encryption? Yes.
- Guarantee recipient? No, anyone with the *dat* link could clone.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? No.
- Easy to install? No, it requires installing Node.
- Which OS? All.

### Installation

```
$ npm install -g dat
```

### Usage

To send a file you can just go to the directory that you want and share it with

```
$ dat share
Created new dat in ...
dat://XX
```

Now you can share the `dat://XX` with someone and they can get your data very easily by doing

```
$ dat clone dat://XX
```

The really cool thing about *dat* is that you can keep files in sync. After you clone you can cd into the directory and issue a sync command.

```
$ cd XX
$ dat sync
```

Now any changes from the sender will be automatically updated here.

---

# Tor {#tor}

Tor is nice because in theory it is encrypted and anonymous file transfer. You need to specify the onion handle, but it can easily be created and destroyed relatively quickly.

- End-to-end encryption? Yes.
- Guarantee recipient? No, anyone with the onion link can download.
- Anonymous? Yes.
- Is it fast? No, it takes a few minutes to create an onion.
- Easy to use? Yes.
- Web accessible? Yes.
- Easy to install? Yes.
- Which OS? Linux and OSX.

### Installation 

```
sudo apt-get install tor torsocks
```

### Usage

To send you should just create an onion and serve your files in a directory. There is a simple wrapper in Go that can do this. You can get it with Go `go get -u -v github.com/schollz/onionserve` or [download a release](https://github.com/schollz/onionserve/releases/latest).

```
$ onionserve
Starting and registering onion service, please wait a couple of minutes...
Open Tor browser and navigate to http://XX.onion
Press enter to exit
```

The receiver can browse the files using the Tor-browser or you can download via the command-line if you know the file name:

```
$ torsocks wget http://XX.onion/file 
```

You can also browse all the files at `http://XX.onion` using a Tor browser.