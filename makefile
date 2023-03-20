all: add_changes commit_changes push_changes
mgs ?= "changes made"
add_changes:
	git add .
commit_changes:
	git commit -m "${msg}"
push_changes:
	git push origin busta