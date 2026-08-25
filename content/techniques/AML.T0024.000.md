---
atlas_id: AML.T0024.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут определить, входил ли конкретный образец данных в обучающую выборку модели, или выявить общие характеристики данных в этой выборке. Это создает риски для конфиденциальности. Некоторые подходы...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Infer Training Data Membership
subtechnique_count: 0
subtechnique_of: AML.T0024
tactics:
    - AML.TA0010
title: Определение принадлежности к обучающей выборке
url: /techniques/AML.T0024.000/
---

Злоумышленники могут определить, входил ли конкретный образец данных в обучающую выборку модели, или выявить общие характеристики данных в этой выборке. Это создает риски для конфиденциальности.

Некоторые подходы используют теневую модель, которую можно получить через [обучение прокси-модели через репликацию](/techniques/AML.T0005.001); другие опираются на статистику оценок предсказаний модели.

В результате модель организации-жертвы может раскрыть приватную информацию, например персональные данные людей из обучающей выборки или другие виды защищенной интеллектуальной собственности.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0024/"><span class="relation-id">AML.T0024</span><strong>Эксфильтрация через API инференса ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Пассивная обфускация выходных данных ИИ</strong><p>Рекомендуемые подходы:</p><ul><li>ограничить количество показываемых результатов</li><li>ограничить детализацию онтологии выходных классов</li><li>использовать методы рандомизированного сглаживания</li><li>снизить точность числовых выходных данных</li></ul></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение количества запросов к ИИ-модели</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Test whether model outputs reveal the membership of known training samples. Improve privacy-preserving training, output restriction, access controls, and query monitoring.</p></a>
</div>
