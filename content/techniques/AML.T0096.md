---
atlas_id: AML.T0096
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-12-24"
description: Злоумышленники могут использовать API ИИ-сервиса в системе организации-жертвы для обмена данными. Команды, отправляемые злоумышленником системе организации-жертвы, а часто и результаты их выполнения, передаются внутри...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: AI Service API
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0014
title: API ИИ-сервиса
url: /techniques/AML.T0096/
---

Злоумышленники могут использовать API ИИ-сервиса в системе организации-жертвы для обмена данными. Команды, отправляемые злоумышленником системе организации-жертвы, а часто и результаты их выполнения, передаются внутри обычного трафика ИИ-сервиса.

Канал команд и управления через API ИИ-сервиса может оставаться скрытым, поскольку команды злоумышленника выглядят как часть штатного обмена данными. Поэтому злоумышленник может применять эту технику, чтобы избежать обнаружения. Опора на уже развернутую инфраструктуру в системе организации-жертвы позволяет ему использовать имеющиеся в среде средства и уменьшать заметность операции.

API ИИ-сервисов могут использоваться как C2-каналы, когда злоумышленнику нужно действовать скрытно и сохранять долговременное закрепление для шпионской деятельности [[microsoft]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0014/"><span class="relation-id">AML.TA0014</span><strong>Командование и управление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Authenticate AI service API callers and monitor queries for policy violations and misuse.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0042/"><span class="relation-id">AML.CS0042</span><strong>SesameOp: новый бэкдор использует OpenAI Assistants API как C2-канал</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0014 Командование и управление</span><p>Злоумышленник использовал OpenAI Assistants API как канал передачи команд вредоносному ПО SesameOp. SesameOp выполнял эти команды в системе жертвы и отправлял результаты обратно злоумышленнику по тому же каналу. Команды и результаты передавались в зашифрованном виде. Чтобы скрыть следы, SesameOp удалял объекты Assistants и Messages, которые создавал и использовал для связи.</p></a>
</div>


## Источники

- [SesameOp: Novel backdoor uses OpenAI Assistants API for command and control | Microsoft Security Blog](https://www.microsoft.com/en-us/security/blog/2025/11/03/sesameop-novel-backdoor-uses-openai-assistants-api-for-command-and-control/)
