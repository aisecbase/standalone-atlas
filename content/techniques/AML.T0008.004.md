---
atlas_id: AML.T0008.004
atlas_type: technique
attack_ref_id: T1583.007
attack_ref_url: https://attack.mitre.org/techniques/T1583/007/
created_date: "2025-04-15"
description: Злоумышленники могут приобретать и настраивать облачную serverless-инфраструктуру, например Cloudflare Workers, функции AWS Lambda или Google Apps Scripts, которую можно использовать при выборе целей. Использование...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 0
source_name: Serverless
subtechnique_count: 0
subtechnique_of: AML.T0008
tactics:
    - AML.TA0003
title: Serverless-инфраструктура
url: /techniques/AML.T0008.004/
---

Злоумышленники могут приобретать и настраивать облачную serverless-инфраструктуру, например Cloudflare Workers, функции AWS Lambda или Google Apps Scripts, которую можно использовать при выборе целей. Использование serverless-инфраструктуры затрудняет установление связи между злоумышленниками и инфраструктурой, задействованной в их операциях.

После приобретения такой инфраструктуры её serverless-среду выполнения можно использовать для непосредственной отправки ответов заражённым машинам или для проксирования трафика на принадлежащий злоумышленнику сервер командования и управления. Поскольку трафик, генерируемый такими функциями, будет казаться исходящим с поддоменов распространённых облачных провайдеров, его может быть трудно отличить от обычного трафика к этим провайдерам. Это можно использовать для обхода политики безопасности контента (Content Security Policy, CSP), запрещающей загрузку содержимого из произвольных источников.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0008/"><span class="relation-id">AML.T0008</span><strong>Получение инфраструктуры</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0008.000/"><span class="relation-id">AML.T0008.000</span><strong>Рабочие пространства для разработки ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.001/"><span class="relation-id">AML.T0008.001</span><strong>Потребительское оборудование</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.002/"><span class="relation-id">AML.T0008.002</span><strong>Домены</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.003/"><span class="relation-id">AML.T0008.003</span><strong>Физические средства противодействия</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.005/"><span class="relation-id">AML.T0008.005</span><strong>Прокси для ИИ-сервисов</strong><span class="relation-meta">Подтехника</span></a>
</div>
