---
atlas_id: AML.T0131
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Злоумышленники могут создавать ссылки, которые открывают ИИ-ассистента или агента с уже подставленными входными данными, контролируемыми злоумышленником. При открытии такой ссылки начинается взаимодействие, заданное...
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
title: Специально сформированные ссылки на ИИ-ассистента
url: /techniques/AML.T0131/
---

Злоумышленники могут создавать ссылки, которые открывают ИИ-ассистента или агента с уже подставленными входными данными, контролируемыми злоумышленником. При открытии такой ссылки начинается взаимодействие, заданное злоумышленником, а не сформулированное самой жертвой. Многие ИИ-ассистенты принимают промпт через параметры URL (например, `?q=` или `?prompt=`): при открытии ссылки этот промпт автоматически подставляется, а в некоторых случаях и отправляется. Закодировав выбранный промпт в такой ссылке, злоумышленник может заставить ассистента жертвы выполнять заданные инструкции сразу после открытия ссылки.

Такие ссылки часто маскируют под полезные действия, например кнопку «Summarize with AI» или ссылку для обмена материалами, и распространяют через веб-страницы, электронные письма, документы или сообщения. Поскольку вызванное ссылкой взаимодействие происходит в сессии ИИ-ассистента самой жертвы, специально сформированная ссылка может приводить к различным последствиям в зависимости от заданных инструкций, например к эксфильтрации данных, доступных ассистенту, или к [формированию предвзятых ответов ИИ-агента](/techniques/AML.T0130).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0072/"><span class="relation-id">AML.CS0072</span><strong>AI Recommendation Poisoning via Crafted AI Assistant Links</strong><span class="relation-meta">Актор: Multiple commercial entities; 31 distinct companies identified / Тактика: AML.TA0004 Первичный доступ</span><p>The user clicked the button or link, which opened the AI assistant domain with the operator&#39;s prompt pre-populated in the input field via a `?q=` or `?prompt=` parameter.</p></a>
</div>
