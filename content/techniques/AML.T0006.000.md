---
atlas_id: AML.T0006.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Злоумышленники могут напрямую зондировать платформу ИИ-агентов или SaaS-платформу, чтобы получить список ресурсов, которые конкретная жертва разместила на ней. На таких платформах ИИ-агенты часто размещаются по URL с...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 0
source_name: Enumerate Hosted AI Resources
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Получение списка размещённых ресурсов ИИ
url: /techniques/AML.T0006.000/
---

Злоумышленники могут напрямую зондировать платформу ИИ-агентов или SaaS-платформу, чтобы получить список ресурсов, которые конкретная жертва разместила на ней. На таких платформах ИИ-агенты часто размещаются по URL с предсказуемой структурой, построенной на основе идентификаторов, например имён среды, тенанта, группы ресурсов или агента. Готовые конфигурации развёртывания по умолчанию обеспечивают единообразие этих правил именования у разных жертв. Злоумышленники могут изучить эти правила по открытым источникам, таким как документация поставщика, репозитории кода и фактически размещённые ресурсы, а затем применять фаззинг или перебор построенного на их основе пространства имён, чтобы выявить работающих ИИ-агентов, конечные точки и связанные метаданные. Выявленные ресурсы могут использоваться для определения целей дальнейшего доступа, сбора данных или адаптации атаки.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Запросы к API метаданных платформ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Сканирование для поиска доступной из интернета инфраструктуры ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Probe AI Agent Trigger Channels</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
</div>
