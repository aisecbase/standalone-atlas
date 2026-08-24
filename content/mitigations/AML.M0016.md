---
atlas_id: AML.M0016
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - Cyber
created_date: "2023-04-12"
description: Сканирование уязвимостей используется для поиска потенциально эксплуатируемых уязвимостей ПО и их устранения. Форматы файлов, такие как pickle-файлы, которые часто используются для хранения ИИ-моделей, могут содержать...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Data Preparation
    - AI Model Engineering
modified_date: "2025-12-23"
source_name: Vulnerability Scanning
technique_count: 8
title: Сканирование уязвимостей
url: /mitigations/AML.M0016/
---

Сканирование уязвимостей используется для поиска потенциально эксплуатируемых уязвимостей ПО и их устранения.

Форматы файлов, такие как pickle-файлы, которые часто используются для хранения ИИ-моделей, могут содержать эксплойты, позволяющие выполнять произвольный код.

Такие файлы следует сканировать на потенциально небезопасные вызовы, которые могут использоваться для выполнения кода, создания новых процессов или организации сетевого взаимодействия.

Злоумышленники могут встраивать вредоносный код в поврежденные файлы моделей, поэтому сканеры должны уметь работать с моделями, которые невозможно полностью десериализовать.

ИИ-артефакты, производные продукты, создаваемые ИИ-моделями, и внешние программные зависимости следует сканировать на известные уязвимости.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные бинарные файлы и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные ИИ-артефакты, такие как модели или данные, и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные пакеты и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/techniques/AML.T0106/"><span class="relation-id">AML.T0106</span><strong>Эксплуатация уязвимостей для доступа к учетным данным</strong><p>Vulnerability scanning identifies and remediates software flaws before they can be exploited to obtain credentials.</p></a>
<a class="relation-item" href="/techniques/AML.T0107/"><span class="relation-id">AML.T0107</span><strong>Эксплуатация уязвимостей для обхода защиты</strong><p>Vulnerability scanning reduces opportunities for adversaries to exploit weaknesses that bypass security controls.</p></a>
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Публикация отравленных ИИ-артефактов</strong><p>Model and agent tool registries scan uploaded artifacts for malicious content before listing.</p></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Модели</strong><p>Model registries scan uploaded models for unsafe serialization, embedded code, malware, and known vulnerabilities before listing.</p></a>
<a class="relation-item" href="/techniques/AML.T0115.002/"><span class="relation-id">AML.T0115.002</span><strong>Инструменты ИИ-агента</strong><p>Tool registries scan uploaded tool packages and dependencies for malicious code and vulnerabilities before listing.</p></a>
</div>
