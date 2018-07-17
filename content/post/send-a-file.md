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

## File synchronization

- Syncthing https://syncthing.net/
- Rclone  https://rclone.org/

## Command-line file transfer

- [Tor](#tor)
- [Dat](#dat)
- [zget](#zget)
- [Magic Wormhole](#magicwormhole)
- [toss](#toss)
- [croc](#croc)
- [cowyodel](#cowyodel)
- [netcat](#netcat)
- [IPFS](#ipfs)

## Web-based file transfer

- [transfer.sh](#transfersh)
- [ShareDrop](#sharedrop)
- [instant.io](#instantio)
- [FilePizza](#filepizza)

https://github.com/nneonneo/ffsend
https://github.com/jedisct1/piknik
netcat
ipfs

---

# IPFS {#ipfs}

IPFS is a lot more than just a file transfer utility, namely its an entire [peer-to-peer hypermedia protocol](https://ipfs.io/). 

### Install

If you don't have already, just [install the latest from their website](https://ipfs.io/docs/install/).

### Usage

The sender needs to start an IPFS daemon first and then add the file that they want to send.

```
$ ipfs daemon &
$ ipfs add somefile
added QmSr1saoM3n1Sx8dBs5bz7ozU somefile
```

The recipient can download it using anything they want - as it is accessible now with a web address.

```
$ wget https://ipfs.io/ipfs/QmSr1saoM3n1Sx8dBs5bz7ozU
```

Note, it can take a few minutes for your file to get connected between enough peers and the gateway.


---

# netcat {#netcat}

netcat allows arbitrary TCP and UDP connections and listens. It can be used for easily transfering files over LAN.

### Install

If you don't have already, just install:

```
$ apt-get install netcat pv pigz
```

The additional `pv` is a progress monitor for pipes so you can monitor the progress too!

### Usage

The usage is adapted from https://unix.stackexchange.com/a/238217

The sender should do:

```
$ cat file | pigz | nc -l 8888
```

Then the sender can send off their file

```
$ nc <sender URL> 8888 | pigz -d | pv -atep -s "1g" > file
```

---

# [cowyodel](https://github.com/schollz/cowyodel) {#cowyodel}

*cowyodel* is another one of my own file-sharing tool. This one is unique in that it uploads the file to pastebin, [cowyo.com](https://cowyo.com) which, if it is a text file, allows you to make edits. You can also share binary files though, too.

- End-to-end encryption? Optional.
- Guarantee recipient? No.
- Anonymous? Yes (except through server logs).
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? Yes, for downloading.
- Easy to install? Yes.
- Which OS? All.

### Installation 

Download a release from [the Github releases](https://github.com/schollz/cowyodel/releases/latest) or install with Go:

```
$ go get github.com/schollz/cowyodel
```

### Usage

To send a file, you first upload it to the cowyo.com server. (Note: you can actually host your own cowyo.com server, see [the Github](https://github.com/schollz/cowyo))

```
$ cowyodel upload README.md
Uploaded (textual data). Your codephrase: xx
```

You can then just download it

```
$ cowyodel download xx
```

or you can view it at `https://cowyo.com/xx`.

See the [Github](https://github.com/schollz/cowyodel#upload) for more usage details.

---

# [croc](https://github.com/schollz/croc) {#croc}

*croc* is my own file-sharing program that like [*magic wormhole*](#magicwormhole) features simple, instantaneous end-to-end encryption with guaranteed recipients ensured using a password authenticated key exchange. Unlike *magic wormhole*, though, you can just download a binary and run for any operating system - much easier to install.

- End-to-end encryption? Yes.
- Guarantee recipient? Yes.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? No.
- Easy to install? Yes.
- Which OS? All.

### Installation 

Download a release from [the Github releases](https://github.com/schollz/croc/releases/latest) or install with Go:

```
$ go get github.com/schollz/croc
```

### Usage 

You can send a file with

```
$ croc send somefile
On the other computer, please run:

croc XX
```

and receive a file with 

```
$ croc XX
```

---

# [toss](https://github.com/zerotier/toss) {#toss}

Toss is a convenient ultra-minimal command line tool to send files over LAN. One of the great features though is that you can use *toss* to [stream huge files between systems](https://github.com/zerotier/toss#stream-a-huge-archive-between-systems).

- End-to-end encryption? No.
- Guarantee recipient? Yes.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Eh, requires long addresses.
- Web accessible? No.
- Easy to install? No, requires building from C source.
- Which OS? All.

### Installation 

TODO 

### Usage 

Then you can send a file with

```
$ toss somefile 
somefile/XX
```

and receive a file with 

```
$ catch somefile/XX
```

---

# [Magic Wormhole](https://github.com/warner/magic-wormhole) {#magicwormhole}

This package provides a library and a command-line tool named *wormhole*, which makes it possible to get arbitrary-sized files and directories (or short pieces of text) from one computer to another. 

- End-to-end encryption? Yes.
- Guarantee recipient? Yes.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? Yes.
- Easy to install? No, requires installing Python.
- Which OS? All.

### Installation 

First install Python and then

```
$ pip install magic-wormhole
```

### Usage 

Then you can send a file with

```
$ wormhole send SOMEFILE
```

and receive a file with 

```
$ wormhole receive
```

which will prompt for the wormhole code from the previous.


---

# [FilePizza](https://github.com/kern/filepizza) {#filepizza}

Download/upload files using the WebTorrent protocol (BitTorrent over WebRTC). 

- End-to-end encryption? No.
- Guarantee recipient? Yes.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? Yes.
- Easy to install? Yes.
- Which OS? All.

### Usage 

Go to https://file.pizza and enter your file and then share the magnet link with someone else.


---

# [instant.io](https://github.com/webtorrent/instant.io) {#instantio}

Download/upload files using the WebTorrent protocol (BitTorrent over WebRTC). 

- End-to-end encryption? No.
- Guarantee recipient? Yes.
- Anonymous? No.
- Is it fast? Yes.
- Easy to use? Yes.
- Web accessible? Yes.
- Easy to install? Yes.
- Which OS? All.

### Usage 

Go to https://instant.io/ and enter your file and then share the magnet link with someone else.


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

First install `tor` for creating the server and `torsocks` for getting the file (although a Tor browser also works for getting the file).

```
sudo apt-get install tor torsocks
```

To send you should just create an onion and serve your files in a directory. There is a simple wrapper in Go that can do this. You can get [download a release](https://github.com/schollz/onionserve/releases/latest) or install it with Go:

```
$ go get github.com/schollz/onionserve
```

### Usage

To share a file, you can just serve your directory with the `onionserve` program.

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
