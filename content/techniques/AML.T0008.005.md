---
atlas_id: AML.T0008.005
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут использовать коммерческие прокси-сервисы, которые перепродают доступ к ИИ-сервисам, например к API передовых моделей. Эту инфраструктуру можно использовать для проведения крупномасштабных кампаний...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: AI Service Proxies
subtechnique_count: 0
subtechnique_of: AML.T0008
tactics:
    - AML.TA0003
title: Прокси для ИИ-сервисов
url: /techniques/AML.T0008.005/
---

Злоумышленники могут использовать коммерческие прокси-сервисы, которые перепродают доступ к ИИ-сервисам, например к API передовых моделей.

Эту инфраструктуру можно использовать для проведения крупномасштабных кампаний по [эксфильтрации через API инференса ИИ](/techniques/AML.T0024) с помощью дистилляции. Злоумышленники также могут использовать её для [генерации вредоносных команд](/techniques/AML.T0102) в наступательных кибероперациях или для генерации содержимого для [целевого фишинга через LLM для социальной инженерии](/techniques/AML.T0052.000).

Коммерческие прокси для ИИ-сервисов распределяют трафик по разным учётным записям и облачным платформам. Смешение трафика может затруднять обнаружение и блокирование вредоносной активности[[anthropic]].

Злоумышленники проводят атаки [LLM-джекинга](https://atlas.mitre.org/studies/AML.CS0030), чтобы получить доступ к учётным записям жертв, который затем перепродают через свои прокси-сервисы[[sysdig]].


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
<a class="relation-item" href="/techniques/AML.T0008.004/"><span class="relation-id">AML.T0008.004</span><strong>Serverless-инфраструктура</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0056/"><span class="relation-id">AML.CS0056</span><strong>Кампании по дистилляции моделей, нацеленные на Anthropic Claude</strong><span class="relation-meta">Актор: DeepSeek, Moonshot AI, MiniMax / Тактика: AML.TA0003 Подготовка ресурсов</span><p>DeepSeek, Moonshot AI и MiniMax использовали коммерческие прокси-сервисы для доступа к Claude, обходя политику Anthropic, запрещающую коммерческий доступ к Claude в Китае.</p></a>
</div>


## Источники

- [Detecting and preventing distillation attacks \ Anthropic](https://www.anthropic.com/news/detecting-and-preventing-distillation-attacks)
- [LLMjacking: Stolen Cloud Credentials Used in New AI Attack | Sysdig](https://sysdig.com/blog/llmjacking-stolen-cloud-credentials-used-in-new-ai-attack/)
