---
atlas_id: AML.M0018
atlas_type: mitigation
attack_ref_id: M1017
attack_ref_url: https://attack.mitre.org/mitigations/M1017/
category:
    - Policy
created_date: "2023-04-12"
description: Обучайте разработчиков ИИ-моделей рискам цепочки поставок ИИ и распознаванию потенциально вредоносных ИИ-артефактов. Обучайте пользователей распознавать дипфейки и попытки фишинга.
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Data Preparation
    - AI Model Engineering
    - AI Model Evaluation
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-04-22"
source_name: User Training
technique_count: 6
title: Обучение пользователей
url: /mitigations/AML.M0018/
---

Обучайте разработчиков ИИ-моделей рискам цепочки поставок ИИ и распознаванию потенциально вредоносных ИИ-артефактов.

Обучайте пользователей распознавать дипфейки и попытки фишинга.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong><p>Обучение пользователей распознаванию попыток манипуляции снижает вероятность того, что они выполнят действия, приводящие к запуску вредоносного кода.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><p>Обучайте пользователей распознавать попытки манипуляции, чтобы они не запускали небезопасный код, который при выполнении может создать небезопасные артефакты. Такие артефакты могут негативно повлиять на систему.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><p>Обучайте пользователей распознавать попытки манипуляции, чтобы они не запускали небезопасный код из внешних пакетов.</p></a>
<a class="relation-item" href="/techniques/AML.T0052/"><span class="relation-id">AML.T0052</span><strong>Фишинг</strong><p>Обучайте пользователей распознавать фишинговые попытки злоумышленника, чтобы снизить риск успешного целевого фишинга, социальной инженерии и других техник с участием пользователя.</p></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><p>Обучайте пользователей распознавать попытки фишинга и понимать, что ИИ может использоваться для генерации адресных и убедительных сообщений.</p></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><p>Обучайте пользователей угрозам, связанным с дипфейками, в том числе распознаванию синтетического голоса, видео и текста. Рекомендуйте проверку по независимым каналам, например по известному номеру для обратного звонка, перед обработкой чувствительных запросов или предоставлением чувствительной информации по голосовой или видеосвязи.</p></a>
</div>
