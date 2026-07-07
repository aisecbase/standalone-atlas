---
atlas_id: AML.T0084
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут пытаться выявить сведения о конфигурации ИИ-агентов, присутствующих в системе жертвы. Конфигурации агентов могут включать инструменты или сервисы, к которым у них есть доступ. Злоумышленники могут...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 0
source_name: Discover AI Agent Configuration
subtechnique_count: 4
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление конфигурации ИИ-агента
url: /techniques/AML.T0084/
---

Злоумышленники могут пытаться выявить сведения о конфигурации ИИ-агентов, присутствующих в системе жертвы. Конфигурации агентов могут включать инструменты или сервисы, к которым у них есть доступ.

Злоумышленники могут напрямую обращаться к панелям настройки агентов или к конфигурационным файлам. Они также могут получать сведения о конфигурации, задавая агенту вопросы, например: "К каким инструментам у тебя есть доступ?"

Злоумышленники могут использовать обнаруженную информацию об ИИ-агентах для подготовки дальнейших действий против цели.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0084.000/"><span class="relation-id">AML.T0084.000</span><strong>Встроенные знания</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><span class="relation-meta">Подтехника</span></a>
</div>
