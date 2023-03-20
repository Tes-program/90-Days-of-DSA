mgs ?= "changes made"
commit_changes:
	git add .
	git commit -m "${msg}"
	git push origin main