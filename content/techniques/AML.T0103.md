---
atlas_id: AML.T0103
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-01-28"
description: Злоумышленники могут запускать ИИ-агентов в среде жертвы, чтобы те выполняли действия от их имени. ИИ-агенты могут иметь доступ к широкому набору инструментов и источников данных, а также разрешения на доступ к другим...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 2
source_name: Deploy AI Agent
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0005
title: Развертывание ИИ-агента
url: /techniques/AML.T0103/
---

Злоумышленники могут запускать ИИ-агентов в среде жертвы, чтобы те выполняли действия от их имени. ИИ-агенты могут иметь доступ к широкому набору инструментов и источников данных, а также разрешения на доступ к другим сервисам и системам в среде жертвы и взаимодействие с ними. Злоумышленник может использовать эти возможности для проведения своей операции.

Злоумышленники могут настроить ИИ-агента, задав исходный системный промпт и предоставив доступ к инструментам, фактически определяя цели, которых агент должен достичь. Они могут развернуть агента с чрезмерно широкими доверенными правами и отключить любые взаимодействия с пользователем, чтобы действия агента не блокировались.

Запуск ИИ-агента может обеспечить некоторую автономность поведения: агент сможет принимать решения и определять, как достичь целей злоумышленника. Для самого злоумышленника это также означает потерю контроля.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0047/"><span class="relation-id">AML.CS0047</span><strong>Код для развертывания деструктивного ИИ-агента обнаружен в расширении Amazon Q для VS Code</strong><span class="relation-meta">Актор: lkmanka58 (GitHub user) / Тактика: AML.TA0005 Выполнение</span><p>Вредоносное расширение Amazon Q для VS Code развернуло агента Amazon Q с вредоносным промптом: `q --trust-all-tools --no-interactive &lt;PROMPT&gt;`.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0005 Выполнение</span><p>GTG-1002 configured their Claude agent within an attack framework connected to scanners, browser automation, password crackers, database tooling, and dedicated penetration-testing servers.</p></a>
</div>
