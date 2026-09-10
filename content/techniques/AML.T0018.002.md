---
atlas_id: AML.T0018.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-04-09"
description: Злоумышленники могут встраивать вредоносный код в файлы ИИ-моделей. ИИ-модели могут распространяться как сочетание инструкций и весов. Некоторые форматы, например pickle-файлы, небезопасно десериализовать, поскольку...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Embed Malware
subtechnique_count: 0
subtechnique_of: AML.T0018
tactics:
    - AML.TA0001
    - AML.TA0006
title: Встраивание вредоносного ПО
url: /techniques/AML.T0018.002/
---

Злоумышленники могут встраивать вредоносный код в файлы ИИ-моделей.

ИИ-модели могут распространяться как сочетание инструкций и весов.

Некоторые форматы, например pickle-файлы, небезопасно десериализовать, поскольку они могут содержать опасные вызовы, такие как `exec`.

Модели со встроенным вредоносным ПО могут при этом работать ожидаемым образом.

Это может позволить злоумышленникам добиться выполнения кода, создания канала командного управления (C2) или эксфильтрации данных.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Адаптация атак, связанных с ИИ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Исследователь встроил [Sliver](https://github.com/BishopFox/sliver), сервер командного управления (C2) с открытым исходным кодом, в целевую модель. Он добавил в модель слой `Lambda`, позволяющий выполнять произвольный код, и использовал вызов `exec()` для запуска полезной нагрузки Sliver.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Злоумышленник встроил вредоносное ПО в ИИ-модель, хранившуюся в pickle-файле. Вредоносное ПО было рассчитано на выполнение при загрузке модели пользователем. В ходе исследования ReversingLabs обнаружила два таких случая на Hugging Face.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Специалисты Unit 42 подготовили подконтрольные злоумышленнику артефакты моделей с полезной нагрузкой, запускавшей реверс-шелл при развёртывании или загрузке артефакта.</p></a>
</div>
