---
atlas_id: AML.T0084.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут пытаться выявить источники данных, к которым имеет доступ конкретный агент. Конфигурация ИИ-агента может раскрывать источники данных или знания. Встроенные знания могут включать чувствительные или...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 1
source_name: Embedded Knowledge
subtechnique_count: 0
subtechnique_of: AML.T0084
tactics:
    - AML.TA0008
title: Встроенные знания
url: /techniques/AML.T0084.000/
---

Злоумышленники могут пытаться выявить источники данных, к которым имеет доступ конкретный агент. Конфигурация ИИ-агента может раскрывать источники данных или знания.

Встроенные знания могут включать чувствительные или проприетарные материалы, такие как интеллектуальная собственность, данные клиентов, внутренние политики или даже учетные данные. Определяя, к каким знаниям имеет доступ агент, злоумышленник может лучше понять роль ИИ-агента и потенциально раскрыть конфиденциальную информацию или определить ценные цели для дальнейшей эксплуатации.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0084/"><span class="relation-id">AML.T0084</span><strong>Выявление конфигурации ИИ-агента</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаруживают, что ИИ-агент имеет доступ к источнику данных «Customer Support Account Owners.csv».</p></a>
</div>
