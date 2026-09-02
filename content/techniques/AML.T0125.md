---
atlas_id: AML.T0125
atlas_type: technique
attack_ref_id: T1136
attack_ref_url: https://attack.mitre.org/techniques/T1136/
created_date: "2026-08-31"
description: Adversaries may create an account to maintain access to an AI system or its supporting infrastructure. With a sufficient level of access, creating such accounts may be used to establish secondary credentialed access...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 1
source_name: Create Account
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Create Account
url: /techniques/AML.T0125/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may create an account to maintain access to an AI system or its supporting infrastructure. With a sufficient level of access, creating such accounts may be used to establish secondary credentialed access that do not require persistent remote access tools to be deployed on the system.

Accounts may be created on local systems, in enterprise domains, cloud tenants, identity providers, or individual services. In AI environments, adversaries may create user, service, workload, or automation accounts in AI platforms, model or dataset repositories, development environments, experiment trackers, orchestration services, data stores, and other AI operations infrastructure.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0006 Закрепление</span><p>GTG-1002&#39;s jailbroken Claude agent created a local backdoor account to maintain access to a compromised environment.</p></a>
</div>
