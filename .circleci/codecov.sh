#!/bin/bash -e

COVERAGE_FILE=coverage.txt
PROFILE_OUT=profile.out

echo "" > ${COVERAGE_FILE}

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=${PROFILE_OUT} -covermode=atomic -tags test $d
    if [[ -f ${PROFILE_OUT} ]]; then
        cat ${PROFILE_OUT} >> ${COVERAGE_FILE}
        rm ${PROFILE_OUT}
    fi
done

bash <(curl -s https://codecov.io/bash)
