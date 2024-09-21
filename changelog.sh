LATEST_TAG=$(git describe --tags --abbrev=0 HEAD^)
REPO_URL=$(git config --get remote.origin.url)
REPO_URL=${REPO_URL%.git}
REPO_URL=${REPO_URL#git@}
REPO_URL=${REPO_URL#https://}
REPO_URL=${REPO_URL/:/\/}
REPO_URL="https://${REPO_URL}"
git log --pretty=format:"- [%h](${REPO_URL}/commit/%H) %s" ${LATEST_TAG}..HEAD