---
atlas_id: AML.T0024.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Обучающие данные ИИ-модели могут быть реконструированы за счет использования оценок уверенности модели, доступных через API инференса. Стратегически отправляя запросы к API инференса, злоумышленники могут восстановить...
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
source_name: Invert AI Model
subtechnique_count: 0
subtechnique_of: AML.T0024
tactics:
    - AML.TA0010
title: Инверсия ИИ-модели
url: /techniques/AML.T0024.001/
---

Обучающие данные ИИ-модели могут быть реконструированы за счет использования оценок уверенности модели, доступных через API инференса.

Стратегически отправляя запросы к API инференса, злоумышленники могут восстановить потенциально приватную информацию, встроенную в обучающие данные.

Это может привести к нарушению конфиденциальности, если злоумышленник сможет реконструировать данные чувствительных признаков, используемых в алгоритме.


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
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Рекомендуемые подходы:</p><ul><li>ограничить количество показываемых результатов</li><li>ограничить детализацию онтологии выходных классов</li><li>использовать методы рандомизированного сглаживания</li><li>снизить точность числовых выходных данных</li></ul></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Попытайтесь реконструировать чувствительные записи, атрибуты или сведения, характерные для обучающих данных, на основе выходных данных модели. Устраните утечку за счёт внесения изменений в модель, минимизации выходных данных, мер защиты приватности и ограничения доступа к инференсу.</p></a>
</div>
