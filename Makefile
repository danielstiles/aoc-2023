YEAR?=2023
DAY?='template'

new_day:
	mkdir -p $(DAY)
	cp -R template/* $(DAY)/
	git config --get remote.origin.url | sed 's/^.*\(github.*\).git/module \1\/$(DAY)\n\ngo 1.21.5\n/' > $(DAY)/go.mod
	curl https://adventofcode.com/$(YEAR)/day/$(DAY)/input -H "Cookie: $(shell cat cookie)" > $(DAY)/input.txt
