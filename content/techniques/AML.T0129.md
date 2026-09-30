---
atlas_id: AML.T0129
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may place instructions or triggers in one part of a multimodal input to influence the model while staying unnoticed by human reviewers and by defenses that do not inspect all input modalities. Multimodal...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 0
modified_date: "2026-09-15"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Triggers in Multimodal Inputs
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Triggers in Multimodal Inputs
url: /techniques/AML.T0129/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may place instructions or triggers in one part of a multimodal input to influence the model while staying unnoticed by human reviewers and by defenses that do not inspect all input modalities.

Multimodal systems jointly process text alongside other modalities, yet many moderation and filtering controls operate mainly on the textual channel. A payload placed in a modality that is not inspected, or is inspected differently from text, is still parsed by the model but can escape detection, allowing it to alter model output or carry a cross-modal prompt injection.

Examples include instructions placed in an image, audio, or video channel, or carried in file metadata (e.g. EXIF for images, ID3 tags for audio, or document metadata).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>
