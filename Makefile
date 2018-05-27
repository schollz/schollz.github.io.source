serve:
	hugo server --ignoreCache -D --watch -t twotwothree --bind 0.0.0.0 --enableGitInfo --disableFastRender

build:
	hugo -t twotwothree
	minify -a -r -o tmp/ tmp

publish:
	#go get -v github.com/tdewolff/minify/cmd/minify
	git status -s
	hugo -t twotwothree
	minify -a -r -o tmp/ tmp
	git checkout master 
	git pull origin master
	rsync -avrP tmp/ ./
	rm -rf tmp
	rm -rf themes
	rm -rf content
	rm -rf static
	rm -rf archetypes
	git add .
	git commit -am "Update site"
	git push origin master
	git checkout simplify
	git clone git@github.com:schollz/twotwothree.git themes/twotwothree
	cd themes/twotwothree && git checkout two

update:
	git add content/post/
	git commit -am "Update posts"
	git push origin source
	

# serve:
# 	hugo server --disableLiveReload --bind 0.0.0.0 -t twotwothree -b http://blog.schollz.com --appendPort=false

