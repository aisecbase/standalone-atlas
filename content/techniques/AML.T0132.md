---
atlas_id: AML.T0132
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: AI agents and LLM platforms are deployed across diverse architectures, including standalone servers, SaaS platforms, and low-code builders. All of these often have misconfigured access controls, ranging from missing...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Misconfigured or Publicly Exposed AI Services
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Misconfigured or Publicly Exposed AI Services
url: /techniques/AML.T0132/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

AI agents and LLM platforms are deployed across diverse architectures, including standalone servers, SaaS platforms, and low-code builders. All of these often have misconfigured access controls, ranging from missing authentication in LLM runtimes to permissive public access settings, that significantly expand their attack surface.

AI agent's security posture affects how attackers are able to explore which agentic attack surface is available to them and directly either expands or limits an adversary's ability to discover and interact with a target's AI agents. For example, one internet-facing AI service could be more easily discoverable by any unauthenticated user, while another may require compromised credentials to even be discovered.

Adversaries may discover and interact with AI agents and LLM services as unauthenticated users via a combination of various OSINT and active scanning methods. These could include using search engines to discover accessible agentic deployments (See [Search Open Technical Databases](/techniques/AML.T0000)), search backlinks which could indicate existing or open agents which have been embedded to frontends (See [Search Open Website/Domains](/techniques/AML.T0095)), or directly scanning a target's AI infrastructure (See [Active Scanning](/techniques/AML.T0006)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>Honeypots can safely absorb real attacks and exploit attempts, turning payloads into mapped defensive signatures for a potential target without risking real assets.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0004 Первичный доступ</span><p>The researcher exploited a proxy misconfiguration present in ClawdBot&#39;s control server to gain access to control interfaces that had authentication enabled.</p></a>
</div>
