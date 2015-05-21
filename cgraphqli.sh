#!/bin/bash
# cgraphqli.sh
#
# requirements:
# jq - for output formatting
# rlwrap - optional requirement for readline support (command history)

set -eou pipefail
BASEURL=${1:-"http://localhost:8080/"}

function query(){
	local q=${1:-"{ __schema{ root_fields { name } } }"}
	echo curl -sgG --data-urlencode "\"q=${q}\"" "${BASEURL}"
	curl -sgG --data-urlencode "q=${q}" "${BASEURL}" | jq .
	echo
}

function maybe_rlwrap() {
	if [ "${USING_RLWRAP:-0}" -eq "0" ] && which -s rlwrap; then
		USING_RLWRAP=1 exec rlwrap "$0" "$@"
	fi
}
function main(){
	maybe_rlwrap "$@"
	echo "using base url '${BASEURL}'"
	while read line; do
		query "$line"
	done
}

main "$@"
