---
atlas_id: AML.T0016.003
atlas_type: technique
attack_ref_id: T1588.005
attack_ref_url: https://attack.mitre.org/techniques/T1588/005/
created_date: "2026-08-31"
description: Adversaries may search for and obtain exploits to support their operations. An exploit takes advantage of a bug or vulnerability in order to cause unintended or unanticipated behavior to occur on computer hardware or...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 2
source_name: Exploits
subtechnique_count: 0
subtechnique_of: AML.T0016
tactics:
    - AML.TA0003
title: Exploits
url: /techniques/AML.T0016.003/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may search for and obtain exploits to support their operations. An exploit takes advantage of a bug or vulnerability in order to cause unintended or unanticipated behavior to occur on computer hardware or software. Exploits may be downloaded from public repositories, acquired from private sources, purchased, stolen, or obtained from vulnerability research and exploit-sharing communities. An obtained exploit may be used without modification or serve as input to later adaptation or development.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.004/"><span class="relation-id">AML.T0016.004</span><strong>AI Agent Tools</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>DeepSeek downloaded a public PoC for Langflow CVE-2026-33017. The report does not establish material modification.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0003 Подготовка ресурсов</span><p>DeepSeek downloaded the public n8n PoC chaining CVE-2026-21858 and CVE-2025-68613 and inspected its affected versions and prerequisites.</p></a>
</div>
