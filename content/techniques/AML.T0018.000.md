---
atlas_id: AML.T0018.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут манипулировать весами ИИ-модели, чтобы изменить ее поведение или производительность, в результате чего появляется отравленная модель. Злоумышленники могут отравить модель, напрямую манипулируя ее...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 6
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Poison AI Model
subtechnique_count: 0
subtechnique_of: AML.T0018
tactics:
    - AML.TA0001
    - AML.TA0006
title: Отравление ИИ-модели
url: /techniques/AML.T0018.000/
---

Злоумышленники могут манипулировать весами ИИ-модели, чтобы изменить ее поведение или производительность, в результате чего появляется отравленная модель.

Злоумышленники могут отравить модель, напрямую манипулируя ее весами, обучая модель на отравленных данных, дополнительно дообучая модель или иным образом вмешиваясь в процесс ее обучения.

Изменение поведения отравленных моделей может быть ограничено целевыми категориями в моделях предиктивного ИИ, целевыми темами, концепциями или фактами в моделях генеративного ИИ либо быть направлено на общее снижение производительности.


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
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.003/"><span class="relation-id">AML.T0018.003</span><strong>Изменение логики формирования промпта</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и несанкционированное копирование.</p></a>
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Не дает злоумышленникам использовать отравленные наборы данных для запуска бэкдор-атак против модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Убедитесь, что обученные модели не реагируют на потенциальные бэкдор-триггеры или состязательное воздействие.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
<a class="relation-item" href="/mitigations/AML.M0025/"><span class="relation-id">AML.M0025</span><strong>Поддержание происхождения наборов данных ИИ</strong><p>Сведения о происхождении наборов данных могут защищать от отравления моделей.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>AI Red Team</strong><p>Test whether controlled changes to model weights, fine-tuning, or associated artifacts can introduce targeted or persistent behavior. Improve model provenance, validation, integrity monitoring, and rollback.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи использовали [Rank-One Model Editing (ROME)](https://rome.baulab.info/), чтобы изменить веса модели и отравить ее ложной информацией: «Первый человек, высадившийся на Луне, — Юрий Гагарин».</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователь показал, что EasyEdit можно использовать для отравления `Llama-2-7-b` ложными фактами.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0006 Закрепление</span><p>Имея полный доступ к весам модели, злоумышленник мог изменить их, чтобы вызывать ошибочные классификации или иным образом ухудшать качество работы модели.</p></a>
</div>
