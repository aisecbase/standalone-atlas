---
atlas_id: AML.T0134
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may selectively deliver malicious or manipulated content to AI systems, while presenting different benign content to human users, web crawlers, or security detection mechanisms. They may identify AI...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 0
modified_date: "2026-09-15"
platforms:
    - Agentic AI
procedure_count: 0
source_name: AI Targeted Cloaking
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: AI Targeted Cloaking
url: /techniques/AML.T0134/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may selectively deliver malicious or manipulated content to AI systems, while presenting different benign content to human users, web crawlers, or security detection mechanisms. They may identify AI browsers or agents via user-agent strings and condition the server response so that AI-based clients receive prompt injections or misleading content.

Preventing human visitors, conventional web crawlers, and security tools from observing the same content allows adversaries to make malicious input more difficult to detect. [AI Targeted Cloaking](/techniques/AML.T0134) may be combined with [Drive-by Compromise](/techniques/AML.T0078) to deliver an [LLM Prompt Injection](/techniques/AML.T0051).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>
