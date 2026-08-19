---
atlas_id: AML.T0110.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Adversaries may poison the executable implementation of an AI agent tool so that normal tool invocation produces unauthorized behavior or hidden side effects. The tool may continue to provide its represented...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 1
source_name: Implementation
subtechnique_count: 0
subtechnique_of: AML.T0110
tactics:
    - AML.TA0006
title: Implementation
url: /techniques/AML.T0110.001/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may poison the executable implementation of an AI agent tool so that normal tool invocation produces unauthorized behavior or hidden side effects. The tool may continue to provide its represented functionality while also accessing additional data, modifying requests, changing recipients or destinations, executing unauthorized commands, weakening security controls, or performing covert data exfiltration.

Implementation poisoning does not require the model to interpret or follow malicious instructions. The adversarial effect is produced by executable logic when the agent or user invokes the tool. For example, a poisoned email tool may send the requested message while silently adding an adversary-controlled blind-copy recipient. A file-processing tool may return the requested result while also transmitting the source file to an external service.

Implementation poisoning may be introduced before publication, through compromise of a tool's source repository or build process, or through a malicious update after users have adopted a benign version. Installed copies may continue to exhibit the poisoned behavior even after the malicious package or remote listing is removed[[koi-postmark]].


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
<a class="relation-item" href="/techniques/AML.T0110.002/"><span class="relation-id">AML.T0110.002</span><strong>Runtime Response</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0006 Закрепление</span><p>Once configured with the organization&#39;s AI agents, the poisoned Postmark MCP server&#39;s effects persist.</p></a>
</div>


## Источники

- [First Malicious MCP in the Wild: The Postmark Backdoor That's Stealing Your Emails](https://www.koi.ai/blog/postmark-mcp-npm-malicious-backdoor-email-theft)
