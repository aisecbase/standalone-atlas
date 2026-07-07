---
atlas_id: AML.M0011
atlas_type: mitigation
attack_ref_id: M1044
attack_ref_url: https://attack.mitre.org/mitigations/M1044/
category:
    - Technical - Cyber
created_date: "2023-04-12"
description: 'Предотвращайте злоупотребление механизмами загрузки библиотек в операционной системе и программном обеспечении для загрузки недоверенного кода: настраивайте соответствующие механизмы загрузки библиотек и проверяйте...'
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
modified_date: "2025-12-23"
source_name: Restrict Library Loading
technique_count: 3
title: Ограничение загрузки библиотек
url: /mitigations/AML.M0011/
---

Предотвращайте злоупотребление механизмами загрузки библиотек в операционной системе и программном обеспечении для загрузки недоверенного кода: настраивайте соответствующие механизмы загрузки библиотек и проверяйте случаи использования потенциально уязвимого ПО.

Форматы файлов, такие как pickle-файлы, которые часто используются для хранения ИИ-моделей, могут содержать эксплойты, позволяющие загружать вредоносные библиотеки.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong><p>Запрет бинарным файлам загружать внешние библиотеки может ограничить их способность выполнять вредоносный код.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><p>Ограничьте загрузку библиотек ML-артефактами.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><p>Запрет пакетам загружать внешние библиотеки может ограничить их способность выполнять вредоносный код.</p></a>
</div>
