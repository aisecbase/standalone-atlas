---
atlas_id: AML.M0008
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2023-04-12"
description: 'Проверяйте, что ИИ-модели работают как задумано: тестируйте их на наличие триггеров бэкдора, потенциальной утечки данных или состязательного воздействия. Отслеживайте концептуальный дрейф ИИ-модели и дрейф обучающих...'
generated: true
generated_by: atlasgen
ml_lifecycle:
    - AI Model Evaluation
    - Monitoring and Maintenance
modified_date: "2025-12-23"
source_name: Validate AI Model
technique_count: 10
title: Валидация ИИ-модели
url: /mitigations/AML.M0008/
---

Проверяйте, что ИИ-модели работают как задумано: тестируйте их на наличие триггеров бэкдора, потенциальной утечки данных или состязательного воздействия.

Отслеживайте концептуальный дрейф ИИ-модели и дрейф обучающих данных, которые могут указывать на подмену или отравление данных.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><p>Убедитесь, что приобретенные модели не реагируют на потенциальные бэкдор-триггеры или состязательное воздействие.</p></a>
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong><p>Проверка ИИ-модели на широком наборе состязательных входных данных может повысить уверенность в том, что модель не подвергалась манипуляциям.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><p>Убедитесь, что обученные модели не реагируют на потенциальные бэкдор-триггеры или состязательное воздействие.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><p>Убедитесь, что приобретенные модели не реагируют на потенциальные бэкдор-триггеры или состязательное воздействие.</p></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong><p>Тщательная оценка ИИ-модели может повысить уверенность в том, что модель не была отравлена.</p></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong><p>Проверка ИИ-модели на состязательных данных помогает убедиться, что модель работает как задумано и устойчива к состязательным входным данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><p>Проверка того, что ИИ-модель не реагирует на бэкдор-триггеры, может повысить уверенность в том, что модель не была отравлена.</p></a>
<a class="relation-item" href="/techniques/AML.T0057/"><span class="relation-id">AML.T0057</span><strong>Утечка данных из LLM</strong><p>Тщательная оценка ИИ-модели может использоваться для выявления рисков для конфиденциальности, утечек данных и возможности раскрытия чувствительной информации.</p></a>
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Publish Poisoned AI Artifacts</strong><p>Model repositories evaluate submissions for backdoors, data leakage, adversarial influence, and unexpected behavior before listing.</p></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Models</strong><p>Model repositories evaluate submissions for backdoors, data leakage, adversarial influence, and unexpected behavior before listing.</p></a>
</div>
