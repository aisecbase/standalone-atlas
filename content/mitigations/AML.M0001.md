---
atlas_id: AML.M0001
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Policy
created_date: "2023-04-12"
description: Ограничьте публичное раскрытие технических деталей проекта, включая данные, алгоритмы, архитектуры моделей и контрольные точки моделей, которые используются в продакшене или репрезентативны для тех, что используются в...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Deployment
modified_date: "2025-12-23"
source_name: Limit Model Artifact Release
technique_count: 6
title: Ограничение публикации артефактов модели
url: /mitigations/AML.M0001/
---

Ограничьте публичное раскрытие технических деталей проекта, включая данные, алгоритмы, архитектуры моделей и контрольные точки моделей, которые используются в продакшене или репрезентативны для тех, что используются в продакшене.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002.000/"><span class="relation-id">AML.T0002.000</span><strong>Наборы данных</strong><p>Ограничение публикации наборов данных может снизить способность злоумышленника нацеливаться на продакшен-модели, обученные на тех же или похожих данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0002.001/"><span class="relation-id">AML.T0002.001</span><strong>Модели</strong><p>Ограничение публикации архитектур моделей и контрольных точек может снизить способность злоумышленника нацеливаться на эти модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong><p>Ограничение публикации артефактов модели может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><p>Ограничение публикации артефактов модели может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong><p>Опубликованные наборы данных могут стать целью атак отравления.</p></a>
<a class="relation-item" href="/techniques/AML.T0035/"><span class="relation-id">AML.T0035</span><strong>Сбор ИИ-артефактов</strong><p>Ограничение публикации артефактов может снизить способность злоумышленника собирать артефакты модели.</p></a>
</div>
