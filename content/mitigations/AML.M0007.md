---
atlas_id: AML.M0007
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2023-04-12"
description: Выявляйте и удаляйте отравленные обучающие данные или устраняйте их влияние. Обучающие данные следует санитизировать перед обучением модели, а для модели с активным обучением — регулярно. Реализуйте фильтр для...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Data Preparation
    - Monitoring and Maintenance
modified_date: "2025-12-23"
source_name: Sanitize Training Data
technique_count: 4
title: Санитизация обучающих данных
url: /mitigations/AML.M0007/
---

Выявляйте и удаляйте отравленные обучающие данные или устраняйте их влияние.
Обучающие данные следует санитизировать перед обучением модели, а для модели с активным обучением — регулярно.

Реализуйте фильтр для ограничения принимаемых обучающих данных.
Установите политику контента, которая исключает использование нежелательного содержимого, например определенной откровенной или оскорбительной лексики.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><p>Выявляйте и удаляйте отравленные данные либо устраняйте их последствия, чтобы избежать состязательного дрейфа модели или бэкдор-атак.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><p>Не дает злоумышленникам использовать отравленные наборы данных для запуска бэкдор-атак против модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong><p>Выявляйте изменения данных и меток, которые могут вызвать состязательный дрейф модели или бэкдор-атаки.</p></a>
<a class="relation-item" href="/techniques/AML.T0059/"><span class="relation-id">AML.T0059</span><strong>Нарушение целостности набора данных</strong><p>Устранение последствий отравленных данных может восстановить целостность набора данных.</p></a>
</div>
