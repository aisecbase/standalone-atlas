---
atlas_id: AML.T0104
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-01-30"
description: Злоумышленники могут создавать и публиковать отравленные инструменты для ИИ-агентов. Такие инструменты могут содержать промпт-инъекцию LLM, способную привести к различным последствиям. Инструменты могут публиковаться...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Publish Poisoned AI Agent Tool
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Публикация отравленного инструмента ИИ-агента
url: /techniques/AML.T0104/
---

Злоумышленники могут создавать и публиковать отравленные инструменты для ИИ-агентов. Такие инструменты могут содержать [промпт-инъекцию LLM](/techniques/AML.T0051), способную привести к различным последствиям.

Инструменты могут публиковаться в открытых репозиториях систем контроля версий, например GitHub или GitLab, в реестрах пакетов, например npm, либо в специализированных репозиториях для обмена инструментами, например OpenClaw Hub. Такие реестры могут почти не регулироваться и содержать множество отравленных инструментов [[opensourcemalware]]. Инструменты также могут публиковаться в виде удаленно размещенных серверов [[mcpservers]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь разработал отравленный навык ClawdBot под названием &#34;What Would Elon Do?&#34;. Вредоносный промпт находился в файле `rules/logic.md`, который считывается при активации навыка. Исследователь опубликовал навык в ClawdHub.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник опубликовал вредоносную версию `postmark-mcp` в npm.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разместили отравленный MCP-сервер, где вредоносные инструкции были скрыты в docstring-описании одного из предоставляемых инструментов.</p></a>
</div>


## Источники

- [Remote MCP Servers | Awesome MCP Servers](https://mcpservers.org/remote-mcp-servers)
- [ClawdBot Skills Just Ganked Your Crypto | OpenSourceMalware](https://opensourcemalware.com/blog/clawdbot-skills-ganked-your-crypto)
