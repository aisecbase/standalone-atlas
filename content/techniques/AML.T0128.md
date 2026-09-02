---
atlas_id: AML.T0128
atlas_type: technique
attack_ref_id: T1584
attack_ref_url: https://attack.mitre.org/techniques/T1584/
created_date: "2026-08-31"
description: Adversaries may compromise third-party infrastructure and repurpose it to support attacks against AI system. Rather than buying, leasing, registering, or otherwise legitimately acquiring a resource, the adversary...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 1
source_name: Compromise Infrastructure
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Compromise Infrastructure
url: /techniques/AML.T0128/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may compromise third-party infrastructure and repurpose it to support attacks against AI system. Rather than buying, leasing, registering, or otherwise legitimately acquiring a resource, the adversary gains unauthorized control of infrastructure owned or operated by another party.

Compromised infrastructure may include physical or cloud servers, domains, network devices, third-party web and DNS services, software or artifact repositories, development workspaces, compute services, and other externally hosted resources.

In operations involving AI systems, adversaries may compromise infrastructure used for model development, artifact hosting, dataset processing, evaluation, inference, agent tooling, or AI operations. They may repurpose this infrastructure to host malicious artifacts, stage payloads, run tools or agents, relay traffic, capture credentials, provide command and control, process collected data, or launch attacks against additional systems.

Compromised infrastructure may appear trustworthy because it uses a legitimate provider, established domain, valid certificate, reputable service, or expected AI development platform. It may also provide network access, compute resources, service identities, or trusted relationships that would be difficult for the adversary to establish directly.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>GTG-1002 operated dedicated penetration-testing servers accessible through MCP to support remote command execution, simultaneous tool coordination, and persistent operational state across campaign sessions.</p></a>
</div>
