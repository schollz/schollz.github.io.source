// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/sh"
	"github.com/spf13/viper"
)

const VIMRC = `set nocompatible
set backspace=2
autocmd TextChanged,TextChangedI <buffer> silent write
func! WordProcessorModeCLI()
  set columns=80
  setlocal formatoptions=1 
  setlocal noexpandtab 
  map j gj 
  map k gk
  set complete+=s
  set formatprg=par
  setlocal wrap 
  setlocal linebreak 
  set foldcolumn=3
  highlight Normal ctermfg=black ctermbg=grey
  hi NonText ctermfg=grey guifg=grey
  hi Folded guibg=grey
endfu
com! WPCLI call WordProcessorModeCLI()
au BufEnter * set noro
`

// Serve the hugo site.
func Serve() (err error) {
	err = sh.RunV("hugo", strings.Fields("server --ignoreCache -D --watch -t twotwothree --bind 0.0.0.0 --enableGitInfo --disableFastRender")...)
	return
}

var NewFile string

// New creates a new post.
func New() (err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Post title: ")
	slug, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	slug = strings.Join(strings.Fields(strings.ToLower(slug)), "-")
	postName := "post/" + slug + ".md"
	err = sh.RunV("hugo", "new", postName)
	fmt.Printf("created post 'content/%s'", postName)
	NewFile = "content/" + postName
	return
}

// Publish will publish and minify the blog.
func Publish() (err error) {
	// get config.toml settings
	configBytes, _ := ioutil.ReadFile("config.toml")
	viper.SetConfigType("toml") // or viper.SetConfigType("YAML")
	viper.ReadConfig(bytes.NewBuffer(configBytes))
	githubPublish := viper.Get("githubPublish")
	if githubPublish != nil {
		if _, errExists := os.Stat("tmp"); os.IsNotExist(errExists) {
			// clone github to tmp
			fmt.Println("getting the current repo...")
			err = sh.RunV("git", "clone", "git@github.com:"+githubPublish.(string), "tmp")
			if err != nil {
				return
			}
		} else {
			os.Chdir("tmp")
			fmt.Println("pulling the current repo...")
			err = sh.RunV("git", "pull")
			if err != nil {
				return
			}
			os.Chdir("..")
		}
	}

	fmt.Println("generating site...")
	err = sh.RunV("hugo", "-t", "twotwothree")
	if err != nil {
		return
	}
	fmt.Println("getting minifier...")
	err = sh.RunV("go", "get", "-v", "github.com/tdewolff/minify/cmd/minify")
	if err != nil {
		return
	}
	err = sh.RunV("minify", "-a", "-r", "-o", "tmp/", "tmp")
	if err != nil {
		return
	}

	if githubPublish != nil {
		// commit the changes
		os.Chdir("tmp")
		defer os.Chdir("..")
		err = sh.RunV("git", "add", ".")
		if err != nil {
			return
		}
		var status string
		status, err = sh.Output("git", "status", "-s")
		if err != nil {
			return
		}
		modified := strings.Count(status, "M ")
		added := strings.Count(status, "A ")
		commitMessage := fmt.Sprintf("add %d and modify %d files\n\n%s", added, modified, status)
		err = sh.RunV("git", "commit", "-am", commitMessage)
		if err != nil {
			return
		}
		err = sh.RunV("git", "push", "origin", "master")
		if err != nil {
			return
		}
	}
	return
}

// Write will create a new post and open up vim to start writing.
func Write() (err error) {
	New() // New can fail
	dir, fname := filepath.Split(NewFile)
	fmt.Println(path.Join(dir, "."+fname+".swp"))
	os.Remove(path.Join(dir, "."+fname+".swp"))
	go sh.Run("hugo", strings.Fields("server --ignoreCache -D --watch -t twotwothree --bind 0.0.0.0 --enableGitInfo --disableFastRender")...)
	vimrc, err := ioutil.TempFile("", "vimrc")
	if err != nil {
		return
	}
	defer os.Remove(vimrc.Name()) // clean up
	err = ioutil.WriteFile(vimrc.Name(), []byte(VIMRC), 0755)
	if err != nil {
		return
	}

	vimExecutable := "vim"
	if runtime.GOOS == "windows" {
		vimExecutable = "vim.exe"
	}
	err = sh.RunV(vimExecutable, "-u", vimrc.Name(), "-c", "WPCLI", "+", "+startinsert", NewFile)
	return
}

// Push will automatically push the current content.
func Push() (err error) {
	err = sh.RunV("git", "add", "content/")
	if err != nil {
		return
	}
	err = sh.RunV("git", "commit", "-am", "updated")
	if err != nil {
		return
	}
	err = sh.RunV("git", "push")
	if err != nil {
		return
	}
	return
}

// Update will fetch the latest from hugocraft.
func Update() (err error) {
	err = sh.RunV("git", "remote", "add", "upstream", "https://github.com/schollz/hugocraft.git")
	if err != nil {
		return
	}
	err = sh.RunV("git", "fetch", "upstream")
	if err != nil {
		return
	}
	err = sh.RunV("git", "merge", "upstream/master")
	if err != nil {
		return
	}
}

// A var named Default indicates which target is the default.
var Default = Serve
