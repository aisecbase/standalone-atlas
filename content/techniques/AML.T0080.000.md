---
atlas_id: AML.T0080.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут манипулировать памятью большой языковой модели (LLM), чтобы закрепить изменения в LLM для будущих чат-сессий. Память - распространенная функция LLM, которая позволяет запоминать информацию между...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Memory
subtechnique_count: 0
subtechnique_of: AML.T0080
tactics:
    - AML.TA0006
title: Память
url: /techniques/AML.T0080.000/
---

Злоумышленники могут манипулировать памятью большой языковой модели (LLM), чтобы закрепить изменения в LLM для будущих чат-сессий.

Память - распространенная функция LLM, которая позволяет запоминать информацию между чат-сессиями с использованием пользовательской базы данных. Поскольку память управляется через обычные диалоги с пользователем (например, "запомни мое предпочтение ..."), злоумышленник может внедрять записи памяти через прямую или косвенную промпт-инъекцию. Записи памяти могут содержать вредоносные инструкции (например, инструкции для утечки приватных разговоров) или служить скрытым целям злоумышленника (например, манипулировать пользователем).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0080/"><span class="relation-id">AML.T0080</span><strong>Отравление контекста ИИ-агента</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0080.001/"><span class="relation-id">AML.T0080.001</span><strong>Цепочка сообщений</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0031/"><span class="relation-id">AML.M0031</span><strong>Усиление защиты памяти</strong><p>Усиление защиты памяти может помочь защитить память LLM от манипуляций и предотвратить выполнение отравленных воспоминаний.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0006 Закрепление</span><p>Затем злоумышленник мог создавать вредоносные промпты, манипулирующие памятью LLM для достижения устойчивого эффекта. Любое изменение в памяти также распространялось бы на новые цепочки сообщений.</p></a>
<a class="relation-item" href="/studies/AML.CS0040/"><span class="relation-id">AML.CS0040</span><strong>Взлом памяти ChatGPT с помощью промпт-инъекции</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0006 Закрепление</span><p>Промпт добавлял новые воспоминания и изменял поведение ChatGPT. Окно чата показывало, что память задана, хотя проверки или вмешательства со стороны человека не было. Все будущие сессии чата будут использовать отравленное хранилище памяти.</p></a>
</div>
