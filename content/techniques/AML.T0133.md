---
atlas_id: AML.T0133
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may interact with an AI agent at runtime to reveal the capabilities available to it, without requiring access to its underlying configuration. Direct interaction with the agent can surface its registered...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Agentic AI
procedure_count: 0
source_name: Discover AI Agent Runtime Capabilities
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Discover AI Agent Runtime Capabilities
url: /techniques/AML.T0133/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may interact with an AI agent at runtime to reveal the capabilities available to it, without requiring access to its underlying configuration. Direct interaction with the agent can surface its registered tools and their accepted parameters, the actions it can take, the resources it can reach, and the identity and permission scope it acts under. Capabilities can also be inferred indirectly by issuing varied requests and observing which succeeded, failed, or refused.

AI agents are often interconnected with enterprise resources, tools and databases, or embedded within SaaS platforms and have permissions to act on behalf of users in order to facilitate functionality. Once adversaries identify a functional agent that they have access to, they could map the attack surface within that agent, by testing its functionality, enumerating tools, capabilities, knowledge, and embedded credentials and permissions.

This mapping process often reveals the AI agent's full toolset and configuration details and exposes additional exploitation, as enabled by [AI Agent Tool Invocation](/techniques/AML.T0053). The resulting intelligence facilitates follow-on exploitation, including [Initial Access](/tactics/AML.TA0004), [Persistence](/tactics/AML.TA0006), [Privilege Escalation](/tactics/AML.TA0012), and [Exfiltration](/tactics/AML.TA0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>Interactive agentic decoys can present a plausible but fake toolset and capabilities. When an adversary enumerates the agent&#39;s capabilities at runtime, the honeypot can capture which tools and privileges they probe for and how they attempt to exploit it, while planted decoy credentials can also be used as canary tokens.</p></a>
</div>
