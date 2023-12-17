YEAR?=2023
DAY?='template'
GO_VERSION?=1.21.5

new_day:
	mkdir -p $(DAY)
	cp -R template/* $(DAY)/
	git config --get remote.origin.url | sed 's/^.*\(github.*\).git/module \1\/$(DAY)\n\ngo $(GO_VERSION)\n/' > $(DAY)/go.mod
	curl https://adventofcode.com/$(YEAR)/day/$(DAY)/input -H "Cookie: $(shell cat cookie)" > $(DAY)/input.txt
