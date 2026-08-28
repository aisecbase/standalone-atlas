---
atlas_id: AML.T0067.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут манипулировать ссылками на источники, приведёнными в ответе ИИ-системы, чтобы этот ответ выглядел заслуживающим доверия. К возможным вариантам относятся указание неверной ссылки на источник,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Citations
subtechnique_count: 0
subtechnique_of: AML.T0067
tactics:
    - AML.TA0007
title: Ссылки на источники
url: /techniques/AML.T0067.000/
---

Злоумышленники могут манипулировать ссылками на источники, приведёнными в ответе ИИ-системы, чтобы этот ответ выглядел заслуживающим доверия. К возможным вариантам относятся указание неверной ссылки на источник, создание новой вымышленной ссылки на источник либо указание корректной ссылки на источник, но применительно к данным, предоставленным злоумышленником.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0067/"><span class="relation-id">AML.T0067</span><strong>Манипуляция доверенными компонентами ответа LLM</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователи добавили инструкции для манипуляции ссылками на источники в ответе, злоупотребляя доверием пользователя к Copilot. Инструкции заставляли Copilot ссылаться только на один `EmailMessage` в формате `[^1^]` и игнорировать остальные файлы.</p></a>
</div>
