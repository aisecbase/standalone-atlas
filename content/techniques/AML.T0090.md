---
atlas_id: AML.T0090
atlas_type: technique
attack_ref_id: T1003
attack_ref_url: https://attack.mitre.org/techniques/T1003/
created_date: "2025-10-27"
description: Злоумышленники могут извлекать учетные данные из кэшей ОС, памяти приложений или других источников на скомпрометированной системе. Учетные данные часто представлены в виде хэша или открытого текста и могут включать...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: OS Credential Dumping
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0013
title: Дамп учетных данных ОС
url: /techniques/AML.T0090/
---

Злоумышленники могут извлекать учетные данные из кэшей ОС, памяти приложений или других источников на скомпрометированной системе. Учетные данные часто представлены в виде хэша или открытого текста и могут включать имена пользователей и пароли, токены приложений или другие ключи аутентификации.

Учетные данные могут использоваться для [латерального перемещения](/tactics/AML.TA0015) к другим ИИ-сервисам, таким как ИИ-агенты, LLM или API ИИ-инференса. Учетные данные также могут дать злоумышленнику доступ к другим программным инструментам и источникам данных, относящимся к жизненному циклу AI DevOps.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0013/"><span class="relation-id">AML.TA0013</span><strong>Доступ к учетным данным</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0011/"><span class="relation-id">AML.M0011</span><strong>Ограничение загрузки библиотек</strong><p>Restrict library loading to block credential-dumping paths that inject or load malicious code into processes holding authentication material.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0013 Доступ к учетным данным</span><p>Злоумышленник подключался к процессу или читал память напрямую из `/proc` в Linux либо открывал handle к процессу LLM-приложения в Windows. Затем он сканировал память процесса, чтобы извлечь токен аутентификации жертвы. Это можно легко сделать, применив регулярное выражение к каждой выделенной странице памяти процесса.</p></a>
</div>
