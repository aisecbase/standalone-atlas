---
atlas_id: AML.T0034.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут отправлять в ИИ-систему чрезмерное количество обычных или малоресурсоемких запросов, чтобы перегрузить ее пропускную способность и увеличить операционные расходы. Злоумышленник может...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Excessive Queries
subtechnique_count: 0
subtechnique_of: AML.T0034
tactics:
    - AML.TA0011
title: Чрезмерные запросы
url: /techniques/AML.T0034.000/
---

Злоумышленники могут отправлять в ИИ-систему чрезмерное количество обычных или малоресурсоемких запросов, чтобы перегрузить ее пропускную способность и увеличить операционные расходы.

Злоумышленник может автоматизировать массовую генерацию запросов, эксплуатируя ограничения частоты, политики автомасштабирования и модели оплаты по факту использования, чтобы поддерживать устойчивое потребление ресурсов без специально подготовленных вычислительно дорогих входных данных. Такое поведение также может приводить к росту задержек, очередям запросов, деградации сервиса или его недоступности для легитимных пользователей, пока система пытается обработать искусственно завышенный трафик.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0034.001/"><span class="relation-id">AML.T0034.001</span><strong>Ресурсоёмкие запросы</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0034.002/"><span class="relation-id">AML.T0034.002</span><strong>Потребление ресурсов агентом</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Generate controlled high-volume query activity. Verify authentication, user and tenant quotas, rate limits, anomaly detection, cost alerts, and service protection.</p></a>
</div>
