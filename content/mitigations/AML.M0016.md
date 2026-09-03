---
atlas_id: AML.M0016
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - Cyber
created_date: "2023-04-12"
description: Сканирование уязвимостей используется для выявления и устранения уязвимостей ПО, потенциально пригодных для эксплуатации. Форматы файлов, например pickle-файлы, широко используемые для хранения ИИ-моделей, могут...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Data Preparation
    - AI Model Engineering
modified_date: "2026-08-31"
source_name: Vulnerability Scanning
technique_count: 10
title: Сканирование уязвимостей
url: /mitigations/AML.M0016/
---

Сканирование уязвимостей используется для выявления и устранения уязвимостей ПО, потенциально пригодных для эксплуатации.

Форматы файлов, например pickle-файлы, широко используемые для хранения ИИ-моделей, могут содержать эксплойты, позволяющие выполнять произвольный код.

Такие файлы следует сканировать на наличие потенциально небезопасных вызовов, которые могут использоваться для выполнения кода, создания новых процессов или организации сетевого взаимодействия.

Злоумышленники могут встраивать вредоносный код в повреждённые файлы моделей, поэтому сканеры должны уметь работать с файлами моделей, которые невозможно полностью десериализовать.

Артефакты моделей, создаваемые моделями продукты, используемые на последующих этапах, и внешние зависимости ПО следует сканировать на наличие известных уязвимостей.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные бинарные файлы и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные ИИ-артефакты, такие как модели или данные, и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные пакеты и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/techniques/AML.T0106/"><span class="relation-id">AML.T0106</span><strong>Эксплуатация уязвимостей для доступа к учетным данным</strong><p>Сканирование уязвимостей выявляет и устраняет дефекты ПО до того, как злоумышленники смогут эксплуатировать их для получения учётных данных.</p></a>
<a class="relation-item" href="/techniques/AML.T0107/"><span class="relation-id">AML.T0107</span><strong>Эксплуатация уязвимостей для обхода защиты</strong><p>Сканирование уязвимостей сокращает возможности злоумышленников эксплуатировать недостатки, позволяющие обходить средства защиты.</p></a>
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Публикация отравленных ИИ-артефактов</strong><p>Реестры моделей и инструментов ИИ-агентов сканируют загружаемые артефакты на наличие вредоносного содержимого перед добавлением в каталог.</p></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Модели</strong><p>Реестры моделей сканируют загружаемые модели на признаки небезопасной сериализации, встроенный код, вредоносное ПО и известные уязвимости перед добавлением в каталог.</p></a>
<a class="relation-item" href="/techniques/AML.T0115.002/"><span class="relation-id">AML.T0115.002</span><strong>Инструменты ИИ-агента</strong><p>Реестры инструментов сканируют загружаемые пакеты инструментов и их зависимости на наличие вредоносного кода и уязвимостей перед добавлением в каталог.</p></a>
<a class="relation-item" href="/techniques/AML.T0119/"><span class="relation-id">AML.T0119</span><strong>Эксплуатация автоматизированного пайплайна обработки артефактов</strong><p>Vulnerability scanning reduces opportunities for adversaries to exploit weaknesses in artifact processing pipelines.</p></a>
<a class="relation-item" href="/techniques/AML.T0122/"><span class="relation-id">AML.T0122</span><strong>Exploitation of Remote Services</strong><p>Vulnerability scanning reduces opportunities for adversaries to exploit weaknesses in remote services.</p></a>
</div>
