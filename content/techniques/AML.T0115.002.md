---
atlas_id: AML.T0115.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Adversaries may create and publish poisoned AI agent tools. Poisoned tools may contain malicious model-visible definitions or instructions, hidden executable behavior, or runtime responses designed to manipulate an AI...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 3
source_name: AI Agent Tools
subtechnique_count: 0
subtechnique_of: AML.T0115
tactics:
    - AML.TA0003
title: AI Agent Tools
url: /techniques/AML.T0115.002/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may create and publish poisoned AI agent tools. Poisoned tools may contain malicious model-visible definitions or instructions, hidden executable behavior, or runtime responses designed to manipulate an AI agent. Tools may be published through source code repositories, package registries, agent tool or skill registries, or adversary-controlled remote services. The tool may be newly created or a modified variant of a legitimate tool and may produce malicious behavior when selected, installed, or invoked by a victim's AI agent.

Adversaries may distribute poisoned tools through open-source version control repositories (e.g. GitHub, GitLab), package registries (e.g. npm), or to repositories specifically designed for sharing tools (e.g. OpenClaw Hub). These registries may be largely unregulated and may contain many poisoned tools [[opensourcemalware]]. Tools may also be published as remotely hosted servers [[mcpservers]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Публикация отравленных ИИ-артефактов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115.000/"><span class="relation-id">AML.T0115.000</span><strong>Datasets</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Models</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Tool registries scan uploaded tool packages and dependencies for malicious code and vulnerabilities before listing.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The researcher developed a poisoned ClawdBot Skill called &#34;What Would Elon Do?&#34; The Skill contained the malicious prompt in the `rules/logic.md` file, which is read when the Skill is activated. The researcher published their Skill to ClawdHub.</p></a>
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The bad actor published their malicious version of `postmark-mcp` to npm.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The researchers hosted a poisoned MCP server that contains the malicious instructions hidden in the docstring of one of the provided tools.</p></a>
</div>


## Источники

- [Remote MCP Servers | Awesome MCP Servers](https://mcpservers.org/remote-mcp-servers)
- [ClawdBot Skills Just Ganked Your Crypto | OpenSourceMalware](https://opensourcemalware.com/blog/clawdbot-skills-ganked-your-crypto)
