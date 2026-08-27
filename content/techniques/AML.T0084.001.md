---
atlas_id: AML.T0084.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут выявлять инструменты, к которым имеет доступ ИИ-агент. Определив доступные инструменты, злоумышленник может понять, какие действия можно выполнить через агента и к каким дополнительным ресурсам он...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Tool Definitions
subtechnique_count: 0
subtechnique_of: AML.T0084
tactics:
    - AML.TA0008
title: Определения инструментов
url: /techniques/AML.T0084.001/
---

Злоумышленники могут выявлять инструменты, к которым имеет доступ ИИ-агент. Определив доступные инструменты, злоумышленник может понять, какие действия можно выполнить через агента и к каким дополнительным ресурсам он может получить доступ. Эти сведения могут раскрыть доступ к внешним источникам данных, таким как OneDrive или SharePoint, либо выявить пути эксфильтрации, например возможность отправлять электронные письма. Это помогает злоумышленникам определить ИИ-агентов, которые представляют наибольшую ценность или открывают лучшие возможности для атаки.


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
<a class="relation-item" href="/techniques/AML.T0084.000/"><span class="relation-id">AML.T0084.000</span><strong>Встроенные знания</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Avoid publicly documenting sensitive agent tool definitions and capabilities.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследователи делают вывод, что у ИИ-агента есть инструмент для отправки писем.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаруживают, что ИИ-агент имеет доступ к инструменту Salesforce `get-records`, который можно использовать для получения записей CRM.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0008 Выявление</span><p>Исследователи выявили инструменты, доступные Claude Code Action, и сравнили пути их выполнения. Они установили, что подпроцессы Bash могли выполняться внутри Bubblewrap с очищенным окружением, тогда как встроенный инструмент Read осуществлял прямой внутрипроцессный доступ к файлам за пределами этой границы изоляции.</p></a>
</div>
