---
atlas_id: AML.T0000
atlas_type: technique
attack_ref_id: T1596
attack_ref_url: https://attack.mitre.org/techniques/T1596/
created_date: "2021-05-13"
description: Злоумышленники могут искать общедоступные исследования и техническую документацию, чтобы понять, как и где ИИ используется в организации-жертве. Злоумышленник может использовать эти сведения, чтобы определить цели...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 10
source_name: Search Open Technical Databases
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Поиск в открытых технических базах данных
url: /techniques/AML.T0000/
---

Злоумышленники могут искать общедоступные исследования и техническую документацию, чтобы понять, как и где ИИ используется в организации-жертве.

Злоумышленник может использовать эти сведения, чтобы определить цели атаки или адаптировать существующую атаку и тем самым повысить её эффективность.

Организации часто используют в продакшене модели, построенные на основе открытых архитектур и обученные на дополнительных проприетарных данных.

Знание лежащей в основе архитектуры позволяет злоумышленнику создавать более реалистичные прокси-модели ([Создание прокси-модели ИИ](/techniques/AML.T0005)).

Злоумышленник может искать в этих источниках публикации авторов, работающих в организации-жертве.

Исследовательские и технические материалы могут быть представлены научными статьями, опубликованными в [журналах и материалах конференций](/techniques/AML.T0000.000) или размещёнными в [репозиториях препринтов](/techniques/AML.T0000.001), а также [техническими блогами](/techniques/AML.T0000.002).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0000.000/"><span class="relation-id">AML.T0000.000</span><strong>Журналы и материалы конференций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.001/"><span class="relation-id">AML.T0000.001</span><strong>Репозитории препринтов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0000.002/"><span class="relation-id">AML.T0000.002</span><strong>Технические блоги</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничьте связь между публично раскрытыми подходами и данными, моделями и алгоритмами, используемыми в продакшене.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0002 Разведка</span><p>Обнаружение DGA широко используется для выявления ботнетов в академической среде и индустрии. Исследовательская группа искала научные публикации, связанные с обнаружением DGA.</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0002 Разведка</span><p>Исследователи изучили публично доступную информацию об ИИ-детекторе вредоносного ПО Cylance. Они собирали эти сведения из разных источников, включая публичные выступления и патентные заявки Cylance.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0002 Разведка</span><p>Исследователи использовали опубликованные научные статьи, чтобы определить наборы данных и архитектуры моделей, применявшиеся целевыми сервисами машинного перевода.</p></a>
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0002 Разведка</span><p>Используя публичную документацию по GPT-2, исследователи собрали сведения о наборе данных, архитектуре модели и гиперпараметрах обучения.</p></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0002 Разведка</span><p>Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0002 Разведка</span><p>Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0002 Разведка</span><p>Команда сначала провела разведку, чтобы собрать сведения о целевой ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0002 Разведка</span><p>Исследователь искал цели в Shodan по заголовку веб-интерфейса управления ClawdBot — `Clawdbot Control` — и обнаружил сотни интерфейсов ClawdBot, открытых в публичном интернете.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek выполнил запрос в FOFA и получил записи о 84 экземплярах Langflow, доступных из интернета. Эти записи указывали лишь на доступность экземпляров извне и не подтверждали, что они являлись уязвимыми целями.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Злоумышленник использовал Hermes Agent на базе DeepSeek при попытках эксплуатации Langflow и n8n</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek queried FOFA for n8n deployments. FOFA reported 647,017 global results and 25,209 in China; these were not confirmed vulnerable systems.</p></a>
</div>
