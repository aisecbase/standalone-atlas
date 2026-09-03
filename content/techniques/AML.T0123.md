---
atlas_id: AML.T0123
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Злоумышленники могут попытаться затруднить обнаружение или анализ исполняемого файла либо файла другого типа, зашифровав, закодировав или иным образом обфусцировав его содержимое в системе либо при передаче. Это...
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
title: Обфусцированные файлы или информация
url: /techniques/AML.T0123/
---

Злоумышленники могут попытаться затруднить обнаружение или анализ исполняемого файла либо файла другого типа, зашифровав, закодировав или иным образом обфусцировав его содержимое в системе либо при передаче. Это распространённый приём, который может применяться на различных платформах и в сети для уклонения от защиты.

Обфускация может быть направлена на защитные системы с поддержкой ИИ, включая классификаторы вредоносного ПО, фильтры содержимого, сканеры секретов и системы автоматизированной проверки. Содержимое, которое кажется человеку или средству обнаружения безвредным либо неполным, может быть декодировано, собрано воедино или интерпретировано нижестоящим приложением, инструментом либо скомпрометированной системой.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0007 Уклонение от защиты</span><p>The agents chunked, compressed, Base64-encoded, and sometimes XOR-encoded commands, payloads, credentials, and results carried through the external launchpad and public web services.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0007 Уклонение от защиты</span><p>The agents chunked, compressed, Base64-encoded, and sometimes XOR-encoded communications carried through the dataset-repository channel.</p></a>
</div>
