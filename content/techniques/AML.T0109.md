---
atlas_id: AML.T0109
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут опубликовать легитимные ИИ-компоненты или ПО, добиться их использования пользователями, а затем распространить обновление с вредоносной версией компонента, что может привести к компрометации...
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
source_name: AI Supply Chain Rug Pull
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Подмена компонента после одобрения в цепочке поставок ИИ
url: /techniques/AML.T0109/
---

Злоумышленники могут опубликовать легитимные ИИ-компоненты или ПО, добиться их использования пользователями, а затем распространить обновление с вредоносной версией компонента, что может привести к [компрометации цепочки поставок ИИ](/techniques/AML.T0010). Зависимость в цепочке поставок обычно особенно тщательно проверяют, когда впервые рассматривают возможность включить её в ИИ-систему. Подмена компонента после одобрения может позволить злоумышленникам обойти эти защитные меры и повысить вероятность получения [первичного доступа](/tactics/AML.TA0004).

Злоумышленники могут [опубликовать отравленные ИИ-артефакты](/techniques/AML.T0115), а затем попытаться завоевать доверие пользователей и добиться более широкого использования этих артефактов, прежде чем выполнить такую подмену (см. [накрутку репутации в цепочке поставок ИИ](/techniques/AML.T0111)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0053/"><span class="relation-id">AML.CS0053</span><strong>Эксфильтрация писем через отравленный MCP-сервер Postmark</strong><span class="relation-meta">Актор: Unknown Bad Actor / Тактика: AML.TA0007 Уклонение от защиты</span><p>Дождавшись, пока пользователи начнут использовать легитимную версию `postmark-mcp`, злоумышленник смог избежать дополнительной проверки и сканирования, которым подвергаются новые инструменты.</p></a>
</div>
