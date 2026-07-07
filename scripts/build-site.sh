#!/usr/bin/env bash
set -euo pipefail

update_data="${UPDATE_ATLAS_DATA:-1}"
require_release_tag="${REQUIRE_ATLAS_RELEASE_TAG:-0}"
hugo_args=(--minify)

if [[ "$update_data" != "0" ]]; then
  if [[ "$require_release_tag" == "1" && -z "${ATLAS_RELEASE_TAG:-}" ]]; then
    echo "ATLAS_RELEASE_TAG is required when REQUIRE_ATLAS_RELEASE_TAG=1." >&2
    exit 1
  fi
  scripts/update-atlas-data.sh
fi

go run ./cmd/atlasgen

rm -rf public

if [[ -n "${HUGO_BASEURL:-}" ]]; then
  hugo_args+=(--baseURL "$HUGO_BASEURL")
fi

hugo "${hugo_args[@]}" "$@"
