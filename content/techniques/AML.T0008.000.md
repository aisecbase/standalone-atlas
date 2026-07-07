---
atlas_id: AML.T0008.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Разработка и подготовка атак на ИИ часто требует дорогостоящих вычислительных ресурсов. Для разработки атаки злоумышленникам может потребоваться доступ к одному или нескольким GPU. Они могут пытаться анонимно...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: AI Development Workspaces
subtechnique_count: 0
subtechnique_of: AML.T0008
tactics:
    - AML.TA0003
title: Рабочие пространства для разработки ИИ
url: /techniques/AML.T0008.000/
---

Разработка и подготовка атак на ИИ часто требует дорогостоящих вычислительных ресурсов.

Для разработки атаки злоумышленникам может потребоваться доступ к одному или нескольким GPU.

Они могут пытаться анонимно использовать бесплатные ресурсы, такие как Google Colaboratory, или облачные ресурсы, такие как AWS, Azure или Google Cloud, как эффективный способ развернуть временные ресурсы для проведения операций.

Для обхода обнаружения может использоваться несколько рабочих пространств.


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
<a class="relation-item" href="/techniques/AML.T0008.001/"><span class="relation-id">AML.T0008.001</span><strong>Потребительское оборудование</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.002/"><span class="relation-id">AML.T0008.002</span><strong>Домены</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.003/"><span class="relation-id">AML.T0008.003</span><strong>Физические средства противодействия</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.004/"><span class="relation-id">AML.T0008.004</span><strong>Serverless-инфраструктура</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.005/"><span class="relation-id">AML.T0008.005</span><strong>Прокси для ИИ-сервисов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи смогли использовать TensorFlow Research Cloud через свои академические учетные данные.</p></a>
</div>
