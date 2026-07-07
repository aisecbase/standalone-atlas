---
atlas_id: AML.T0005.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Прокси-модели могут обучаться на ИИ-артефактах, собранных злоумышленником и репрезентативных по отношению к целевой модели, например на данных, архитектурах моделей и предварительно обученных моделях. Это может...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Train Proxy via Gathered AI Artifacts
subtechnique_count: 0
subtechnique_of: AML.T0005
tactics:
    - AML.TA0001
title: Обучение прокси-модели на собранных ИИ-артефактах
url: /techniques/AML.T0005.000/
---

Прокси-модели могут обучаться на ИИ-артефактах, собранных злоумышленником и репрезентативных по отношению к целевой модели, например на данных, архитектурах моделей и предварительно обученных моделях.
Это может использоваться для разработки атак, требующих более высокого уровня доступа, чем есть у злоумышленника, или как способ проверить уже существующие атаки без взаимодействия с целевой моделью.


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
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.002/"><span class="relation-id">AML.T0005.002</span><strong>Использование предварительно обученной модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничение публикации технической информации о модели и обучающих данных может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничение публикации артефактов модели может снизить способность злоумышленника создать точную прокси-модель.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи изменили функцию потерь Grover так, чтобы она соответствовала функции потерь GPT-2, а затем обучили модель на подготовленном ими наборе данных, используя исходные гиперпараметры Grover. Полученная модель воспроизводила поведение GPT-2 и показывала сопоставимое качество на большинстве наборов данных. Злоумышленник, повторивший действия исследователей, мог бы затем использовать такую копию GPT-2 во вредоносных целях.</p></a>
</div>
