---
atlas_id: AML.T0024
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут эксфильтровать приватную информацию через доступ к API инференса ИИ-модели. Известно, что ИИ-модели могут раскрывать приватную информацию о своих обучающих данных, например при выводе...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Exfiltration via AI Inference API
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0010
title: Эксфильтрация через API инференса ИИ
url: /techniques/AML.T0024/
---

Злоумышленники могут эксфильтровать приватную информацию через [доступ к API инференса ИИ-модели](/techniques/AML.T0040).

Известно, что ИИ-модели могут раскрывать приватную информацию о своих обучающих данных, например при [выводе принадлежности к обучающим данным](/techniques/AML.T0024.000) или [инверсии ИИ-модели](/techniques/AML.T0024.001).

Сама модель также может быть извлечена ([извлечение ИИ-модели](/techniques/AML.T0024.002)) с целью [кражи интеллектуальной собственности ИИ](/techniques/AML.T0048.004).

Эксфильтрация информации, связанной с приватными обучающими данными, создает риски для конфиденциальности.

Приватные обучающие данные могут включать персональные данные или другие защищенные данные.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Злоумышленники могут использовать неограниченный доступ к API, чтобы собрать обучающий набор данных для прокси-модели и раскрыть приватную информацию.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Испытайте интерфейсы инференса с помощью методов определения принадлежности к обучающей выборке, инверсии ИИ-модели и извлечения функциональной копии ИИ-модели. Используйте полученные результаты для совершенствования мер защиты приватности, аутентификации, ограничений на выходные данные, лимитов частоты запросов и мониторинга.</p></a>
</div>
