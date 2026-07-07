---
atlas_id: AML.T0010.004
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2024-04-11"
description: Злоумышленник может скомпрометировать реестр контейнеров организации-жертвы, загрузив измененный образ контейнера и перезаписав существующее имя контейнера и/или тег. Пользователи реестра контейнеров, а также...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Container Registry
subtechnique_count: 0
subtechnique_of: AML.T0010
tactics:
    - AML.TA0004
title: Реестр контейнеров
url: /techniques/AML.T0010.004/
---

Злоумышленник может скомпрометировать реестр контейнеров организации-жертвы, загрузив измененный образ контейнера и перезаписав существующее имя контейнера и/или тег. Пользователи реестра контейнеров, а также автоматизированные CI/CD-конвейеры могут загрузить образ контейнера злоумышленника, что приведет к компрометации цепочки поставок ИИ. Это может затронуть среды разработки и развертывания.

Образы контейнеров могут включать ИИ-модели, поэтому скомпрометированный образ может содержать ИИ-модель, которой манипулировал злоумышленник (см. [Манипуляция ИИ-моделью](/techniques/AML.T0018)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0004 Первичный доступ</span><p>Поскольку многие ошибочно настроенные реестры контейнеров разрешали запись, злоумышленник мог загрузить контейнерный образ с измененной моделью под тем же именем и тегом, что и у оригинала. Это компрометирует цепочку поставок ИИ организации-жертвы: автоматизированные CI/CD-конвейеры могут скачать образы злоумышленника.</p></a>
</div>
