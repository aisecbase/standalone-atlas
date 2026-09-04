---
atlas_id: AML.T0017.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: 'Злоумышленники могут разрабатывать или существенно адаптировать средства, предназначенные для расширения возможностей ИИ-агента: инструменты, интеграции или серверы, предоставляющие инструменты. С их помощью агент...'
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: AI Agent Tools
subtechnique_count: 0
subtechnique_of: AML.T0017
tactics:
    - AML.TA0003
title: Инструменты ИИ-агента
url: /techniques/AML.T0017.002/
---

Злоумышленники могут разрабатывать или существенно адаптировать средства, предназначенные для расширения возможностей ИИ-агента: инструменты, интеграции или серверы, предоставляющие инструменты. С их помощью агент может взаимодействовать с операционными системами, браузерами, сетями, облачными сервисами, хранилищами данных, репозиториями ПО, системами управления идентификацией или другими внешними ресурсами.

Инструменты ИИ-агента могут быть реализованы в виде серверов Model Context Protocol (MCP), плагинов, навыков, коннекторов, библиотек функций, адаптеров для управления компьютером, брокеров выполнения, удалённых API или аналогичных интерфейсов, которые может вызывать модель. Разработка может включать создание исполняемой логики, описаний инструментов, передаваемых модели, инструкций по порядку действий, схем входных данных, методов аутентификации и механизмов управления разрешениями, а также подготовку пакета, необходимого для предоставления агенту соответствующей возможности.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017/"><span class="relation-id">AML.T0017</span><strong>Разработка средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017.000/"><span class="relation-id">AML.T0017.000</span><strong>Состязательные атаки на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.001/"><span class="relation-id">AML.T0017.001</span><strong>Автономная разработка эксплойтов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The actor created two Hermes skills. web-terminal-exploitation encoded a procedure for unauthenticated WebSocket exploitation, while fofa-cyberspace-search instructed DeepSeek to use the actor&#39;s fofoapi.py script for internet asset enumeration. The observed FOFA workflow is consistent with the latter skill; the report does not attribute an action in the recovered session to web-terminal-exploitation.</p></a>
</div>
