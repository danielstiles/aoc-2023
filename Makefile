DAY?='template'

new_day:
	mkdir -p $(DAY)
	git config --get remote.origin.url | sed 's/^.*\(github.*\).git/module \1\/$(DAY)\n\ngo 1.20\n/' > $(DAY)/go.mod
	cp -R template/* $(DAY)/
