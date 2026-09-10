# AGENTS.md

Guidance for coding agents working on this repository.

## Project

This is an unofficial static reference site for MITRE ATLAS, currently focused on the Russian localization. It builds a Hugo site from the official `mitre-atlas/atlas-data` release data plus local translation overlays.

Keep the site clearly unofficial. Do not remove the disclaimer, official ATLAS link, source data link, or licensing/notice references.

## Important Paths

- `data/atlas/ATLAS.yaml` - upstream ATLAS data, normally downloaded from the latest `mitre-atlas/atlas-data` release. Since `v2026.05`, this is the ATLAS YAML v6 release asset (`ATLAS-YYYY.MM.yaml`), not legacy `dist/ATLAS.yaml`.
- `data/generated/atlas.yaml` - generated adapter output for Hugo templates. Do not edit by hand.
- `data/generated/*.yaml`, `reports/translation-coverage.md`, and `content/translation-coverage.md` are generated snapshots committed for review. Regenerate them with `cmd/atlasgen`; do not edit them by hand.
- `data/translations/ru.yaml` - Russian translation overlay by ATLAS object ID.
- `data/translations/<lang>.yaml` - optional translation overlays for other languages.
- `cmd/atlasgen/` - generator that creates Hugo content pages and translation coverage data.
- `content/` - generated and hand-written Hugo content.
- `layouts/` - Hugo templates for lists, object pages, matrix, resources, and knowledge graph.
- `assets/css/site.css` - main stylesheet.
- `reports/translation-coverage.md` - generated translation coverage report.

## Common Commands

Build from the latest ATLAS data release:

```bash
make build
```

Regenerate content from local data without downloading:

```bash
UPDATE_ATLAS_DATA=0 make build
```

Run only the generator:

```bash
make generate
```

Run the generator with a non-default translation overlay:

```bash
ATLAS_LANG=es go run ./cmd/atlasgen
ATLAS_TRANSLATION_FILE=data/translations/es-MX.yaml go run ./cmd/atlasgen
```

Run Hugo locally:

```bash
hugo server -D
```

Production build:

```bash
hugo --minify
```

Check the latest upstream release metadata:

```bash
curl -fsSL https://api.github.com/repos/mitre-atlas/atlas-data/releases/latest
```

## Workflow Rules

- Commit every logical change. Keep commits small and named in English.
- Before committing, run at least `hugo --minify` for template/content/CSS changes.
- If generator code changes, run `gofmt -w cmd/atlasgen/*.go`, `go run ./cmd/atlasgen`, and `go test ./cmd/atlasgen`.
- When running Go tests in this Codex workspace, use an escalated/out-of-sandbox command for `go test`. The sandbox can make `go test ./cmd/atlasgen` fail with `package testing is not in std (/usr/local/go/src/testing)`, while the same command succeeds outside the sandbox.
- For UI verification, prefer Firefox headless with a separate temporary profile. Chrome headless has been unreliable in this workspace and can hang after writing screenshots.
- Do not edit generated object pages in `content/tactics`, `content/techniques`, `content/mitigations`, or `content/studies` by hand unless the task is explicitly about generated output. Prefer changing `cmd/atlasgen` or `data/translations/ru.yaml`.
- Do not edit `data/generated/atlas.yaml` by hand. It is written by `cmd/atlasgen` and consumed by Hugo templates.
- Do not commit `public/`, `resources/_gen/`, `.hugo_build.lock`, or `.DS_Store`.
- Preserve aliases when moving public pages, especially `/knowledge-graph/` and `/about/`.
- GitHub Pages builds from committed `data/atlas/ATLAS.yaml` on push. The scheduled weekly workflow intentionally downloads latest upstream ATLAS data so the public coverage page can expose newly untranslated or stale translations. Manual workflow runs download upstream data only when `atlas_release_tag` is provided.

## ATLAS Data Workflow

- `scripts/update-atlas-data.sh` should prefer the release asset named like `ATLAS-YYYY.MM.yaml`. The legacy `dist/ATLAS.yaml` path is only a fallback for older releases.
- Since `v2026.05`, upstream ATLAS separates content version (`collection.version`, for example `2026.05`) from data format version (`format-version`, for example `6.0.0`). Do not compare only the old top-level `version`.
- After updating upstream data, run `go run ./cmd/atlasgen` before Hugo. The generator writes object pages, translation reports, and `data/generated/atlas.yaml`.
- Do not run Hugo in parallel with `go run ./cmd/atlasgen` or `make generate`. The generator removes and recreates generated object pages, so Hugo can panic on transient missing files if it starts during generation. If this happens, wait for the generator to finish and rerun `hugo --minify`.
- The generator supports legacy ATLAS YAML and v6. For v6 it adapts keyed maps and first-class relationships into the local Hugo model.
- The v6 adapter must preserve:
  - matrix tactic order from `relationships[ATLAS-matrix].sequences`;
  - technique tactics from `achieves`;
  - subtechnique parent links from `specializes`;
  - mitigation technique uses from `mitigates`;
  - case-study procedures from `employs`;
  - technique `platforms`.
- When changing the adapter, add or update Go tests that cover the relationship mapping. Do not rely only on Hugo rendering to catch data loss.

## Content Rules

- Keep translations separate from upstream data.
- For a new localization, add a separate overlay in `data/translations/<lang>.yaml`; do not mix languages in `data/translations/ru.yaml`.
- `EN` in the language switcher currently points to the official ATLAS site, not to a local English build. Mark external destinations clearly.
- Do not imply affiliation with MITRE.
- For contributions to ATLAS itself, point users to `https://atlas.mitre.org/resources/contribute`.
- Official resource pages are mirrored in `content/resources/` where requested. Preserve source links and image paths, and avoid changing technical meaning while translating or adapting them.
- If an official resource page moves in the current MITRE ATLAS SPA, update the mirrored page title, `url`, aliases, and `source_url` to match the current route while preserving old local aliases.

