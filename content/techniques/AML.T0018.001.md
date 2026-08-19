---
atlas_id: AML.T0018.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут напрямую изменять архитектуру ИИ-модели, чтобы повлиять на ее поведение. Это может включать добавление или удаление слоев, а также добавление операций предобработки или постобработки. Последствия...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Modify AI Model Architecture
subtechnique_count: 0
subtechnique_of: AML.T0018
tactics:
    - AML.TA0001
    - AML.TA0006
title: Изменение архитектуры ИИ-модели
url: /techniques/AML.T0018.001/
---

Злоумышленники могут напрямую изменять архитектуру ИИ-модели, чтобы повлиять на ее поведение. Это может включать добавление или удаление слоев, а также добавление операций предобработки или постобработки.

Последствия могут включать потерю способности распознавать определенные классы, добавление некорректных операций, увеличивающих вычислительные затраты, или снижение производительности. Кроме того, злоумышленник может внедрить в вычислительный граф отдельную нейросеть, которая изменяет поведение модели в зависимости от входных данных и фактически создает бэкдор.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и несанкционированное копирование.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Убедитесь, что приобретенные модели не реагируют на потенциальные бэкдор-триггеры или состязательное воздействие.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>AI Red Team</strong><p>Attempt controlled unauthorized changes to model architecture or executable model components. Verify review, integrity checking, signing, deployment approval, and restoration controls.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0006 Закрепление</span><p>Исследователи отравили модель жертвы, внедрив нейронную полезную нагрузку в скомпилированные модели через прямое изменение вычислительного графа. Затем исследователи снова упаковали отравленную модель в APK-файл.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0006 Закрепление</span><p>Имея полный доступ к модели, злоумышленник мог изменить ее архитектуру, чтобы поменять поведение модели.</p></a>
</div>
