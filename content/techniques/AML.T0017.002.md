---
atlas_id: AML.T0017.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may develop or materially adapt tools, integrations, or tool servers designed to extend the capabilities of an AI agent. These capabilities may allow an agent to interact with operating systems, browsers,...
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
title: AI Agent Tools
url: /techniques/AML.T0017.002/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may develop or materially adapt tools, integrations, or tool servers designed to extend the capabilities of an AI agent. These capabilities may allow an agent to interact with operating systems, browsers, networks, cloud services, data stores, software repositories, identity systems, or other external resources.

AI agent tools may be implemented as Model Context Protocol servers, plugins, skills, connectors, function libraries, computer-use adapters, execution brokers, remote APIs, or similar model-callable interfaces. Development may include creating executable functionality, model-visible tool descriptions, procedural instructions, input schemas, authentication methods, permission handling, or packaging needed to make a capability available to an agent.


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
