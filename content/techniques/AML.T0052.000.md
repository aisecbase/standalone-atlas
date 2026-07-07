---
atlas_id: AML.T0052.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленники могут превращать LLM в инструмент целевой социальной инженерии. LLM способны взаимодействовать с пользователями в текстовых диалогах. Злоумышленник может дать им инструкции выманивать у пользователя...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Spearphishing via Social Engineering LLM
subtechnique_count: 0
subtechnique_of: AML.T0052
tactics:
    - AML.TA0004
    - AML.TA0015
title: Целевой фишинг через LLM для социальной инженерии
url: /techniques/AML.T0052.000/
---

Злоумышленники могут превращать LLM в инструмент целевой социальной инженерии.

LLM способны взаимодействовать с пользователями в текстовых диалогах.

Злоумышленник может дать им инструкции выманивать у пользователя чувствительную информацию и действовать как эффективный инструмент социальной инженерии.

Они могут быть нацелены на конкретные персоны, заданные злоумышленником.

Это позволяет злоумышленникам масштабировать целевой фишинг и добиваться от отдельных людей раскрытия приватной информации, например учетных данных для привилегированных систем.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0052/"><span class="relation-id">AML.T0052</span><strong>Фишинг</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0018/"><span class="relation-id">AML.M0018</span><strong>Обучение пользователей</strong><p>Обучайте пользователей распознавать попытки фишинга и понимать, что ИИ может использоваться для генерации адресных и убедительных сообщений.</p></a>
<a class="relation-item" href="/mitigations/AML.M0034/"><span class="relation-id">AML.M0034</span><strong>Обнаружение дипфейков</strong><p>Обнаружение дипфейков можно использовать для выявления и блокировки фишинговых попыток, использующих сгенерированный контент.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0004 Первичный доступ</span><p>Вредоносная инструкция заставляет Bing Chat перейти на «пиратскую» манеру общения и незаметно убеждать пользователя раскрыть персональные данные, например имя. Затем она побуждает пользователя перейти по ссылке, в URL которой эти данные уже закодированы.</p></a>
</div>
