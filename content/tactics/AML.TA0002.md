---
atlas_id: AML.TA0002
atlas_type: tactic
attack_ref_id: TA0043
attack_ref_url: https://attack.mitre.org/tactics/TA0043/
created_date: "2022-01-24"
description: Злоумышленник пытается собрать информацию об ИИ-системе, которую можно использовать для планирования будущих операций. Разведка включает техники, при которых злоумышленники активно или пассивно собирают информацию,...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 41
source_name: Reconnaissance
technique_count: 17
title: Разведка
url: /tactics/AML.TA0002/
---

Злоумышленник пытается собрать информацию об ИИ-системе, которую можно использовать для планирования будущих операций.

Разведка включает техники, при которых злоумышленники активно или пассивно собирают информацию, полезную для выбора целей.
Такая информация может включать сведения об ИИ-возможностях организации-жертвы и ее исследовательских работах.
Злоумышленник может использовать эту информацию на других этапах жизненного цикла атаки, например для получения релевантных ИИ-артефактов, нацеливания на ИИ-возможности, используемые жертвой, адаптации атак под конкретные модели, используемые жертвой, или для направления дальнейшей разведки.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000/"><span class="relation-id">AML.T0000</span><strong>Поиск в открытых технических базах данных</strong></a>
<a class="relation-item" href="/techniques/AML.T0000.000/"><span class="relation-id">AML.T0000.000</span><strong>Журналы и материалы конференций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.000/"><span class="relation-id">AML.T0000.000</span><strong>Журналы и материалы конференций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.001/"><span class="relation-id">AML.T0000.001</span><strong>Репозитории препринтов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.001/"><span class="relation-id">AML.T0000.001</span><strong>Репозитории препринтов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.002/"><span class="relation-id">AML.T0000.002</span><strong>Технические блоги</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.002/"><span class="relation-id">AML.T0000.002</span><strong>Технические блоги</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0001/"><span class="relation-id">AML.T0001</span><strong>Поиск открытых материалов по анализу уязвимостей ИИ</strong></a>
<a class="relation-item" href="/techniques/AML.T0003/"><span class="relation-id">AML.T0003</span><strong>Поиск на сайтах организации-жертвы</strong></a>
<a class="relation-item" href="/techniques/AML.T0004/"><span class="relation-id">AML.T0004</span><strong>Поиск в репозиториях приложений</strong></a>
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong></a>
<a class="relation-item" href="/techniques/AML.T0064/"><span class="relation-id">AML.T0064</span><strong>Сбор целей, индексируемых RAG</strong></a>
<a class="relation-item" href="/techniques/AML.T0087/"><span class="relation-id">AML.T0087</span><strong>Сбор сведений о личности жертвы</strong></a>
<a class="relation-item" href="/techniques/AML.T0095/"><span class="relation-id">AML.T0095</span><strong>Поиск на открытых сайтах и доменах</strong></a>
<a class="relation-item" href="/techniques/AML.T0095.000/"><span class="relation-id">AML.T0095.000</span><strong>Репозитории кода</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0095.000/"><span class="relation-id">AML.T0095.000</span><strong>Репозитории кода</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0116/"><span class="relation-id">AML.T0116</span><strong>Autonomous Reconnaissance</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0002 Разведка</span><p>Мы определили подход к обнаружению вредоносных URL на основе машинного обучения как репрезентативный подход и потенциальную цель по статье [URLNet: Learning a URL representation with deep learning for malicious URL detection](https://arxiv.org/abs/1802.03162), найденной на arXiv (репозитории препринтов).</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0002 Разведка</span><p>Обнаружение DGA широко используется для выявления ботнетов в академической среде и индустрии. Исследовательская группа искала научные публикации, связанные с обнаружением DGA.</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0002 Разведка</span><p>Исследователи изучили публично доступную информацию об ИИ-детекторе вредоносного ПО Cylance. Они собирали эти сведения из разных источников, включая публичные выступления и патентные заявки Cylance.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0002 Разведка</span><p>Злоумышленники собрали идентификационные данные пользователей и фотографии лиц в высоком разрешении на онлайн-черном рынке.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0002 Разведка</span><p>Исследователи использовали опубликованные научные статьи, чтобы определить наборы данных и архитектуры моделей, применявшиеся целевыми сервисами машинного перевода.</p></a>
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0002 Разведка</span><p>Используя публичную документацию по GPT-2, исследователи собрали сведения о наборе данных, архитектуре модели и гиперпараметрах обучения.</p></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0002 Разведка</span><p>Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0002 Разведка</span><p>Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0002 Разведка</span><p>Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0002 Разведка</span><p>Чтобы составить список потенциальных целевых моделей, исследователи искали в Google Play приложения, которые могли содержать встроенные модели глубокого обучения, по ключевым словам, связанным с глубоким обучением.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0002 Разведка</span><p>Исследователи провели обзор состязательных атак на ML-компоненты антивирусных продуктов. Они обнаружили, что техники, заимствованные из атак на классификаторы изображений, уже успешно применялись в области защиты от вредоносного ПО. Однако было неясно, эффективны ли такие подходы против ML-компонента промышленных антивирусных решений.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0002 Разведка</span><p>Использование компанией Kaspersky антивирусных детекторов на основе ML публично описано на сайте компании. На практике злоумышленник мог бы использовать эту информацию для выбора цели.</p></a>
</div>


Показано 12 из 41 примеров.
