#!/bin/sh

source ./hack/go-mod-env.sh

if [[ -z ${CI} ]]; then
    ./hack/go-vet.sh
    ./hack/go-fmt.sh
else
    BASE_SHA=$(echo ${JOB_SPEC} | python -c "import sys, json; print(json.load(sys.stdin)['refs']['base_sha'])")
    REPO_LINK=$(echo ${JOB_SPEC} | python -c "import sys, json; print(json.load(sys.stdin)['refs']['repo_link'])")
    if [[ -z ${BASE_SHA} ]]; then
        echo "base_sha not found, can't execute diff"
        exit 1
    else
        git remote add origin ${REPO_LINK}
        git fetch origin ${BASE_SHA}
        VERSION=$(go run getversion.go)
        RESULT=$(git diff --name-only ${BASE_SHA} | grep "^config/" | grep -v "^config/${VERSION}")
        if [[ ${RESULT} ]]; then
            echo "Detected changes to an older version's config file(s). Current version changes are only allowed in config/${VERSION}."
            echo "Undo changes to the following files -"
            echo "${RESULT}"
            exit 1
        fi
    fi
fi

go test -count=1 ./...