# Releasing

Releases for the Go ecosystem and on GitHub are done via git tags:

1. `git checkout main` to get on the main branch
2. `git tag --sort=-creatordate` to see last created tags
3. `git tag vX.X.X` (e.g. `git tag v0.1.5`) to create a new tag
4. `git push --tags` to push the tags to the remote

The [`cd`](.github/workflows/cd.yml) workflow will be triggered on tag push. It will build the binaries and create a new release with the binaries uploaded.