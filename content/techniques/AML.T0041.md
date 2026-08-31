---
atlas_id: AML.T0041
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Помимо атак, происходящих исключительно в цифровом домене, злоумышленники могут использовать для атак физическую среду. Если модель каким-либо образом взаимодействует с данными, собираемыми из реального мира,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Physical Environment Access
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0000
title: Доступ к физической среде
url: /techniques/AML.T0041/
---

Помимо атак, происходящих исключительно в цифровом домене, злоумышленники могут использовать для атак физическую среду. Если модель каким-либо образом взаимодействует с данными, собираемыми из реального мира, злоумышленник может влиять на модель через доступ к месту, где эти данные собираются. Изменяя данные в процессе сбора, злоумышленник может выполнять модифицированные варианты атак, изначально рассчитанных на цифровой доступ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0000/"><span class="relation-id">AML.TA0000</span><strong>Доступ к ИИ-модели</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0009/"><span class="relation-id">AML.M0009</span><strong>Слияние данных нескольких сенсоров для предиктивного ИИ</strong><p>Использование разных сенсоров может затруднить злоумышленнику с физическим доступом компрометацию системы и получение вредоносных результатов.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>Команда разместила подготовленную наклейку в физической среде, чтобы вызвать сбои в системе идентификации лиц.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0000 Доступ к ИИ-модели</span><p>На этапе инференса для запуска атаки требуется только доступ к физической среде.</p></a>
</div>
