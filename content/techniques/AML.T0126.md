---
atlas_id: AML.T0126
atlas_type: technique
attack_ref_id: T1119
attack_ref_url: https://attack.mitre.org/techniques/T1119/
created_date: "2026-08-31"
description: Adversaries may use automated techniques to collect data from AI systems and supporting enterprise environments. Automation may use scripts, command interpreters, command-line tools, or AI agent tools to identify,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 2
source_name: Automated Collection
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Automated Collection
url: /techniques/AML.T0126/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may use automated techniques to collect data from AI systems and supporting enterprise environments. Automation may use scripts, command interpreters, command-line tools, or AI agent tools to identify, retrieve, copy, or aggregate data without a human selecting each individual item.

Collection criteria may be fixed, such as file name, type, location, owner, date. Automation may also collect repeatedly, monitor for new material, traverse related resources, or combine data from local systems, cloud services, repositories, databases, object stores, and application APIs.

In AI environments, targeted material may include models, datasets, configurations, conversation histories, retrieval databases, deployment information, logs, and other operational data.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0009 Сбор материалов</span><p>GTG-1002&#39;s jailbroken Claude agent automatically collected and processed large volumes of victim data and categorized the results according to their intelligence value. It generated comprehensive documentation covering discovered services, harvested credentials, sensitive data, exploitation techniques, and attack progression to support subsequent campaign activity.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0009 Сбор материалов</span><p>Using the acquired accesses, the framework automatically retrieved and aggregated the reported personnel, account, configuration, credential, and network information.</p></a>
</div>
