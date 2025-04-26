.PHONY: sync
sync:
	@echo "Syncing files..."
	rm -rf lessons/
	cd tool/getlessons; \
		go run .
	@echo "Syncing README..."
	rm -f README.md
	cp README.md.template README.md
	cd tool/listexercises; \
		go run . >> ../../README.md
	@echo "Complete."