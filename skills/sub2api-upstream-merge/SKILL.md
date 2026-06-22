---
name: sub2api-upstream-merge
description: Merge upstream Wei-Shaw/sub2api release tags into this repository's product-edition branch. Use only for the current sub2api project when the user asks to merge, sync, update, or pull an upstream release such as v0.1.137, including optional version synchronization, validation, commit creation, and pushing to origin/product-edition.
---

# Sub2API Upstream Merge

Use this skill only inside `/home/ubuntu/dev/sub2api`.

Prefer the bundled script over manual git commands:

```bash
node skills/sub2api-upstream-merge/scripts/merge-upstream-release.js v0.1.137
```

To push after a clean merge:

```bash
node skills/sub2api-upstream-merge/scripts/merge-upstream-release.js v0.1.137 --push
```

## Workflow

1. Confirm the request names an upstream release tag, for example `v0.1.137`.
2. Run the script from the repository root. It will:
   - require the repository root to be `/home/ubuntu/dev/sub2api`
   - require the current branch to be `product-edition`
   - require a clean working tree
   - ensure the `upstream` remote points at `Wei-Shaw/sub2api`
   - fetch the upstream tag
   - merge the tag with `--no-ff`
   - set `backend/cmd/server/VERSION` to the tag version without the leading `v`
   - commit the version bump when needed
   - run backend tests, backend build, frontend lint, frontend typecheck, critical frontend tests, and frontend build
   - push only when `--push` is passed
3. If the merge conflicts, inspect `git status` and resolve manually. Do not push until validation passes.
4. Report the merge commit, any version commit, verification commands, and push result.

## Notes

- Upstream release tags can leave `backend/cmd/server/VERSION` at the previous version. Always let the script synchronize it to the requested tag.
- The script intentionally does not run `golangci-lint`; this environment may not have it installed. Run it separately if available.
- The frontend build writes to `backend/internal/web/dist`; the script verifies that build output does not leave tracked diffs.
