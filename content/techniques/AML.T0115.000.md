---
aliases:
    - /techniques/AML.T0019/
atlas_id: AML.T0115.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут публиковать отравленные наборы данных, предназначенные для обучения или дообучения ИИ-моделей. Набор данных может быть вновь созданным либо представлять собой модифицированный вариант легитимного...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Datasets
subtechnique_count: 0
subtechnique_of: AML.T0115
tactics:
    - AML.TA0003
title: Наборы данных
url: /techniques/AML.T0115.000/
---

Злоумышленники могут публиковать отравленные наборы данных, предназначенные для обучения или дообучения ИИ-моделей. Набор данных может быть вновь созданным либо представлять собой модифицированный вариант легитимного набора данных и может содержать изменённые образцы, метки, аннотации или метаданные.

Злоумышленники могут распространять отравленные наборы данных через репозитории наборов данных, репозитории кода, файлообменные сервисы или скомпрометированные источники данных. Жертва, которая включает такой набор данных в пайплайн обучения, может пострадать от [отравления обучающих данных](/techniques/AML.T0020), возможно, в результате [компрометации цепочки поставок ИИ](/techniques/AML.T0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Публикация отравленных ИИ-артефактов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.002/"><span class="relation-id">AML.T0115.002</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Репозитории наборов данных проверяют поступающие материалы и до публикации в каталоге помещают в карантин отравленные образцы данных, метки, аннотации или метаданные.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Затем злоумышленник может разместить отравленные данные на подконтрольных ему доменах. В рамках этого испытания исследователи отслеживали запросы к подконтрольным им URL и тем самым фиксировали скачивания, чтобы продемонстрировать наличие активных пользователей датасета.</p></a>
</div>
