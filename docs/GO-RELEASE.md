# Go binary

we build go binary using go releaser

```
goreleaser release --snapshot --clean
```

or 

```
make release-check
```

export GITHUB Token
```
export GITHUB_TOKEN="YOUR_GH_TOKEN"
```

up the git tag
```
git tag v0.1.45; git push origin v0.1.45
```
