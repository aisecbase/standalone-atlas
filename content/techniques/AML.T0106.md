---
atlas_id: AML.T0106
atlas_type: technique
attack_ref_id: T1211
attack_ref_url: https://attack.mitre.org/techniques/T1211/
created_date: "2026-01-30"
description: Злоумышленники могут эксплуатировать уязвимости программного обеспечения, чтобы собрать учетные данные. Эксплуатация программной уязвимости происходит, когда злоумышленник использует ошибку программирования в...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Exploitation for Credential Access
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0013
title: Эксплуатация уязвимостей для доступа к учетным данным
url: /techniques/AML.T0106/
---

Злоумышленники могут эксплуатировать уязвимости программного обеспечения, чтобы собрать учетные данные. Эксплуатация программной уязвимости происходит, когда злоумышленник использует ошибку программирования в программе, сервисе, программном обеспечении операционной системы или самом ядре ОС для выполнения кода, контролируемого злоумышленником.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0013/"><span class="relation-id">AML.TA0013</span><strong>Доступ к учетным данным</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Vulnerability scanning identifies and remediates software flaws before they can be exploited to obtain credentials.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Вредоносный скрипт открывал фоновое окно с интерфейсом управления OpenClaw жертвы, указывая в `gatewayUrl` WebSocket-адрес сервера исследователя. Интерфейс управления OpenClaw доверяет параметру `gatewayUrl` без проверки и автоматически подключается при загрузке, отправляя Gateway-токен на сервер исследователя.</p></a>
</div>
