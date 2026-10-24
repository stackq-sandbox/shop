# shop

Sandbox repository for exercising [stackq](https://github.com/shresthaujjwal/stackq)
against live GitHub: stacked PRs, the merge queue, atomic landing and priority.

The Go code is a small fake storefront. It exists so there is something real to
build, vet and test, and so a large branch (`feature/new-checkout`) has a
believable shape to split.

No personal access token is used. The `stackq` workflow runs with the built-in
`GITHUB_TOKEN`, and stackq itself is vendored under `tools/stackq` so the
workflow does not need to fetch another (private) repository.
