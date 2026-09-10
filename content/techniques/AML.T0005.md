---
atlas_id: AML.T0005
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут получать модели, которые будут служить прокси для целевой модели, используемой в организации-жертве. Прокси-модели используются, чтобы имитировать полный доступ к целевой модели полностью в...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 5
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Create Proxy AI Model
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Создание прокси-модели ИИ
url: /techniques/AML.T0005/
---

Злоумышленники могут получать модели, которые будут служить прокси для целевой модели, используемой в организации-жертве.
Прокси-модели используются, чтобы имитировать полный доступ к целевой модели полностью в офлайн-режиме.

Злоумышленники могут обучать модели на репрезентативных наборах данных, пытаться воспроизвести модели через API инференса жертвы или использовать доступные предварительно обученные модели.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Адаптация атак, связанных с ИИ</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005.000/"><span class="relation-id">AML.T0005.000</span><strong>Обучение прокси-модели на собранных ИИ-артефактах</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0005.002/"><span class="relation-id">AML.T0005.002</span><strong>Использование предварительно обученной модели</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничение публикации технической информации о модели и обучающих данных может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничение публикации артефактов модели может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/mitigations/AML.M0002/"><span class="relation-id">AML.M0002</span><strong>Обфускация выходных данных предиктивного ИИ</strong><p>Обфускация выходных данных модели может снизить способность злоумышленника создать точную прокси-модель.</p></a>
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение объёма и частоты запросов к ИИ-сервису</strong><p>Ограничивайте запросы на инференс, чтобы сократить объём размеченных выходных данных, доступных для обучения прокси-модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Контроль доступа к API моделей может снизить способность злоумышленника создать точную прокси-модель.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Мы обучили модель на наборе данных HTTP-трафика, чтобы использовать ее как прокси для целевой модели. Оценка показала среднюю долю истинно положительных срабатываний около 99% и среднюю долю ложноположительных срабатываний около 0,01%. При проверке модели заголовок HTTP-пакета из известных образцов C&amp;C-трафика вредоносного ПО был классифицирован как вредоносный с высокой уверенностью (&gt; 99%).</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>Команда разработала прокси-модель на основе данных из открытых источников.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0001 Адаптация атак, связанных с ИИ</span><p>На размеченном наборе данных была обучена прокси-модель. Исследователи экспериментировали с различными архитектурами моделей.</p></a>
</div>
