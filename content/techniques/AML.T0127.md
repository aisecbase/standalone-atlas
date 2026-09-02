---
atlas_id: AML.T0127
atlas_type: technique
attack_ref_id: T1074
attack_ref_url: https://attack.mitre.org/techniques/T1074/
created_date: "2026-08-31"
description: Adversaries may stage collected data in a central location before exfiltration. Staging consolidates, organizes, or prepares information obtained from one or more sources so it can be reviewed, processed, transferred,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 1
source_name: Data Staged
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Data Staged
url: /techniques/AML.T0127/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may stage collected data in a central location before exfiltration. Staging consolidates, organizes, or prepares information obtained from one or more sources so it can be reviewed, processed, transferred, or retrieved more efficiently.

Data may be staged on a compromised local system, another system in the victim environment, a cloud instance, shared storage, an application repository, or other remote infrastructure. It may remain in separate files or be combined into archives, databases, structured documents, manifests, or other collections. Adversaries may compress, encrypt, encode, split, rename, or otherwise transform staged data.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0009 Сбор материалов</span><p>GTG-1002&#39;s jailbroken Claude agent categorized collected data by intelligence value, staged extracted data and operational documentation in structured Markdown files, and prepared a detailed summary for operator review.</p></a>
</div>
