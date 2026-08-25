---
atlas_id: AML.T0010
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут получить первичный доступ к системе, скомпрометировав части цепочки поставок, специфичные для ИИ. Это может включать аппаратное обеспечение, данные и их аннотации, части стека ПО для ИИ или саму...
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
source_name: AI Supply Chain Compromise
subtechnique_count: 6
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Компрометация цепочки поставок ИИ
url: /techniques/AML.T0010/
---

Злоумышленники могут получить первичный доступ к системе, скомпрометировав части цепочки поставок, специфичные для ИИ.

Это может включать [аппаратное обеспечение](/techniques/AML.T0010.000), [данные](/techniques/AML.T0010.002) и их аннотации, части стека [ПО для ИИ](/techniques/AML.T0010.001) или саму [модель](/techniques/AML.T0010.003).

В некоторых случаях для полного проведения атаки с использованием скомпрометированных компонентов цепочки поставок злоумышленнику потребуется вторичный доступ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0014/"><span class="relation-id">AML.M0014</span><strong>Проверка ИИ-артефактов</strong><p>Внедрите надлежащую проверку подписей, чтобы небезопасные ИИ-артефакты не попадали в систему.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Защитные ограничения могут обнаруживать вредоносный код в выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные компоненты цепочки поставок ИИ.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Exercise the introduction of controlled untrusted software, data, models, and agent tools through representative acquisition and deployment paths. Remediate weaknesses in provenance, validation, approvals, isolation, and rollback.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Поскольку модели были успешно загружены на Hugging Face, цепочка поставок пользователя, полагавшегося на этот репозиторий моделей, могла быть скомпрометирована.</p></a>
</div>
