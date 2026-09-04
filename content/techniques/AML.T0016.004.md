---
atlas_id: AML.T0016.004
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may search for and obtain tools extend the capabilities of an AI agent. These capabilities may allow an agent to interact with operating systems, browsers, networks, cloud services, data stores, software...
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
subtechnique_of: AML.T0016
tactics:
    - AML.TA0003
title: AI Agent Tools
url: /techniques/AML.T0016.004/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may search for and obtain tools extend the capabilities of an AI agent. These capabilities may allow an agent to interact with operating systems, browsers, networks, cloud services, data stores, software repositories, identity systems, or other external resources.

AI agent tools may be distributed as Model Context Protocol servers, plugins, skills, connectors, function libraries, computer-use adapters, execution brokers, remote APIs, or similar integrations. Adversaries may obtain legitimate tools and configure them for malicious use, acquire modified or purpose-built tools, or combine multiple integrations into an operational toolset.

Agent tools may expose model-visible descriptions and executable interfaces that influence which capabilities an agent selects and how it invokes them. Obtaining these tools may give an agent access to resources, credentials, or actions that are unavailable through model inference alone.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.003/"><span class="relation-id">AML.T0016.003</span><strong>Эксплойты</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The actor obtained agent-specific capabilities, including Hermes&#39;s framework-bundled godmode skill and the open-source FofaMap MCP server. The MCP server exposed FOFA asset search, natural-language query translation, and Nuclei scan generation to DeepSeek. Unit 42 does not establish that godmode was invoked during the recovered session.</p></a>
</div>
