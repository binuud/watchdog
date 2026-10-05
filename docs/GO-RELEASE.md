# Go binary

we build go binary using go releaser

```
goreleaser release --snapshot --clean
```

or 

```
make release-check
```

## Release GO binaries

export GITHUB Token
```
export GITHUB_TOKEN="YOUR_GH_TOKEN"
```

up the git tag
```
git tag v0.1.45; git push origin v0.1.45
```

Build the binary for windows, darwin, linux on arm and x86 arch, and push same to github
```
make release
```