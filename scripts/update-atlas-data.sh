#!/usr/bin/env bash
set -euo pipefail

repo="${ATLAS_DATA_REPO:-mitre-atlas/atlas-data}"
tag="${ATLAS_RELEASE_TAG:-}"
dest="${ATLAS_DATA_DEST:-data/atlas/ATLAS.yaml}"
api_base="${GITHUB_API_URL:-https://api.github.com}"
token="${GH_TOKEN:-${GITHUB_TOKEN:-}}"

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT

headers=(
  -H "Accept: application/vnd.github+json"
  -H "X-GitHub-Api-Version: 2022-11-28"
)
if [[ -n "$token" ]]; then
  headers+=(-H "Authorization: Bearer $token")
fi

if [[ -n "$tag" ]]; then
  release_url="${api_base}/repos/${repo}/releases/tags/${tag}"
else
  release_url="${api_base}/repos/${repo}/releases/latest"
fi

release_json="${tmpdir}/release.json"
curl -fsSL "${headers[@]}" "$release_url" -o "$release_json"

if ! command -v jq >/dev/null 2>&1; then
  echo "Warning: jq is not installed; falling back to brittle sed parsing of GitHub release JSON." >&2
fi

json_field() {
  local field="$1"
  if command -v jq >/dev/null 2>&1; then
    jq -r --arg field "$field" '.[$field] // empty' "$release_json"
  else
    sed -nE "s/.*\"${field}\"[[:space:]]*:[[:space:]]*\"([^\"]*)\".*/\1/p" "$release_json" | head -n 1
  fi
}

tag_name="$(json_field tag_name)"
tarball_url="$(json_field tarball_url)"
asset_url=""
if command -v jq >/dev/null 2>&1; then
  asset_url="$(jq -r '.assets[]? | select(.name | test("^ATLAS-[0-9]{4}\\.[0-9]{2}(\\.[0-9]+)?\\.yaml$")) | .browser_download_url' "$release_json" | head -n 1)"
else
  asset_url="$(sed -nE 's/.*"browser_download_url"[[:space:]]*:[[:space:]]*"([^"]*\/ATLAS-[0-9]{4}\.[0-9]{2}(\.[0-9]+)?\.yaml)".*/\1/p' "$release_json" | head -n 1)"
fi

if [[ -n "$asset_url" ]]; then
  mkdir -p "$(dirname "$dest")"
  curl -fsSL -L "${headers[@]}" "$asset_url" -o "${tmpdir}/ATLAS.yaml"
  mv "${tmpdir}/ATLAS.yaml" "$dest"
  echo "Downloaded ${repo} ${tag_name:-latest} v6 asset -> ${dest}"
  exit 0
fi

if [[ -z "$tarball_url" ]]; then
  echo "Could not find tarball_url in GitHub release response: $release_url" >&2
  exit 1
fi

archive="${tmpdir}/atlas-data.tar.gz"
curl -fsSL -L "${headers[@]}" "$tarball_url" -o "$archive"

atlas_path="$(tar -tzf "$archive" | grep '/dist/ATLAS.yaml$' | head -n 1 || true)"
if [[ -z "$atlas_path" ]]; then
  echo "Release ${tag_name:-unknown} does not contain dist/ATLAS.yaml" >&2
  exit 1
fi

mkdir -p "$(dirname "$dest")"
tar -xOf "$archive" "$atlas_path" > "${tmpdir}/ATLAS.yaml"
mv "${tmpdir}/ATLAS.yaml" "$dest"

echo "Downloaded ${repo} ${tag_name:-latest} -> ${dest}"
