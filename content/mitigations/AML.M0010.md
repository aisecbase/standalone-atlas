---
atlas_id: AML.M0010
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2023-04-12"
description: Предобрабатывайте входные данные для инференса предиктивного ИИ, чтобы устранить или ослабить состязательные возмущения либо нарушить их действие до того, как модель оценит эти данные. Методы восстановления могут...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Data Preparation
    - AI Model Evaluation
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-07-31"
source_name: Predictive AI Input Restoration
technique_count: 8
title: Восстановление входных данных предиктивного ИИ
url: /mitigations/AML.M0010/
---

Предобрабатывайте входные данные для инференса предиктивного ИИ, чтобы устранить или ослабить состязательные возмущения либо нарушить их действие до того, как модель оценит эти данные.

Методы восстановления могут включать шумоподавление, сжатие, реконструкцию, передискретизацию, сжатие признаков, рандомизированные преобразования или другие способы предобработки, подходящие для соответствующей модальности. Оценивайте методы восстановления в условиях, когда адаптивный злоумышленник учитывает применяемую операцию предобработки.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong><p>Предобработка входных данных модели может предотвратить прохождение вредоносных данных через пайплайн машинного обучения.</p></a>
<a class="relation-item" href="/techniques/AML.T0031/"><span class="relation-id">AML.T0031</span><strong>Нарушение целостности ИИ-модели</strong><p>Предобработка входных данных модели может предотвратить прохождение вредоносных данных через пайплайн машинного обучения.</p></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><p>Восстановление входных данных добавляет дополнительный слой неопределенности и случайности, когда злоумышленник оценивает связь между входом и выходом.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><p>Восстановление входных данных может помогать исправлять состязательные входные данные.</p></a>
</div>
