// +build mage

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

// Serve the hugo site
func Serve() (err error) {
	err = sh.RunV("hugo", strings.Fields("server --ignoreCache -D --watch -t twotwothree --bind 0.0.0.0 --enableGitInfo --disableFastRender")...)
	return
}

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
	return
}

// Publish will publish and minify the blog
func Publish() (err error) {
	fmt.Println("generating site...")
	err = sh.RunV("hugo", "-t", "twotwothree")
	if err != nil {
		return
	}
	fmt.Println("minifying...")
	err = sh.RunV("minify", "-a", "-r", "-o", "tmp/", "tmp")
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
