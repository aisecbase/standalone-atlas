---
atlas_id: AML.T0046
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут перегружать ИИ-систему нерелевантными данными, из-за которых растёт число срабатываний. В результате аналитики организации-жертвы могут тратить время на проверку и исправление неверных инференсов....
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Spamming AI System with Chaff Data
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0011
title: Зашумление ИИ-системы нерелевантными данными
url: /techniques/AML.T0046/
---

Злоумышленники могут перегружать ИИ-систему нерелевантными данными, из-за которых растёт число срабатываний. В результате аналитики организации-жертвы могут тратить время на проверку и исправление неверных инференсов.

Злоумышленники также могут перегружать ИИ-агентов избыточными проверяемыми событиями низкой серьёзности или агентными действиями, которые требуют участия человека. Это заставляет организацию-жертву тратить время на ручную проверку работы агентной ИИ-системы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0004/"><span class="relation-id">AML.M0004</span><strong>Ограничение количества запросов к ИИ-модели</strong><p>Limit volume and rate of queries to protect the system from chaff data spam.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Аутентификация для моделей в продакшене может помочь предотвратить анонимный спам шумовыми данными.</p></a>
</div>
