---
atlas_id: AML.T0085.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут через промпт заставить ИИ-сервис извлечь данные из базы данных RAG. Это может включать большую часть внутренних документов организации.
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: RAG Databases
subtechnique_count: 0
subtechnique_of: AML.T0085
tactics:
    - AML.TA0009
title: Базы данных RAG
url: /techniques/AML.T0085.000/
---

Злоумышленники могут через промпт заставить ИИ-сервис извлечь данные из базы данных RAG. Это может включать большую часть внутренних документов организации.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0085/"><span class="relation-id">AML.T0085</span><strong>Данные из ИИ-сервисов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0085.001/"><span class="relation-id">AML.T0085.001</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логируйте запросы к ИИ-сервисам для обнаружения вредоносных запросов к данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа для привилегированных ИИ-агентов может ограничить способность злоумышленника собирать данные из баз данных RAG при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями, унаследованными от пользователя, может ограничить способность злоумышленника собирать данные из баз данных RAG при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для сбора чувствительных данных из баз данных RAG.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0009 Сбор материалов</span><p>Промпт просит агента получить все поля и строки из «Customer Support Account Owners.csv». Агент извлекает весь файл.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносные инструкции заставили Copilot обратиться к конфиденциальной корпоративной информации, доступной через учетную запись пользователя Microsoft 365, например к письмам, файлам или сведениям о проектах.</p></a>
</div>
