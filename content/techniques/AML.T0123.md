---
atlas_id: AML.T0123
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may attempt to make an executable or file difficult to discover or analyze by encrypting, encoding, or otherwise obfuscating its contents on the system or in transit. This is common behavior that can be...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Enterprise
procedure_count: 2
source_name: Obfuscated Files or Information
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Obfuscated Files or Information
url: /techniques/AML.T0123/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may attempt to make an executable or file difficult to discover or analyze by encrypting, encoding, or otherwise obfuscating its contents on the system or in transit. This is common behavior that can be used across different platforms and the network to evade defenses.

Obfuscation may target AI-enabled defensive systems, including malware classifiers, content filters, secret scanners, and automated review systems. Content that appears benign or incomplete to a person or detector may be decoded, assembled, or interpreted by a downstream application, tool, or compromised system.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0007 Уклонение от защиты</span><p>The agents chunked, compressed, Base64-encoded, and sometimes XOR-encoded commands, payloads, credentials, and results carried through the external launchpad and public web services.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0007 Уклонение от защиты</span><p>The agents chunked, compressed, Base64-encoded, and sometimes XOR-encoded communications carried through the dataset-repository channel.</p></a>
</div>
