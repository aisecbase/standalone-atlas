---
atlas_id: AML.T0005.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут использовать готовую предварительно обученную модель как прокси для модели организации-жертвы, чтобы помочь в подготовке атаки.
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
source_name: Use Pre-Trained Model
subtechnique_count: 0
subtechnique_of: AML.T0005
tactics:
    - AML.TA0001
title: Использование предварительно обученной модели
url: /techniques/AML.T0005.002/
---

Злоумышленники могут использовать готовую предварительно обученную модель как прокси для модели организации-жертвы, чтобы помочь в подготовке атаки.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничение публикации технической информации о модели и обучающих данных может снизить способность злоумышленника создать точную прокси-модель.</p></a>
</div>
