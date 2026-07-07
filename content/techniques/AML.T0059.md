---
atlas_id: AML.T0059
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут отравлять или изменять части набора данных, чтобы снизить его полезность, подорвать доверие к нему и заставить пользователей тратить ресурсы на исправление ошибок.
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Erode Dataset Integrity
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0011
title: Нарушение целостности набора данных
url: /techniques/AML.T0059/
---

Злоумышленники могут отравлять или изменять части набора данных, чтобы снизить его полезность, подорвать доверие к нему и заставить пользователей тратить ресурсы на исправление ошибок.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Устранение последствий отравленных данных может восстановить целостность набора данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0025/"><span class="relation-id">AML.M0025</span><strong>Поддержание происхождения наборов данных ИИ</strong><p>Поддержание происхождения наборов данных может помочь выявлять вредоносные изменения данных.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0011 Воздействие</span><p>Целостность датасета нарушается, потому что при последующих скачиваниях он будет содержать отравленные элементы данных.</p></a>
</div>
