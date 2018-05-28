// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/spf13/viper"
)

const VIMRC = `set nocompatible
set backspace=2
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

// Serve the hugo site
func Serve() (err error) {
	err = sh.RunV("hugo", strings.Fields("server --ignoreCache -D --watch -t twotwothree --bind 0.0.0.0 --enableGitInfo --disableFastRender")...)
	return
}

var NewFile string

// New creates a new post
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

// Publish will publish and minify the blog
func Publish() (err error) {
	os.RemoveAll("tmp")
	// get config.toml settings
	configBytes, _ := ioutil.ReadFile("config.toml")
	viper.SetConfigType("toml") // or viper.SetConfigType("YAML")
	viper.ReadConfig(bytes.NewBuffer(configBytes))
	githubPublish := viper.Get("githubPublish")
	if githubPublish != nil {
		// clone github to tmp
		fmt.Println("getting the current repo...")
		err = sh.RunV("git", "clone", "git@github.com:"+githubPublish.(string), "tmp")
		if err != nil {
			return
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

// Write will open up vim to start writing
func Write() (err error) {
	mg.Deps(New)
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

// Update will automatically update
func Update() (err error) {
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

// The first sentence in the comment will be the short help text shown with mage -l.
// The rest of the comment is long help text that will be shown with mage -h <target>
func Target() {
	// by default, the log stdlib package will be set to discard output.
	// Running with mage -v will set the output to stdout.
	log.Printf("Hi!")
}

// A var named Default indicates which target is the default.
var Default = Serve
