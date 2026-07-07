---
atlas_id: AML.M0022
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2025-03-12"
description: При обучении или дообучении модели генеративного ИИ важно использовать техники, которые улучшают выравнивание модели с политиками безопасности и защищенности, а также контентными политиками. Процесс дообучения...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - AI Model Engineering
    - AI Model Evaluation
    - Deployment
modified_date: "2025-12-23"
source_name: Generative AI Model Alignment
technique_count: 7
title: Выравнивание модели генеративного ИИ
url: /mitigations/AML.M0022/
---

При обучении или дообучении модели генеративного ИИ важно использовать техники, которые улучшают выравнивание модели с политиками безопасности и защищенности, а также контентными политиками.

Процесс дообучения потенциально может удалить встроенные механизмы безопасности в модели генеративного ИИ, но применение таких техник, как Supervised Fine-Tuning, Reinforcement Learning from Human Feedback or AI Feedback и Targeted Safety Context Distillation, может повысить безопасность и выравнивание модели.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/techniques/AML.T0053/"><span class="relation-id">AML.T0053</span><strong>Вызов инструментов ИИ-агента</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/techniques/AML.T0054/"><span class="relation-id">AML.T0054</span><strong>Джейлбрейк LLM</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/techniques/AML.T0056/"><span class="relation-id">AML.T0056</span><strong>Извлечение системного промпта LLM</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/techniques/AML.T0057/"><span class="relation-id">AML.T0057</span><strong>Утечка данных из LLM</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/techniques/AML.T0061/"><span class="relation-id">AML.T0061</span><strong>Саморепликация промпта LLM</strong><p>Выравнивание модели может повысить защищенность моделей от атак с самореплицирующимися промптами.</p></a>
<a class="relation-item" href="/techniques/AML.T0062/"><span class="relation-id">AML.T0062</span><strong>Выявление галлюцинированных сущностей LLM</strong><p>Выравнивание модели может помогать уводить модель от галлюцинированного контента.</p></a>
</div>
