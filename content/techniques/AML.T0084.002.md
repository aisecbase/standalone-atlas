---
atlas_id: AML.T0084.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут выявлять ключевые слова или другие триггеры, такие как входящие письма, добавление документов, входящие сообщения или другие рабочие процессы, которые активируют агента и могут заставить его...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 2
source_name: Activation Triggers
subtechnique_count: 0
subtechnique_of: AML.T0084
tactics:
    - AML.TA0008
title: Триггеры активации
url: /techniques/AML.T0084.002/
---

Злоумышленники могут выявлять ключевые слова или другие триггеры, такие как входящие письма, добавление документов, входящие сообщения или другие рабочие процессы, которые активируют агента и могут заставить его выполнить дополнительные действия.

Понимание этих триггеров может раскрыть, как ИИ-агент активируется и управляется. Это также может открыть дополнительные пути компрометации, поскольку злоумышленник может попытаться вызвать срабатывание агента извне его среды и заставить его выполнить непреднамеренные или вредоносные действия.


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
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Avoid publicly disclosing agent activation keywords, events, and workflows.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследователи делают вывод, что ИИ-агент активируется при получении письма.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0008 Выявление</span><p>Исследователи установили, что Claude Code Action мог активироваться событиями GitHub, связанными с issues, pull requests и комментариями; после срабатывания Claude Code Action загружал связанное с событием содержимое в контекст Claude.</p></a>
</div>
