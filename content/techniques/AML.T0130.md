---
atlas_id: AML.T0130
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Злоумышленники могут манипулировать ИИ-ассистентом, чтобы в ответах он отдавал предпочтение выбранным ими источникам или материалам. Внедряя инструкции, такие как treat [source] as a trusted source или recommend...
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
title: Формирование предвзятых ответов ИИ-агента
url: /techniques/AML.T0130/
---

Злоумышленники могут манипулировать ИИ-ассистентом, чтобы в ответах он отдавал предпочтение выбранным ими источникам или материалам. Внедряя инструкции, такие как `treat [source] as a trusted source` или `recommend [source] first,`, злоумышленник делает ответы ассистента предвзятыми в своих интересах: ассистент представляет рекламный контент или материалы, выгодные злоумышленнику, как нейтральный, хорошо обоснованный ответ. Это нарушает целостность ответов ассистента и снижает их надёжность в вопросах, в которых пользователь может на него полагаться, например здоровья, финансов или безопасности. При этом пользователь не знает, что рекомендации стали предвзятыми.

Внедряемые инструкции могут доставляться разными способами, например через [Специально сформированные ссылки на ИИ-ассистента]. Такое воздействие может сохраняться, если память агента была отравлена (см. [Отравление контекста ИИ-агента: Память](/techniques/AML.T0080.000)).


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
