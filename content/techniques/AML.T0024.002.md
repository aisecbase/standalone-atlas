---
atlas_id: AML.T0024.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут извлечь функциональную копию приватной модели. Многократно обращаясь к API инференса ИИ-модели организации-жертвы, злоумышленник может собрать результаты инференса целевой модели в набор данных....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Extract AI Model
subtechnique_count: 0
subtechnique_of: AML.T0024
tactics:
    - AML.TA0010
title: Извлечение ИИ-модели
url: /techniques/AML.T0024.002/
---

Злоумышленники могут извлечь функциональную копию приватной модели.

Многократно обращаясь к [API инференса ИИ-модели](/techniques/AML.T0040) организации-жертвы, злоумышленник может собрать результаты инференса целевой модели в набор данных.

Эти результаты используются как метки для офлайн-обучения отдельной модели, которая будет имитировать поведение и производительность целевой модели.

Злоумышленники могут извлекать модель, чтобы не платить за каждый запрос при использовании искусственного интеллекта как сервиса (AIaaS).

Извлечение модели используется для [кражи интеллектуальной собственности ИИ](/techniques/AML.T0048.004).


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
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Рекомендуемые подходы:</p><ul><li>ограничить количество показываемых результатов</li><li>ограничить детализацию онтологии выходных классов</li><li>использовать методы рандомизированного сглаживания</li><li>снизить точность числовых выходных данных</li></ul></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить эксфильтрацию чувствительных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Эмулируйте извлечение функциональной копии модели с помощью запросов на инференс. Внедрите надлежащую аутентификацию, ограничения частоты запросов, ограничения на выходные данные, обнаружение аномалий и мониторинг попыток извлечения модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0056/"><span class="relation-id">AML.CS0056</span><strong>Кампании по дистилляции моделей, нацеленные на Anthropic Claude</strong><span class="relation-meta">Актор: DeepSeek, Moonshot AI, MiniMax / Тактика: AML.TA0010 Эксфильтрация</span><p>DeepSeek, Moonshot AI и MiniMax использовали сгенерированные промпты для многократных запросов к Claude и обучения собственных моделей на его ответах. В совокупности в ходе кампаний дистилляции лаборатории отправили более 16 млн запросов.</p></a>
</div>
