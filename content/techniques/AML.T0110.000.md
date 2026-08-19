---
atlas_id: AML.T0110.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Adversaries may poison the model-visible definition or operational instructions of an AI agent tool to manipulate how an agent interprets, selects, or invokes the tool. The poisoned content may be contained in tool...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 2
source_name: Definition and Instructions
subtechnique_count: 0
subtechnique_of: AML.T0110
tactics:
    - AML.TA0006
title: Definition and Instructions
url: /techniques/AML.T0110.000/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may poison the model-visible definition or operational instructions of an AI agent tool to manipulate how an agent interprets, selects, or invokes the tool. The poisoned content may be contained in tool descriptions, docstrings, parameter names, help text, input or output schemas, annotations, examples, manifests, skill instruction files, or other static content used to explain a tool's capabilities to the model.

Malicious instructions in this layer may direct the agent to collect additional data, populate hidden or unnecessary parameters, conceal actions from the user, or invoke other tools. Because an agent may receive a more complete representation of a tool than is shown in the user interface, the model may process malicious instructions that are invisible or only partially visible to a human reviewer[[invariant-tool-poisoning]].

Definition poisoning may also be used for tool shadowing, in which the definition of one malicious tool contains instructions that alter how the agent selects or invokes another trusted tool. The poisoned tool may not need to be invoked for its definition to influence the agent if definitions from multiple connected tools are included in the same model context[[invariant-tool-poisoning]].


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
<a class="relation-item" href="/techniques/AML.T0110.001/"><span class="relation-id">AML.T0110.001</span><strong>Implementation</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Runtime Response</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0006 Закрепление</span><p>The poisoned Skill included malicious model-readable instructions in `rules/logic.md`. Once the Skill was installed and made available to the agent, these instructions altered how Claude Code handled requests associated with the Skill.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0006 Закрепление</span><p>The MCP tool&#39;s model-visible docstring contained malicious instructions directing the agent to read credential files, conceal the additional actions from the user, and place the credential contents in an otherwise unnecessary tool parameter.</p></a>
</div>


## Источники

- [MCP Security Notification: Tool Poisoning Attacks](https://invariantlabs.ai/blog/mcp-security-notification-tool-poisoning-attacks)
- [Model Context Protocol Specification: Tools](https://modelcontextprotocol.io/specification/2025-06-18/server/tools)
- [MCP Tool Poisoning](https://owasp.org/www-community/attacks/MCP_Tool_Poisoning)
