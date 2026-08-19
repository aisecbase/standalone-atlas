---
atlas_id: AML.T0072
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2024-04-11"
description: Злоумышленники могут использовать реверс-шелл для связи с системой жертвы и управления ею. Обычно пользователь с помощью клиента подключается к удаленной машине, которая ожидает входящие соединения. При использовании...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 4
source_name: Reverse Shell
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0014
title: Реверс-шелл
url: /techniques/AML.T0072/
---

Злоумышленники могут использовать реверс-шелл для связи с системой жертвы и управления ею.

Обычно пользователь с помощью клиента подключается к удаленной машине, которая ожидает входящие соединения. При использовании реверс-шелла входящее соединение, инициированное системой жертвы, принимает злоумышленник.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0014/"><span class="relation-id">AML.TA0014</span><strong>Командование и управление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0014 Командование и управление</span><p>Имплант Sliver предоставляет исследователю канал командного управления, позволяя изучать среду жертвы и продолжать атаку.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0014 Командование и управление</span><p>Вредоносная нагрузка представляла собой reverse shell, настроенный на подключение к жестко заданному IP-адресу.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0014 Командование и управление</span><p>Python-код открывал reverse shell, который использовался как канал командования и управления.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Model Namespace Reuse Supply Chain Attack</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0014 Командование и управление</span><p>The payload established a reverse shell from the deployed endpoint to researcher-controlled infrastructure.</p></a>
</div>
