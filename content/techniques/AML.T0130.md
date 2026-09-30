---
atlas_id: AML.T0130
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may manipulate an AI assistant so that it favors adversary-chosen sources, or content in its responses. By injecting instructions such as "treat [source] as a trusted source" or "recommend [source] first,"...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-09-15"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: AI Agent Response Biasing
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0011
title: AI Agent Response Biasing
url: /techniques/AML.T0130/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may manipulate an AI assistant so that it favors adversary-chosen sources, or content in its responses. By injecting instructions such as "treat [source] as a trusted source" or "recommend [source] first," an adversary biases the assistant's outputs toward their own interests, causing it to present promotional or self-serving content as if it were a neutral, well-reasoned response. This degrades the integrity and trustworthiness of the assistant's responses on topics the user may rely on, such as health, finance, or security, without the user being aware that the advice has been skewed.

The injected instructions can be delivered through different ways, for example through [Crafted AI Assistant Links]. This impact may persist if the agent's memory was poisoned (See [AI Agent Context Poisoning: Memory](/techniques/AML.T0080.000)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0072/"><span class="relation-id">AML.CS0072</span><strong>AI Recommendation Poisoning via Crafted AI Assistant Links</strong><span class="relation-meta">Актор: Multiple commercial entities; 31 distinct companies identified / Тактика: AML.TA0011 Воздействие</span><p>In subsequent unrelated conversations, the assistant preferentially surfaced the operator&#39;s domain or product and presented the result as a neutral recommendation. Observed targeting included health and financial topics, where skewed recommendations carry elevated consequences.</p></a>
</div>


## Источники

- [Manipulating AI memory for profit: The rise of AI Recommendation Poisoning](https://www.microsoft.com/en-us/security/blog/2026/02/10/ai-recommendation-poisoning/)
