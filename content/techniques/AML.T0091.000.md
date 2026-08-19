---
atlas_id: AML.T0091.000
atlas_type: technique
attack_ref_id: T1550.001
attack_ref_url: https://attack.mitre.org/techniques/T1550/001/
created_date: "2025-10-28"
description: Злоумышленники могут использовать украденные токены доступа к приложениям, чтобы обходить штатный процесс аутентификации и получать доступ к ограниченным учетным записям, информации или сервисам на удаленных системах....
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Application Access Token
subtechnique_count: 0
subtechnique_of: AML.T0091
tactics:
    - AML.TA0015
title: Токен доступа к приложению
url: /techniques/AML.T0091.000/
---

Злоумышленники могут использовать украденные токены доступа к приложениям, чтобы обходить штатный процесс аутентификации и получать доступ к ограниченным учетным записям, информации или сервисам на удаленных системах. Такие токены обычно похищают у пользователей или сервисов и используют вместо учетных данных для входа.

Токены доступа к приложениям используются для выполнения авторизованных API-запросов от имени пользователя или сервиса и часто применяются для доступа к ресурсам в облачных приложениях, контейнерных приложениях, software-as-a-service (SaaS) и AI-as-a-service (AIaaS). Они часто используются в ИИ-сервисах, таких как чат-боты, LLM и API предиктивного инференса.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0091/"><span class="relation-id">AML.T0091</span><strong>Использование альтернативных средств аутентификации</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0091.001/"><span class="relation-id">AML.T0091.001</span><strong>Cookie веб-сессии</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Enforce authorization and monitor production AI API use for anomalous activity associated with replayed access tokens.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0015 Латеральное перемещение</span><p>Злоумышленник использовал извлеченный токен, чтобы аутентифицироваться в backend-сервисе LLM.</p></a>
</div>
