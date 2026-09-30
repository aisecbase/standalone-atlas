---
atlas_id: AML.T0131
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may craft links that open an AI assistant or agent with attacker-controlled input already supplied, so that opening the link initiates an interaction the adversary defines rather than one the target...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-09-15"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Crafted AI Assistant Links
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Crafted AI Assistant Links
url: /techniques/AML.T0131/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may craft links that open an AI assistant or agent with attacker-controlled input already supplied, so that opening the link initiates an interaction the adversary defines rather than one the target composed. Many AI assistants accept a prompt through URL parameters (for example `?q=` or `?prompt=`) that is automatically populated, and in some cases submitted, when the link is opened. By encoding a chosen prompt into such a link, an adversary can cause the target's assistant to act on supplied instructions as soon as the link is opened.

These links are frequently disguised as helpful actions, such as a "Summarize with AI" button or a share link, and distributed through web pages, emails, documents, or messages. Because the resulting interaction runs in the target's own assistant session, a crafted link can drive a range of downstream impacts depending on the supplied instructions, such as exfiltrating data the assistant can access or for [AI Recommendation Poisoning](/techniques/AML.T0131).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0072/"><span class="relation-id">AML.CS0072</span><strong>AI Recommendation Poisoning via Crafted AI Assistant Links</strong><span class="relation-meta">Актор: Multiple commercial entities; 31 distinct companies identified / Тактика: AML.TA0004 Первичный доступ</span><p>The user clicked the button or link, which opened the AI assistant domain with the operator&#39;s prompt pre-populated in the input field via a `?q=` or `?prompt=` parameter.</p></a>
</div>
