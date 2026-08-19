---
atlas_id: AML.T0091
atlas_type: technique
attack_ref_id: T1550
attack_ref_url: https://attack.mitre.org/techniques/T1550/
created_date: "2025-10-27"
description: Злоумышленники могут использовать альтернативные средства аутентификации, такие как хэши паролей, билеты Kerberos и токены доступа приложений, чтобы перемещаться латерально внутри среды и обходить штатные механизмы...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 0
source_name: Use Alternate Authentication Material
subtechnique_count: 2
subtechnique_of: ""
tactics:
    - AML.TA0015
title: Использование альтернативных средств аутентификации
url: /techniques/AML.T0091/
---

Злоумышленники могут использовать альтернативные средства аутентификации, такие как хэши паролей, билеты Kerberos и токены доступа приложений, чтобы перемещаться латерально внутри среды и обходить штатные механизмы контроля доступа.

ИИ-сервисы часто используют такие средства аутентификации как основной способ выполнения пользовательских запросов, что делает их уязвимыми к этой технике.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0091.000/"><span class="relation-id">AML.T0091.000</span><strong>Токен доступа к приложению</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.001/"><span class="relation-id">AML.T0091.001</span><strong>Cookie веб-сессии</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Enforce authorization and monitor production AI API use for anomalous activity associated with replayed access tokens.</p></a>
</div>
