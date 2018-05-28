---
title: "My Favorite Vimrc"
date: 2018-05-28T10:25:45-07:00
tags: [coding]
slug: my-favorite-vimrc
written: ["2018-05-28","2018-05","2018"]
---

I've been enjoying writing in `vim` a lot more. There is a nice and simple way to make writing a lot easier (word-wrapping, nice color scheme). You can just make the following config file (save as `~/.vimrc`):

```
func! WordProcessorMode() 
  set columns=80
  setlocal formatoptions=1 
  setlocal noexpandtab 
  map j gj 
  map k gk
  setlocal spell spelllang=en_us 
  set complete+=s
  set formatprg=par
  setlocal wrap 
  setlocal linebreak 
  set foldcolumn=3
  highlight Normal ctermfg=black ctermbg=grey
  hi NonText ctermfg=grey guifg=grey
endfu 
com! WP call WordProcessorMode()
```

Then, while in `vim` you can activate the word processing mode by typing `:WP`.


