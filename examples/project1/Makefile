CONFLUENCE_ANCESTOR_PAGE=1719369746
CONFLUENCE_SPACE_KEY=BT

.PHONY: test
test:
	go run ./scripts/integration-tests/main.go


.PHONY: publish
publish-to-confluence:
	gomplate -d results.json -f results.tmpl > results.html
	cfcli -u ${CONFLUENCE_USER} -p ${CONFLUENCE_TOKEN} -s ${CONFLUENCE_SITE}  \
		-k ${CONFLUENCE_SPACE_KEY} -f results.html -a ${CONFLUENCE_ANCESTOR_PAGE} \
		-t "database test results" add-or-update-page

