---
atlas_id: AML.T0018
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут манипулировать артефактом ИИ-модели или включёнными в его состав компонентами, чтобы изменить поведение ИИ-системы, внедрить вредоносный код или создать долговременно сохраняющуюся вредоносную...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 4
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Manipulate AI Model
subtechnique_count: 4
subtechnique_of: ""
tactics:
    - AML.TA0001
    - AML.TA0006
title: Манипуляция ИИ-моделью
url: /techniques/AML.T0018/
---

Злоумышленники могут манипулировать артефактом ИИ-модели или включёнными в его состав компонентами, чтобы изменить поведение ИИ-системы, внедрить вредоносный код или создать долговременно сохраняющуюся вредоносную функциональность. Это может включать изменение весов модели, её архитектуры или логики формирования промпта.

Артефакты, подвергшиеся манипуляции, могут сохранять ожидаемое поведение в обычных условиях, активируя вредоносное поведение только при обработке определённых входных данных, в определённых контекстах или при определённых условиях развёртывания.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать подмену ИИ-артефактов и их неавторизованное изменение.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Проверка ИИ-модели на широком наборе состязательных входных данных может повысить уверенность в том, что модель не подвергалась манипуляциям.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>В контролируемых условиях попытайтесь изменить или подменить модели, веса, адаптеры и связанную с ними конфигурацию. Устраните недостатки механизмов авторизации, обеспечения целостности артефактов, процедур одобрения развёртывания, мониторинга и восстановления.</p></a>
</div>
