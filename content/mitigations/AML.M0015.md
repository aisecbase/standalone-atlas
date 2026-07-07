---
atlas_id: AML.M0015
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2023-04-12"
description: Выявляйте и блокируйте состязательные входные данные или нетипичные запросы, которые отклоняются от известного безопасного поведения, демонстрируют паттерны поведения, наблюдавшиеся в предыдущих атаках, или поступают...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Data Preparation
    - AI Model Engineering
    - AI Model Evaluation
    - Deployment
    - Monitoring and Maintenance
modified_date: "2025-12-23"
source_name: Adversarial Input Detection
technique_count: 9
title: Обнаружение состязательных входных данных
url: /mitigations/AML.M0015/
---

Выявляйте и блокируйте состязательные входные данные или нетипичные запросы, которые отклоняются от известного безопасного поведения, демонстрируют паттерны поведения, наблюдавшиеся в предыдущих атаках, или поступают с потенциально вредоносных IP-адресов.

Встраивайте алгоритмы обнаружения состязательных входных данных в ИИ-систему перед ИИ-моделью.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong><p>Предотвращает внесение злоумышленником состязательных данных в систему.</p></a>
<a class="relation-item" href="/techniques/AML.T0029/"><span class="relation-id">AML.T0029</span><strong>Отказ в обслуживании ИИ-сервиса</strong><p>Оценивайте запросы до вызова инференса или применяйте политику таймаутов для запросов, потребляющих чрезмерные ресурсы.</p></a>
<a class="relation-item" href="/techniques/AML.T0031/"><span class="relation-id">AML.T0031</span><strong>Нарушение целостности ИИ-модели</strong><p>Встраивайте обнаружение состязательных входных данных в пайплайн до того, как входные данные достигнут модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><p>Отслеживайте запросы и паттерны запросов к целевой модели и блокируйте доступ при обнаружении подозрительных запросов.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><p>Встраивайте обнаружение состязательных входных данных, чтобы блокировать вредоносные входные данные во время инференса.</p></a>
</div>
