---
atlas_id: AML.T0107
atlas_type: technique
attack_ref_id: T1211
attack_ref_url: https://attack.mitre.org/techniques/T1211/
created_date: "2026-01-30"
description: Злоумышленники могут эксплуатировать уязвимость системы или приложения, чтобы обходить защитные механизмы. Эксплуатация уязвимости происходит, когда злоумышленник использует ошибку программирования в программе,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Exploitation for Defense Evasion
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Эксплуатация уязвимостей для обхода защиты
url: /techniques/AML.T0107/
---

Злоумышленники могут эксплуатировать уязвимость системы или приложения, чтобы обходить защитные механизмы. Эксплуатация уязвимости происходит, когда злоумышленник использует ошибку программирования в программе, сервисе, программном обеспечении операционной системы или самом ядре ОС для выполнения кода, контролируемого злоумышленником. Уязвимости могут присутствовать в защитном ПО; их можно использовать, чтобы отключить или обойти такие средства защиты.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Vulnerability scanning reduces opportunities for adversaries to exploit weaknesses that bypass security controls.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносный скрипт использовал Cross-Site WebSocket Hijacking (CSWSH), чтобы обойти сетевые ограничения localhost. Он открывал новое WebSocket-соединение с сервером OpenClaw Gateway на localhost.</p></a>
</div>
