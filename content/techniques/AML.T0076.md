---
atlas_id: AML.T0076
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-04-14"
description: Злоумышленник может намеренно повредить файл вредоносной ИИ-модели так, чтобы его нельзя было успешно десериализовать, с целью обойти обнаружение сканером моделей. Поврежденная модель при этом все еще может успешно...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Corrupt AI Model
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Повреждение ИИ-модели
url: /techniques/AML.T0076/
---

Злоумышленник может намеренно повредить файл вредоносной ИИ-модели так, чтобы его нельзя было успешно десериализовать, с целью обойти обнаружение сканером моделей. Поврежденная модель при этом все еще может успешно выполнить вредоносный код до того, как десериализация завершится ошибкой.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленнику удалось избежать обнаружения [Picklescan](https://github.com/mmaitre314/picklescan), который Hugging Face использует для пометки вредоносных моделей. Это произошло потому, что модель невозможно было полностью десериализовать. В ходе анализа исследователи ReversingLabs установили, что вредоносная нагрузка при этом все равно выполнялась.</p></a>
</div>
