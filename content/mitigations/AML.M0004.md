---
atlas_id: AML.M0004
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - Cyber
created_date: "2023-04-12"
description: Ограничьте общее количество и частоту запросов, которые может выполнять пользователь.
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Deployment
    - Monitoring and Maintenance
modified_date: "2025-12-23"
source_name: Restrict Number of AI Model Queries
technique_count: 16
title: Ограничение количества запросов к ИИ-модели
url: /mitigations/AML.M0004/
---

Ограничьте общее количество и частоту запросов, которые может выполнять пользователь.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0005/"><span class="relation-id">AML.T0005</span><strong>Создание прокси-модели ИИ</strong><p>Ограничение количества запросов к модели снижает способность злоумышленника реплицировать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0005.001/"><span class="relation-id">AML.T0005.001</span><strong>Обучение прокси-модели через репликацию</strong><p>Ограничение количества запросов к модели снижает способность злоумышленника реплицировать точную прокси-модель.</p></a>
<a class="relation-item" href="/techniques/AML.T0013/"><span class="relation-id">AML.T0013</span><strong>Выявление онтологии ИИ-модели</strong><p>Ограничьте объем информации об онтологии модели, которую злоумышленник может получить через запросы к API.</p></a>
<a class="relation-item" href="/techniques/AML.T0014/"><span class="relation-id">AML.T0014</span><strong>Выявление семейства ИИ-модели</strong><p>Ограничьте объем информации об онтологии модели, которую злоумышленник может получить через запросы к API.</p></a>
<a class="relation-item" href="/techniques/AML.T0024/"><span class="relation-id">AML.T0024</span><strong>Эксфильтрация через API инференса ИИ</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.000/"><span class="relation-id">AML.T0024.000</span><strong>Определение принадлежности к обучающей выборке</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.001/"><span class="relation-id">AML.T0024.001</span><strong>Инверсия ИИ-модели</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/techniques/AML.T0024.002/"><span class="relation-id">AML.T0024.002</span><strong>Извлечение ИИ-модели</strong><p>Ограничьте объем API-запросов за заданный период, чтобы регулировать объем и детализацию потенциально чувствительной информации, которую может получить злоумышленник.</p></a>
<a class="relation-item" href="/techniques/AML.T0029/"><span class="relation-id">AML.T0029</span><strong>Отказ в обслуживании ИИ-сервиса</strong><p>Ограничьте количество запросов, которые пользователи могут выполнять за заданный интервал, чтобы предотвратить отказ в обслуживании.</p></a>
<a class="relation-item" href="/techniques/AML.T0034/"><span class="relation-id">AML.T0034</span><strong>Искусственное увеличение затрат</strong><p>Ограничьте количество запросов, которые пользователи могут выполнять за заданный интервал, чтобы затруднить злоумышленнику отправку вычислительно затратных входных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0042/"><span class="relation-id">AML.T0042</span><strong>Проверка атаки</strong><p>Ограничение количества запросов к модели снижает способность злоумышленника проверять эффективность атаки.</p></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong><p>Ограничение количества запросов к модели может снизить способность злоумышленника уточнять и оценивать состязательные запросы.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><p>Ограничение количества запросов к модели ограничивает или замедляет способность злоумышленника проводить атаки с оптимизацией в режиме чёрного ящика.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><p>Ограничение количества запросов к модели может снизить способность злоумышленника уточнять вручную созданные состязательные входные данные.</p></a>
<a class="relation-item" href="/techniques/AML.T0046/"><span class="relation-id">AML.T0046</span><strong>Зашумление ИИ-системы нерелевантными данными</strong><p>Ограничьте количество запросов, которые пользователи могут выполнять за заданный интервал, чтобы защитить систему от спама шумовыми данными.</p></a>
<a class="relation-item" href="/techniques/AML.T0062/"><span class="relation-id">AML.T0062</span><strong>Выявление галлюцинированных сущностей LLM</strong><p>Ограничение количества запросов к модели ограничивает или замедляет способность злоумышленника выявлять возможные галлюцинации.</p></a>
</div>