## Translation Workflow

- Before starting a translation session, check the latest `mitre-atlas/atlas-data` release tag and ATLAS content version against local `data/atlas/ATLAS.yaml`. For ATLAS v6, use `collection.version` and `format-version`; do not rely on the legacy top-level `version`. If a newer release exists, tell the user and prefer updating the data first, then regenerate content and review any translations flagged by source hashes.
- For English to Russian ATLAS review sessions, use only two translation agents: `ru_translator` for the primary translation/audit pass and `ru_editor` for the final review/verdict. The main agent is responsible for structural validation, including Markdown, YAML, generated Hugo content, and layout/build checks.
- Before applying final Russian translation edits, present the proposed fixes to the user and wait for explicit approval.
- Translate ATLAS objects one at a time and send each proposed translation to the user for review before editing files.
- For each review item, include the object ID, original English title, original English description, proposed Russian title, and proposed Russian description.
- Always present translation proposals in a Markdown table with the original English text in the left `Original` column and the Russian translation in the right `Предлагаемый перевод` column. This applies to initial proposals, individual phrase corrections, and final editor-reviewed versions. Repeat the corresponding original alongside each revised proposal even if it appeared earlier in the conversation; never present a Russian-only proposal for approval. Include this table in the final user-facing response, not only in a progress update.
- If a prior Russian ATLAS translation is available as a reference, show only short relevant terminology or excerpts for comparison. Treat it as a terminology aid, not as the source of truth.
- Do not execute code downloaded from reference sites such as `lms20.ru` or old `atlas.securityhub.ru` artifacts. Only use passive text inspection, for example `curl`/`rg` over downloaded text, when checking old terminology.
- Use cybersecurity-oriented Russian terminology. Prefer clear threat-intelligence/AppSec language over academic machine-learning phrasing.
- Use `злоумышленник` for `adversary`, `организация-жертва` or `целевая организация` for `victim organization` depending on context, and `состязательные атаки на ИИ` for `Adversarial AI Attacks`.
- After the user approves a translation, update `data/translations/ru.yaml`, include `name_sha256` plus the relevant body source hash (`description_sha256` or `summary_sha256`), run `go run ./cmd/atlasgen`, run `hugo --minify`, and commit the logical change.
- Do not hand-edit generated object pages for translation changes; let `cmd/atlasgen` update `content/`, reports, and generated data.
- Mitigation relationship `use` descriptions render as relation-card text on both mitigation and technique pages. Translate repeated source strings once in top-level `technique_uses` in `data/translations/ru.yaml`; use `objects[AML.Mxxxx].techniques` only for mitigation-specific overrides.

## Translation Review Rules

- Coverage reports and source hashes only prove freshness against the English source. They do not prove editorial quality; still compare Russian text against the English source for omissions, semantic drift, untranslated labels, broken code spans, broken links, and numeric changes.
- When reviewing `data/translations/ru.yaml`, compare against `data/generated/atlas.yaml` or `data/atlas/ATLAS.yaml`. For case-study procedures, check the exact procedure step by index and by `tactic`/`technique`.
- Review generated relation-card paragraphs in `content/techniques/*.md` and `content/mitigations/*.md`; they come from mitigation `use` relationships and can leak English even when object titles and descriptions are translated.
- Do not summarize away operational details from ATLAS case studies. Preserve embedded prompts, commands, URLs, code blocks, markers, delimiters, payload strings, file names, and concrete examples. Translate the surrounding explanatory prose, but keep literal payload/code text intact when it is the evidence or mechanism of the attack.
- If the English source uses inline HTML only to show prompt/code snippets, it is acceptable to represent the same content in Russian Markdown with code fences, as long as the technical content is preserved.
- Translate linked ATLAS technique/tactic labels in Russian prose when a Russian label already exists. Avoid leftover English labels such as `LLM prompt injection`, `Publish Hallucinated Entities`, `AI Supply Chain Rug Pull`, `Spearphishing Link`, `Exploitation for Client Execution`, or `Malicious File` unless they are explicitly source names, product names, commands, reference titles, or intentionally quoted source text.
- Use `Выявление` for the ATLAS tactic `Discovery` (`AML.TA0008`). Keep `обнаружение` for defensive detection contexts, for example malware detection or deepfake detection.
- Preserve Markdown links, code spans, reference markers, object IDs, percentages, dates, version strings, package names, commands, and API names exactly unless there is a clear reason to localize the surrounding words.
- After a review/fix pass, run targeted searches for leftover English labels, untranslated relation-card paragraphs, and terminology drift, then run `go run ./cmd/atlasgen` and `hugo --minify`. If generated content changes, commit the overlay and generated content together.

## UI Rules

- This is a reference/work tool, not a marketing site. Favor dense, calm, readable UI.
- Check desktop and mobile states for matrix, tables, object pages, resources, and knowledge graph when touching CSS/layouts.
- Avoid hidden horizontal overflow on mobile unless it is an intentional table/matrix scroll surface with clear affordance.
- Keep table IDs and dates non-wrapping.
- The matrix has filters for subtechniques, minimum maturity, and ATLAS v6 platforms. Keep counts, tactic column counts, parent/subtechnique visibility, and mobile wrapping working together.
- Technique pages should show ATLAS v6 `platforms` in the sidebar when present.
- For the knowledge graph, keep node click, hover/focus state, inspector state, and mobile layout working.
- On mobile, the knowledge graph canvas intentionally scrolls horizontally so nodes stay large enough to inspect.
