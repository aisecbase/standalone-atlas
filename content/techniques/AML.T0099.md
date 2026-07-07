---
atlas_id: AML.T0099
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-25"
description: Злоумышленники могут размещать вредоносное содержимое в системе жертвы так, чтобы его мог извлечь инструмент ИИ-агента. Для этого они могут поместить документы в место, данные из которого принимает сервис, для работы...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 0
source_name: AI Agent Tool Data Poisoning
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Отравление данных инструмента ИИ-агента
url: /techniques/AML.T0099/
---

Злоумышленники могут размещать вредоносное содержимое в системе жертвы так, чтобы его мог извлечь инструмент ИИ-агента. Для этого они могут поместить документы в место, данные из которого принимает сервис, для работы с которым у ИИ-агента есть связанные инструменты.

Содержимое может быть подобрано так, чтобы часто попадать в результаты типовых запросов. Материалы злоумышленника могут содержать ложную или вводящую в заблуждение информацию. Они также могут включать промпт-инъекции с вредоносными инструкциями.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>
