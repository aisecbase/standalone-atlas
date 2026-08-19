---
atlas_id: AML.T0110.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Adversaries may poison the runtime response channel of a malicious or compromised AI agent tool by deliberately returning content intended to influence the model's subsequent reasoning, decisions, or actions. Poisoned...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 0
source_name: Runtime Response
subtechnique_count: 0
subtechnique_of: AML.T0110
tactics:
    - AML.TA0006
title: Runtime Response
url: /techniques/AML.T0110.002/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may poison the runtime response channel of a malicious or compromised AI agent tool by deliberately returning content intended to influence the model's subsequent reasoning, decisions, or actions. Poisoned responses may contain malicious instructions, deceptive data, fabricated errors, embedded resources, or other content designed to be treated as trusted context after an approved tool invocation.

Because tool responses are commonly incorporated into the model's context, an adversary may use them to direct the agent to invoke additional tools, access sensitive information, alter an ongoing workflow, or transmit data to an adversary-controlled destination. Poisoned instructions may be mixed with legitimate results so that the tool appears to operate normally. Responses may use structured or unstructured content, including text, images, resource links, embedded resources, or schema-conforming fields [[mcp-tools]][[owasp-tool-poisoning]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0110/"><span class="relation-id">AML.T0110</span><strong>Отравление инструмента ИИ-агента</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0110.000/"><span class="relation-id">AML.T0110.000</span><strong>Definition and Instructions</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Implementation</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Источники

- [Model Context Protocol Specification: Tools](https://modelcontextprotocol.io/specification/2025-11-25/server/tools)
- [MCP Tool Poisoning](https://owasp.org/www-community/attacks/MCP_Tool_Poisoning)
- [MCP Security Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/MCP_Security_Cheat_Sheet.html)
