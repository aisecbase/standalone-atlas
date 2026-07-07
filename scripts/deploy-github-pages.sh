#!/usr/bin/env bash
set -euo pipefail

remote="${GITHUB_PAGES_REMOTE:-origin}"
branch="${GITHUB_PAGES_BRANCH:-gh-pages}"
message="${GITHUB_PAGES_COMMIT_MESSAGE:-Deploy ATLAS RU site}"

if ! git remote get-url "$remote" >/dev/null 2>&1; then
  echo "Git remote '$remote' is not configured." >&2
  echo "Set GITHUB_PAGES_REMOTE or add a remote before deploying." >&2
  exit 1
fi

REQUIRE_ATLAS_RELEASE_TAG="${REQUIRE_ATLAS_RELEASE_TAG:-1}" scripts/build-site.sh

if [[ ! -f public/index.html ]]; then
  echo "public/index.html was not generated." >&2
  exit 1
fi

tmpdir="$(mktemp -d)"
cleanup() {
  git worktree remove --force "$tmpdir" >/dev/null 2>&1 || rm -rf "$tmpdir"
}
trap cleanup EXIT

if git ls-remote --exit-code --heads "$remote" "$branch" >/dev/null 2>&1; then
  git fetch "$remote" "$branch"
  git worktree add --detach "$tmpdir" FETCH_HEAD
  (
    cd "$tmpdir"
    git switch -C "$branch"
  )
else
  git worktree add --detach "$tmpdir"
  (
    cd "$tmpdir"
    git switch --orphan "$branch"
    git rm -rf . >/dev/null 2>&1 || true
  )
fi

find "$tmpdir" -mindepth 1 -maxdepth 1 ! -name .git -exec rm -rf {} +
cp -R public/. "$tmpdir"/
touch "$tmpdir/.nojekyll"

(
  cd "$tmpdir"
  git config user.name >/dev/null 2>&1 || git config user.name "atlas-ru deploy"
  git config user.email >/dev/null 2>&1 || git config user.email "atlas-ru-deploy@example.invalid"
  git add -A
  if git diff --cached --quiet; then
    echo "No GitHub Pages changes to deploy."
    exit 0
  fi
  git commit -m "$message"
  git push "$remote" "HEAD:${branch}"
)
